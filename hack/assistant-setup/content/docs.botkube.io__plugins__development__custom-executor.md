Title: Custom executor | Botkube

URL Source: https://docs.botkube.io/plugins/development/custom-executor

Markdown Content:
You can extend Botkube functionality by writing custom executor plugin. An executor allows you to run a given command, such as `kubectl`, directly in the chat window of each communication platform For example.

Plugin executor is a binary that implements the [executor](https://github.com/kubeshop/botkube/blob/main/proto/executor.proto) Protocol Buffers contract.

Goal[​](https://docs.botkube.io/plugins/development/custom-executor/#goal "Direct link to Goal")
------------------------------------------------------------------------------------------------

This tutorial shows you how to build a custom `echo` executor that responds with a command that was specified by the user in a chat window.

![Image 1: echo-demo](https://docs.botkube.io/assets/images/echo-demo-ff35ba0e3555968d781b9a2ac6961311.gif)

For a final implementation, see the [Botkube template repository](https://docs.botkube.io/plugins/development/quick-start).

Prerequisites[​](https://docs.botkube.io/plugins/development/custom-executor/#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------------------------------------------------------------------

*   Basic understanding of the Go language.
*   [Go](https://golang.org/doc/install) at least 1.18.
    *   See [go.mod](https://github.com/kubeshop/botkube/blob/main/go.mod#L1) for the recommended version used by Botkube team.
*   [GoReleaser](https://goreleaser.com/) at least 1.13.

### Develop plugin business logic[​](https://docs.botkube.io/plugins/development/custom-executor/#develop-plugin-business-logic "Direct link to Develop plugin business logic")

1.  Create an executor plugin directory:
    
    ```
    mkdir botkube-plugins && cd botkube-plugins
    ```
    
2.  Initialize the Go module:
    
    ```
    go mod init botkube-plugins
    ```
    
3.  Create the `main.go` file for the `echo` executor with the following template:
    
    ```
    cat << EOF > main.gopackage mainimport (  "context"  "fmt"  "github.com/MakeNowJust/heredoc"  "github.com/hashicorp/go-plugin"  "github.com/kubeshop/botkube/pkg/api"  "github.com/kubeshop/botkube/pkg/api/executor")// EchoExecutor implements the Botkube executor plugin interface.type EchoExecutor struct{}func main() {  executor.Serve(map[string]plugin.Plugin{    "echo": &executor.Plugin{      Executor: &EchoExecutor{},    },  })}EOF
    ```
    
    This template code imports required packages and registers `EchoExecutor` as the gRPC plugin. At this stage, the `EchoExecutor` service doesn't implement the required [Protocol Buffers](https://github.com/kubeshop/botkube/blob/main/proto/executor.proto) contract. We will add the required methods in the next steps.
    
4.  Download imported dependencies:
    
5.  Add the required `Metadata` method:
    
    ```
    // Metadata returns details about the Echo plugin.func (*EchoExecutor) Metadata(context.Context) (api.MetadataOutput, error) {  return api.MetadataOutput{    Version:     "1.0.0",    Description: "Echo sends back the command that was specified.",    JSONSchema: api.JSONSchema{    Value: heredoc.Doc(`{       "$schema": "http://json-schema.org/draft-04/schema#",       "title": "echo",       "description": "Example echo plugin",       "type": "object",       "properties": {         "formatOptions": {           "description": "Options to format echoed string",           "type": "array",           "items": {             "type": "string",             "enum": [ "bold", "italic" ]           }         }       },       "additionalProperties": false     }`),    },  }, nil}
    ```
    
    The `Metadata` method returns basic information about your plugin. This data is used when the plugin index is generated in an automated way. You will learn more about that in the next steps.
    
    Ąs a part of the `Metadata` method, you can define the plugin dependencies. To learn more about them, see the [Dependencies](https://docs.botkube.io/plugins/development/dependencies) document.
    
6.  Add the required `Execute` method:
    
    ```
    // Execute returns a given command as a response.func (*EchoExecutor) Execute(_ context.Context, in executor.ExecuteInput) (executor.ExecuteOutput, error) {	return executor.ExecuteOutput{		Message: api.NewCodeBlockMessage(response, true),	}, nil}
    ```
    
    The `Execute` method is the heart of your executor plugin. This method runs your business logic and returns the execution output. Next, the Botkube core sends back the response to a given communication platform. If the communication platform supports interactivity, you can construct and return interactive messages containing buttons, dropdowns, input text, and more complex formatting. To learn more about it, see the [Interactive Messages](https://docs.botkube.io/plugins/development/interactive-messages) guide.
    
    For each `Execute` method call, Botkube attaches the list of associated configurations. You will learn more about that in the [**Passing configuration to your plugin**](https://docs.botkube.io/plugins/development/custom-executor/#passing-configuration-to-your-plugin) section.
    
7.  Add the required `Help` method:
    
    ```
    // Help returns help messagefunc (EchoExecutor) Help(context.Context) (api.Message, error) {	btnBuilder := api.NewMessageButtonBuilder()	return api.Message{		Sections: []api.Section{			{				Base: api.Base{					Header:      "Run `echo` commands",					Description: description,				},				Buttons: []api.Button{					btnBuilder.ForCommandWithDescCmd("Run", "echo 'hello world'"),				},			},		},	}, nil}
    ```
    
    You can use `api.NewCodeBlockMessage` or `api.NewPlaintextMessage` helper functions, or construct your own message.
    

Build plugin binaries[​](https://docs.botkube.io/plugins/development/custom-executor/#build-plugin-binaries "Direct link to Build plugin binaries")
---------------------------------------------------------------------------------------------------------------------------------------------------

Now it's time to build your plugin. For that purpose, we will use GoReleaser. It simplifies building Go binaries for different architectures. The important thing is to produce the binaries for the architecture of the host platform where Botkube is running. Adjust the `goos`, `goarch`, and `goarm` properties based on your needs.

note

Instead of GoReleaser, you can use another tool of your choice.

1.  Create the GoReleaser configuration file:
    
    ```
    cat << EOF > .goreleaser.yamlproject_name: botkube-pluginsbefore:  hooks:    - go mod downloadbuilds:  - id: echo    binary: executor_echo_{{ .Os }}_{{ .Arch }}    no_unique_dist_dir: true    env:      - CGO_ENABLED=0    goos:      - linux      - darwin    goarch:      - amd64      - arm64    goarm:      - 7snapshot:  name_template: 'v{{ .Version }}'EOF
    ```
    
2.  Build the binaries:
    
    ```
    goreleaser build --rm-dist --snapshot
    ```
    

Congrats! You just created your first Botkube executor plugin! 🎉

Now it's time to [test it locally with Botkube](https://docs.botkube.io/plugins/development/local-testing). Once you're done testing, see how to [distribute it](https://docs.botkube.io/plugins/development/repo).

Passing configuration to your plugin[​](https://docs.botkube.io/plugins/development/custom-executor/#passing-configuration-to-your-plugin "Direct link to Passing configuration to your plugin")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Sometimes your executor plugin requires a configuration specified by the end-user. Botkube supports such requirement and provides an option to specify plugin configuration under `config`. An example Botkube configuration looks like this:

```
communications:  "default-group":    socketSlack:      channels:        "default":          name: "all-teams"          bindings:            executors:              - echo-team-a              - echo-team-bexecutors:  "echo-team-a": # executor configuration group name    botkube/echo:      enabled: true      config:        formatOptions: ["italic"]  "echo-team-b": # executor configuration group name    botkube/echo:      enabled: true      config:        formatOptions: ["bold"]
```

This means that two different `botkube/echo` plugin configurations were bound to the `all-teams` Slack channel. Under `executor.ExecuteInput{}.Configs`, you will find the list of configurations in the YAML format as specified under the `config` property for each bound and enabled executor. The order of the configuration is the same as specified under the `bindings.executors` property. It's up to the plugin author to merge the passed configurations. You can use our helper function from the plugin extension package (`pluginx`).

```
import (	"context"	"github.com/kubeshop/botkube/pkg/api/executor"	"github.com/kubeshop/botkube/pkg/pluginx")// Config holds the executor configuration.type Config struct {	FormatOptions []string `yaml:"options,omitempty"`}func (EchoExecutor) Execute(_ context.Context, in executor.ExecuteInput) (executor.ExecuteOutput, error) {	var cfg Config	err := pluginx.MergeExecutorConfigs(in.Configs, &cfg)	if err != nil {		return executor.ExecuteOutput{}, err	}	// rest logic}
```

caution

Botkube starts only one process of a given executor plugin, and the list of configuration YAMLs can be different per gRPC call, so you shouldn't cache the merged configuration.

Notes[​](https://docs.botkube.io/plugins/development/custom-executor/#notes "Direct link to Notes")
---------------------------------------------------------------------------------------------------

*   Streaming command response is not supported. As a result, commands like `helm install --wait` doesn't work well, as the response won't be sent until the command finishes. It's recommended to return the response as soon as possible.
*   The interactive message is not yet supported.
