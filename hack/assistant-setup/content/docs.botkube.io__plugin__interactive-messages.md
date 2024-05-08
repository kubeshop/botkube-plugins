Title: Interactive messages | Botkube

URL Source: https://docs.botkube.io/plugin/interactive-messages

Markdown Content:
Some of our communication platforms support interactive messages. If interactivity is supported the `IsInteractivitySupported` property under source or execution context is set to **true**.

In this guide, we describe how interactivity applies to messages, and show you how to build ones.

![Image 1: demo](https://docs.botkube.io/assets/images/demo-msg-aa0290bb67438a1f83c756f4b50842df.gif)

Primitives[​](#primitives"DirectlinktoPrimitives")
------------------------------------------------------

You can construct and return interactive messages containing buttons, dropdowns, input text, and more complex formatting. For that reason, we introduced a generic `api.Messages` model.

*   [Buttons](#buttons)
*   [Dropdowns](#dropdowns)
*   [Input text](#input-text-fields)

For all primitives you need to attach the proper command that will be executed by Botkube. The pattern is:

import "github.com/kubeshop/botkube/pkg/api"command :=  api.MessageBotNamePlaceholder + " <plugin_name>" + " <args>"

The `api.MessageBotNamePlaceholder` constant, is our cross-platform placeholder that is replaced by Botkube core with a proper bot name. It's mandatory, otherwise Botkube core will ignore a given command, e.g. a button click.

The `<plugin_name>` and `<args>` should be replaced based on your needs. You replace `<plugin_name>` with your own plugin name, or other plugin name, like `kubectl`. However, if you use other plugin name, a given command won't work if plugin is not enabled.

### Buttons[​](#buttons"DirectlinktoButtons")

You can define a list of buttons under `Section` object. To construct a given button you can use our helper `api.NewMessageButtonBuilder` builder. For example:

const pluginName = "msg"btnBuilder := api.NewMessageButtonBuilder()msg := api.Message{    Sections: []api.Section{        {            Buttons: []api.Button{                btnBuilder.ForCommandWithDescCmd("Run act1", fmt.Sprintf("%s buttons act1", pluginName)),                btnBuilder.ForCommandWithDescCmd("Run act2", fmt.Sprintf("%s buttons act2", pluginName), api.ButtonStylePrimary),                btnBuilder.ForCommandWithDescCmd("Run act3", fmt.Sprintf("%s buttons act3", pluginName), api.ButtonStyleDanger),            },        },    },}

When a button is clicked, Botkube runs an associated command. For **Run act1** it's `msg buttons act1`. If there is a plugin named `msg` and it is enabled on a given channel, it will be called by Botkube with a given command string. As a result, you can parse input command and return proper output.

If you use only `ForCommandWithoutDesc`, all buttons are render in the same line. ![Image 2: btns-desc](https://docs.botkube.io/assets/images/btns-desc-33bccf898c619eca04aa20266804baad.png)

Otherwise, each button is rendered in new line with the description on the left side and button on the right side. ![Image 3: btns-inline](https://docs.botkube.io/assets/images/btns-inline-a7f68fcaac17c49eb65ef94a16ede58e.png)

### Dropdowns[​](#dropdowns"DirectlinktoDropdowns")

You can define dropdowns under `Section` object. You can split options into groups. Optionally, you can define the initial option. It must be included also under `OptionsGroups`.

cmdPrefix := func(cmd string) string {    return fmt.Sprintf("%s %s %s", api.MessageBotNamePlaceholder, pluginName, cmd)}msg := api.Message{    Sections: []api.Section{        {            Selects: api.Selects{                ID: "select-id",                Items: []api.Select{                    {                        Name:    "two-groups",                        Command: cmdPrefix("selects two-groups"),                        OptionGroups: []api.OptionGroup{                            {                                Name: cmdPrefix("selects two-groups/1"),                                Options: []api.OptionItem{                                    {Name: "BAR", Value: "BAR"},                                    {Name: "BAZ", Value: "BAZ"},                                    {Name: "XYZ", Value: "XYZ"},                                },                            },                            {                                Name: cmdPrefix("selects two-groups/2"),                                Options: []api.OptionItem{                                    {Name: "123", Value: "123"},                                    {Name: "456", Value: "456"},                                    {Name: "789", Value: "789"},                                },                            },                        },                        // MUST be defined also under OptionGroups.Options slice.                        InitialOption: &api.OptionItem{                            Name: "789", Value: "789",                        },                    },                },            },        },    },}

When user select a given option, Botkube runs an associated command and appends selected option at the end. For **BAR** it's `msg selects two-gropus BAR`. If there is a plugin named `msg` and it is enabled on a given channel, it will be called by Botkube with a given command string. As a result, you can parse input command and return proper output.

### Input text fields[​](#input-text-fields"DirectlinktoInputtextfields")

msg := api.Message{    PlaintextInputs: []api.LabelInput{        {            Command:          fmt.Sprintf("%s %s input-text", api.MessageBotNamePlaceholder, pluginName),            DispatchedAction: api.DispatchInputActionOnEnter,            Placeholder:      "String pattern to filter by",            Text:             "Filter output",        },    },}

When user types an input string and clicks enter, Botkube runs an associated command and appends input text in quotes. For example, for input **"text"** it's `msg input-text "test"`. If there is a plugin named `msg` and it is enabled on a given channel, it will be called by Botkube with a given command string. As a result, you can parse input command and return proper output.

Message visibility[​](#message-visibility"DirectlinktoMessagevisibility")
------------------------------------------------------------------------------

If the interactivity is enabled, you can change the default message visibility options:

*   The `OnlyVisibleForYou` property allows you to display the message only to the user who clicked a given button or dropdown. It's taken into account only if respond to the interactive message. It cannot be used initial messages that are automatically sent to a given channel.
*   The `ReplaceOriginalMessage` property allows you to replace the message which triggered interactive command. You cannot change `OnlyVisibleForYou` as the message retains its initial visibility.
