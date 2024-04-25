package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"

	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/loggerx"
	"github.com/kubeshop/botkube/pkg/multierror"
	"github.com/sirupsen/logrus"
)

const (
	marketingSitemapURL = "https://botkube.io/sitemap.xml"
	docsSitemapURL      = "https://docs.botkube.io/sitemap.xml"
	processingAPIURL    = "https://r.jina.ai"
	contentDir          = "content"
	maxRetries          = 5
	retryInterval       = 1 * time.Second
	httpCliTimeout      = 1 * time.Minute
)

var excludedDocsPagesRegex = regexp.MustCompile(`^https:\/\/docs\.botkube\.io\/(?:\d+\.\d+|next)\/.*`)

func main() {
	log := loggerx.New(config.Logger{
		Level:     "info",
		Formatter: "json",
	})

	fetcher := &contentFetcher{
		log: log,
		httpCli: &http.Client{
			Timeout: httpCliTimeout,
		},
	}

	log.Infof("Removing old %q directory...", contentDir)
	err := os.RemoveAll(contentDir)
	loggerx.ExitOnError(err, "while removing old directory")

	log.Info("Fetching Botkube sitemap...")
	marketingPages, err := fetcher.getURLsToDownloadFromSitemap(marketingSitemapURL)
	loggerx.ExitOnError(err, "while fetching Botkube sitemap")

	log.Info("Fetching Botkube docs sitemap...")
	docsPages, err := fetcher.getURLsToDownloadFromSitemap(docsSitemapURL)
	loggerx.ExitOnError(err, "while fetching Botkube docs sitemap")

	log.Info("Preparing list of pages to fetch...")
	pagesToFetch := fetcher.preparePageList(docsPages, marketingPages)
	log.Infof("Found %d pages to fetch", len(pagesToFetch))

	log.Infof("Creating %q directory...", contentDir)
	err = os.MkdirAll(contentDir, os.ModePerm)
	loggerx.ExitOnError(err, "while creating directory")

	errs := multierror.New()
	for i, page := range pagesToFetch {
		filePath, err := fetcher.filePathForURL(page)
		if err != nil {
			errs = multierror.Append(errs, err)
			continue
		}
		log.WithFields(logrus.Fields{
			"url":      page,
			"filePath": filePath,
		}).Infof("Fetching and saving page %d of %d...", i+1, len(pagesToFetch))

		err = retry.Do(
			func() error {
				return fetcher.fetchAndSavePage(page, filePath)
			},
			retry.Attempts(maxRetries),
			retry.OnRetry(func(n uint, err error) {
				log.WithError(err).Errorf("while fetching and saving page %q. Retrying...", page)
			}),
			retry.Delay(retryInterval),
		)

		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	loggerx.ExitOnError(errs.ErrorOrNil(), "while fetching and saving docs pages")

	log.Infof("Saved %d docs pages", len(pagesToFetch))
}

type contentFetcher struct {
	log     logrus.FieldLogger
	httpCli *http.Client
}

type sitemapURLSet struct {
	URLs []sitemapURL `xml:"url"`
}

type sitemapURL struct {
	Loc string `xml:"loc"`
}

func (f *contentFetcher) getURLsToDownloadFromSitemap(sitemapURL string) ([]string, error) {
	log := f.log.WithField("sitemapURL", sitemapURL)
	// nolint:gosec
	res, err := http.Get(sitemapURL)
	if err != nil {
		return nil, fmt.Errorf("while fetching sitemap %q: %w", sitemapURL, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code when fetching Botkube sitemap: %d", res.StatusCode)
	}

	log.Info("Decoding sitemap...")
	var sitemap sitemapURLSet
	err = xml.NewDecoder(res.Body).Decode(&sitemap)
	if err != nil {
		return nil, fmt.Errorf("while decoding sitemap %q: %w", sitemapURL, err)
	}

	var urls []string
	for _, part := range sitemap.URLs {
		urls = append(urls, part.Loc)
	}

	log.Infof("Found %d sitemap entries", len(urls))
	return urls, nil
}

func (f *contentFetcher) fetchAndSavePage(inURL, filePath string) error {
	pageURL := fmt.Sprintf("%s/%s", processingAPIURL, inURL)

	req, err := http.NewRequest(http.MethodGet, pageURL, nil)
	if err != nil {
		return fmt.Errorf("while creating request for page %q: %w", pageURL, err)
	}
	req.Header.Set("Content-Type", "text/event-stream")

	res, err := f.httpCli.Do(req)
	if err != nil {
		return fmt.Errorf("while fetching page %q: %w", pageURL, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code when fetching page %q: %d", pageURL, res.StatusCode)
	}

	// nolint:gosec
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("while creating file %q: %w", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("while writing to file %q: %w", filePath, err)
	}

	return nil
}

func (f *contentFetcher) preparePageList(docsPages, marketingPages []string) []string {
	var out []string
	for _, page := range docsPages {
		// remove all docs for previous and upcoming versions
		if excludedDocsPagesRegex.MatchString(page) {
			continue
		}

		out = append(out, strings.TrimSpace(page))
	}
	for _, page := range marketingPages {
		out = append(out, strings.TrimSpace(page))
	}

	return out
}

func (f *contentFetcher) filePathForURL(inURL string) (string, error) {
	parsedInURL, err := url.Parse(inURL)
	if err != nil {
		return "", fmt.Errorf("while parsing url %q: %w", inURL, err)
	}

	prefix := parsedInURL.Host
	urlPath := strings.Trim(parsedInURL.Path, "/")
	urlPath = strings.Replace(urlPath, "/", "__", -1)

	fileName := prefix
	if urlPath != "" {
		fileName = fmt.Sprintf("%s__%s", prefix, urlPath)
	}

	return fmt.Sprintf("%s/%s.md", contentDir, fileName), nil
}
