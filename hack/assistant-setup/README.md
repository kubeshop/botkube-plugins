# OpenAI Assistant setup

The tool configures the OpenAI assistant for Botkube AI plugin. It uses documents from the `assets` directory for file search capabilities.

## Toolchain

This project uses [Volta](https://github.com/volta-cli/volta) to manage JS tools. Volta automatically downloads and installs the right Node.js version when you run any of the `node` or `npm` commands.

It is recommended to install it before running the script, to ensure the right Node.js version is used.

## Usage

Navigate to the directory `hack/assistant-setup` and execute one of the following commands.

### Install dependencies

To install all dependencies, run:

```sh
npm install
```

### Start app

```sh
export OPENAI_API_KEY=... # your Open AI API key
export OPENAI_ASSISTANT_ID=... # your Open AI Assistant ID

export OPENAI_PROJECT_ID=... # Optional: your Open AI Project ID
export OPENAI_ORG_ID=... # Optional: your Open AI Organization ID

npm run start
```

## Development

### Refetch content for file search

The process uses [Jina.AI Reader API](https://github.com/jina-ai/reader) and usually takes a few minutes.
A free API token is recommended, to increase the API rate limits, and, in a result, speed up the process. You can get if from the [Jina.AI website](https://jina.ai/reader/).
To scrap the content from the latest Botkube website and Botkube Docs, run the following command:

```sh
export AI_ASSISTANT_FETCH_CONTENT_API_KEY="jina_..." # get it from https://jina.ai/reader/
npm run fetch-content
```

By default, before refetching content starts:

- downloaded files starting with [package.json](package.json)`https://botkube.io/blog/*` and `https://botkube.io/learn/*` prefixes are kept and won't be updated,
- all other content is removed.

To refetch all content, run the following command:

```sh
export PURGE_ALL_CONTENT=true # default: false
npm run fetch-content
```

### Format code

To format code, run:

```sh
npm run format
```
