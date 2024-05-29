Title: Architecture | Botkube

URL Source: https://docs.botkube.io/architecture/

Markdown Content:
This document describes high-level Botkube architecture, all components in the system and interactions between them.

Botkube is split into two main parts:

*   Botkube agent: Botkube binary that serves as a bridge between communication platforms (e.g. Slack, Discord) and Botkube plugins (sources and executors).
*   [Botkube plugins](https://docs.botkube.io/plugins/): The executable binaries that communicate with Botkube agent over an RPC interface. Botkube supports two types of plugins, respectively called Source plugins and Executor plugins.

Components[​](#components "Direct link to Components")
------------------------------------------------------

The following diagram visualizes all main components in the system.

![Image 1: Architecture](https://docs.botkube.io/assets/images/arch-light-5dd32e39675b8833f7bcf6cfe2340542.svg#gh-light-mode-only)![Image 2: Architecture](https://docs.botkube.io/assets/images/arch-dark-d40e372bd6c7930979ab40b08b32ebfb.svg#gh-dark-mode-only)

### Plugin repository[​](#plugin-repository "Direct link to Plugin repository")

A plugin repository is a place where plugin binaries and index file are stored. This repository must be publicly available and supports downloading assets via HTTP(s). Any static file server can be used, for instance: GitHub Pages, `s3`, `gcs`, etc.

### Plugin manager[​](#plugin-manager "Direct link to Plugin manager")

Plugin manager takes care of downloading enabled and bound plugins, running a given plugin binary and maintaining the gRPC connection. Under the hood, the [`go-plugin`](https://github.com/hashicorp/go-plugin/) library is used. Plugin manager is responsible both for the executor and source plugins.

### Plugin executor bridge[​](#plugin-executor-bridge "Direct link to Plugin executor bridge")

Plugin executor bridge is resolving the received Botkube command, calling the respective plugin, and sending back the response to a given communication platform.

### Executor[​](#executor "Direct link to Executor")

Executor is a binary that implements the [executor](https://github.com/kubeshop/botkube/blob/main/proto/executor.proto) Protocol Buffers contract. Executor runs a given command, such as `kubectl` one, and returns the response in a synchronous way.

Streaming command response is not supported. As a result, commands which take a lot of time doesn't work well, as the response won't be sent until the command is finished.

### Plugin source bridge[​](#plugin-source-bridge "Direct link to Plugin source bridge")

Plugin source bridge is dispatching received source events to all configured communication channels.

### Source[​](#source "Direct link to Source")

Source is a binary that implements the [source](https://github.com/kubeshop/botkube/blob/main/proto/source.proto) Protocol Buffers contract. Source starts asynchronous streaming of domain-specific events. For example, streaming Kubernetes events.

### Bot[​](#bot "Direct link to Bot")

Bot represents a bidirectional communication platform. Each bot is responsible for authenticating, managing connections, and providing an interface for receiving and sending messages for a given platform like Slack, Discord, etc. Connection is mostly done via WebSocket.

### Sink[​](#sink "Direct link to Sink")

Sink represents a unidirectional communication platform. Each sink is responsible for authenticating, managing connections, and providing an interface for sending messages for a given platform like Elasticsearch, outgoing webhook, etc.
