Title: Botkube v0.16 Release Notes

URL Source: https://botkube.io/blog/botkube-v016-release-notes

Published Time: Nov 28, 2022

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

Botkube v0.16.0 adds new event filter criteria allowing more granular event tuning, and actions, which automatically run commands when specified event types are received.

### Table of Contents

*   [Filters](https://botkube.io/blog/botkube-v016-release-notes#filters)
*   [Actions](https://botkube.io/blog/botkube-v016-release-notes#actions)
*   [Legacy Slack App Deprecation](https://botkube.io/blog/botkube-v016-release-notes#legacy-slack-app-deprecation)
*   [Bug Fixes](https://botkube.io/blog/botkube-v016-release-notes#bug-fixes)
*   [Feedback - We’d Love to Hear From You](https://botkube.io/blog/botkube-v016-release-notes#feedback-we-d-love-to-hear-from-you)

#### Start Using Botkube AI-Powered Assistant Today

The latest version of Botkube is here, v0.16.0. Like peanut butter and chocolate, we've brought together two great parts of Botkube to make your life working with Kubernetes even easier. Botkube is the most modern [ChatOps tool for Kubernetes](https://botkube.io/)!

We've addressed two things that have been frequently requested by users. First, we've added a more configurable filter system for events, meaning you can better filter what kind of events you want to see. Second, we've added actions, which can be configured to run commands automatically when an event is received and the command output is included with the event. These features work behind the scenes in the Botkube engine and are not specific to any communication platform.

_If you have any feature requests or bug reports, reach out to us on [Slack](https://join.botkube.io/) or submit a [Github issue](https://github.com/kubeshop/botkube/issues)_

Filters
-------

Botkube now has additional event and resource constraints that you can use to filter events when you're defining sources. Sources define the event stream you want to watch and are then bound to a communication channel that will receive all matching events.

In earlier versions of Botkube you could create sources based on Kubernetes resource type (e.g. Pod, ConfigMap, etc.), Namespace, and event type (e.g. Create, Delete, Error). This is useful for getting all events of a certain type in a namespace, for example, but in some cases still results in a lot of events that aren't relevant to all users in a channel. In addition, this doesn't provide the granular level of control needed to automatically act on events (see Actions below!).

In Botkube v0.16.0 we've added the following [event filter criteria](https://docs.botkube.io/next/configuration/source#kubernetes-resource-events):

*   **labels and annotations**: These filters simply have to (exactly) match the labels or annotations applied to a resource in order for the event to be logged.
*   **name**: This is a Regex field to match resource names. Events pertaining to any resource that matches the specified pattern will be logged.
*   **event.reason and event.message**: These are Regex fields to match the event reason and event message. Events matching the specified pattern will be logged.

The new filter criteria have lots of great uses. You may have a suspect resource like a flaky pod, for example. Now you can create a "flaky pod" Slack channel and send error events pertaining to that suspect pod with the specific failure reason to that Slack channel and monitor when and how often those failures occur.

They also enable Botkube actions!

Actions
-------

Along with the new, more granular filter options, we have introduced [actions](https://docs.botkube.io/next/configuration/action) to Botkube. Actions allow you to automatically trigger a specific command when a matching event occurs. You can have more context for an event delivered directly to your communication channels with no additional manual work.

Actions consist of two parts, the command and the source binding. The command is the specific executor command that will be run by the action. In v0.16.0 only the kubectl executor is available but more will be coming soon. The command itself is defined using the [Go template](https://golang.org/pkg/text/template/) syntax and allows you to pass the event metadata in order to build the command you need. The source binding specifies the event source that will trigger the action and uses the same sources that can be bound to communication channels. These sources use the event filters discussed earlier.

When a source is bound to both an action and a communication channel, the results of the action command are delivered to the communication channel along with the original event. Actions can be used to get information using the default read-only kubectl configuration. If you want, you can even configure the executor to allow other actions beyond those that are read-only. You could have Botkube restart a hang pod automatically, for example, and the results will be sent to the communication channel along with the event.

![Image 2: Config Map created in Kubernetes Cluster from Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6384b39243714f0c359b33fb_vpk-kQTm3FfWa0WOV6yI7kl5Zg8Fv4kwQDDPDycWIDWbSbuoapqDwW-z95PUaA8qDINbpb92Z1k-gTtYC-qNf83CLnWJ_1nFQ4BJbniN-gywZtU2siAcyz8jIpvwKuGn90vIGCbnChTj7eIV21j7dGlO8_dAFNhtnOETwIPB4EU7-j5EX8S9G3qw1FQeCw.png)

A Botkube action at work

There are two preconfigured actions included with Botkube to help you get started quickly, they are disabled by default:

*   **describe-created resource**: This action runs the kubectl describe command against the resource associated with an event. It is bound to the default k8s-create-events source by default. When enabled, the command will be run each time a resource is created and the event is logged by the k8s-create-events source. The results of the kubectl describe command will then be attached to the event and delivered to any channel to which the k8s-create-events source is bound.
*   **show-logs-on-error**: This action runs kubectl logs against a resource. It is bound to the k8s-err-with-logs-events source by default. Any event captured by this source (triggered by error event types) will include the output of the kubectl logs command and be delivered to any communication channel to which it is bound.

Legacy Slack App Deprecation
----------------------------

With the release of the [new Socket-mode Slack app](https://docs.botkube.io/installation/socketslack/) in [Botkube v0.14.0](https://botkube.io/blog/botkube-v014-release-notes#new-botkube-slack-app), the legacy Slack app is being deprecated. In the next version (v0.17.0), it will no longer be supported in the Botkube app and Botkube will need to be configured to use the new Slack app.

This change is happening because Botkube requires the socket-mode Slack app for the new [interactivity features](https://botkube.io/blog/botkube-v015-release-notes#interactive-kubectl) we introduced in Botkube v0.15.0. We plan to put all of our development efforts into this new app so we can continue to build great interactive features. In addition, Slack is recommending the migration from the legacy app type to socket mode as it introduces more granular permissions for bot users and OAuth tokens. This means you have more control over what Botkube and other Slack apps can do in your Slack workspace.

Slack will be unapproving the legacy Botkube Slack app in their Slack App Directory on December 21, 2022. After this date you will still be able to install the legacy Slack app using the Add to Slack button in the [legacy Slack installation](https://docs.botkube.io/0.15/installation/slack/) guide but you will receive a warning that the app is not authorized by the Slack App Directory. You will need to accept this warning to proceed if you want to keep using the legacy Slack app with Botkube v0.16 or earlier.

Bug Fixes
---------

We have also fixed a couple of small bugs in Botkube that were reported to us by users and found by us. You can see the full list of changes in the [changelog](https://github.com/kubeshop/botkube/releases/tag/v0.16.0).

We appreciate your bug reports and pull requests! If you notice a bug in Botkube, please report it as a [GitHub issue](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

Feedback - We’d Love to Hear From You
-------------------------------------

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the Botkube [GitHub issues](https://github.com/kubeshop/botkube/issues), Botkube [Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
