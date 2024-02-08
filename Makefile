############
# Building #
############

build-plugins: ## Builds all plugins for all defined platforms
	goreleaser build -f .goreleaser.plugin.yaml --rm-dist --snapshot
.PHONY: build-plugins

build-plugins-single: ## Builds all plugins only for current GOOS and GOARCH.
	goreleaser build -f .goreleaser.plugin.yaml --rm-dist --single-target --snapshot
.PHONY: build-plugins-single

build-plugins-archives: ## Builds all plugins for all defined platforms in form of arhcives
	goreleaser release -f .goreleaser.plugin.yaml --rm-dist --snapshot
.PHONY: build-plugins

##############
# Generating #
##############

gen-plugin-index: ## Generate plugins YAML index file.
	go run github.com/kubeshop/botkube/hack -binaries-path "./dist" -use-archive=$(USE_ARCHIVE)
.PHONY: gen-plugin-index

gen-plugins-goreleaser: # Generate Goreleaser config for plugins
	go run github.com/kubeshop/botkube/hack/target/gen-goreleaser

###############
# Developing  #
###############

test: ## Run tests
	go test -v  -race ./...
.PHONY: test

fix-lint-issues: ## Automatically fix lint issues
	go mod tidy
	go mod verify
	golangci-lint run --fix "./..."
.PHONY: fix-lint-issues

#############
# Others    #
#############

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
.PHONY: help
