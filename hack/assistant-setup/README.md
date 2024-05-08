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
export ASSISTANT_ENV=dev # dev or prod
npm run start
```

To use your own assistant, modify the `assistantID` variable in the `index.ts` file.

## Development

### Refetch content for file search

The process uses [Jina.AI Reader API](https://github.com/jina-ai/reader) and usually takes a few minutes minutes.
To scrap the content from the latest Botkube website and Botkube Docs, run the following command:

```sh
npm run fetch-content
```

By default, before refetching content starts:

- downloaded files starting with `https://botkube.io/blog/*` and `https://botkube.io/learn/*` prefixes are kept and won't be updated,
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
