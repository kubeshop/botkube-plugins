package aibrain

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/kubeshop/botkube/pkg/api"
)

const (
	teamsMessageIDSubstr = "thread.tacv2"
)

var (
	quickResponses = []string{
		"Just a moment, please...",
		"Thinking about this one...",
		"Let me check on that for you.",
		"Processing your request...",
		"Working on it!",
		"This one needs some extra thought.",
		"I'm carefully considering your request.",
		"Consulting my super-smart brain...",
		"Cogs are turning...",
		"Accessing the knowledge archives...",
		"Running calculations at lightning speed!",
		"Hold on tight, I'm diving into the details.",
		"I'm here to help!",
		"Happy to look into this for you.",
		"Always learning to do this better.",
		"I want to get this right for you.",
		"Let me see what I can find out.",
		"My circuits are buzzing!",
		"Let me consult with my owl advisor...",
		"Consider it done (or at least, I'll try my best!)",
		"I'll get back to you with the best possible answer.",
	}

	multiLineCodeBlockWithLang = regexp.MustCompile(`(?m)\x60\x60\x60.*`) // \x60 is the backtick
	mdBold                     = regexp.MustCompile(`\*\*(.*?)\*\*`)
	mdStrikethrough            = regexp.MustCompile(`~~([^~]+)~~`)
	mdHeadings                 = regexp.MustCompile(`(?m)^#+ (.*)$`) // (?m) for multiline, see the "grouping" section at https://pkg.go.dev/regexp/syntax
	mdLinks                    = regexp.MustCompile(`\[([^\]]+)]\(([^)]+)\)`)
	mdImages                   = regexp.MustCompile(`!\[([^\]]+)]\(([^)]+)\)`)
)

func pickQuickResponse(messageID string) api.Message {
	rand.New(rand.NewSource(time.Now().UnixNano())) // #nosec G404
	i := rand.Intn(len(quickResponses))             // #nosec G404

	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: quickResponses[i]},
				},
			},
		},
	}
}

func msgUnableToHelp(messageID string) api.Message {
	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: "I am sorry, something went wrong, please try again. 😔"},
				},
			},
		},
	}
}

func msgQuotaExceeded(messageID string) api.Message {
	btnBuilder := api.NewMessageButtonBuilder()

	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: "⚠️ Quota exceeded for current calendar month. Upgrade the Botkube Cloud plan to continue."},
				},
				Buttons: []api.Button{
					btnBuilder.ForURL("Open Botkube Cloud", "https://app.botkube.io"),
				},
			},
		},
	}
}

func msgNoAIAnswer(messageID string) api.Message {
	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: "I am sorry, but I don't have a good answer."},
				},
			},
		},
	}
}

func msgAIAnswer(messageID, text string) api.Message {
	if strings.Contains(messageID, teamsMessageIDSubstr) {
		return api.Message{
			ParentActivityID: messageID,
			BaseBody: api.Body{
				// We use the Plaintext to make sure that Teams renderer will send
				// as simplified card (https://learn.microsoft.com/en-us/microsoftteams/platform/bots/how-to/format-your-bot-messages#format-text-content)
				// instead of AdaptiveCard (https://learn.microsoft.com/en-us/microsoftteams/platform/task-modules-and-cards/cards/cards-format?tabs=adaptive-md%2Cdesktop%2Cconnector-html#format-cards-with-markdown)
				// which doesn't support most of the markdown elements.
				Plaintext: markdownToTeams(text),
			},
		}
	}

	// the Sections.Base.Body is rendered by engine using `fmt.Sprintf` so we need to escape '%' returned
	// from the AI to prevent it from being interpreted as formatting.
	text = strings.ReplaceAll(text, "%", "%%")

	// messageID is set only for Teams or Slack, for others like Discord, Mattermost, we keep the original formatting
	if messageID != "" {
		text = markdownToSlack(text)
	}
	return api.Message{
		ParentActivityID: messageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Body: api.Body{Plaintext: text},
				},
				Context: []api.ContextItem{
					{Text: "AI-generated content may be incorrect."},
				},
			},
		},
	}
}

func markdownToSlack(text string) string {
	text = mdBold.ReplaceAllString(text, "*$1*")
	// italic not needed
	// single code not needed
	text = mdStrikethrough.ReplaceAllString(text, "~$1~")
	text = mdHeadings.ReplaceAllString(text, "*$1*")
	text = mdImages.ReplaceAllString(text, "<$2|$1>")               // this needs to be first otherwise we will end up with changing ![]() to  !<>
	text = multiLineCodeBlockWithLang.ReplaceAllString(text, "```") // drop the syntax name for multi-line code blocks
	return mdLinks.ReplaceAllString(text, "<$2|$1>")
}

func markdownToTeams(text string) string {
	text = mdHeadings.ReplaceAllString(text, "**$1**")
	text = mdImages.ReplaceAllString(text, "[$1]($2)") // get rid of ! to define it as a link instead of image

	// https://learn.microsoft.com/en-us/microsoftteams/platform/task-modules-and-cards/cards/cards-format?tabs=adaptive-md%2Cdesktop%2Cconnector-html#newlines-for-adaptive-cards
	text += "\n\n~AI-generated content may be incorrect.~\n" // add the warning
	return text
}
