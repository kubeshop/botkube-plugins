Title: Interactive output filtering | Botkube

URL Source: https://docs.botkube.io/usage/interactive-output-filtering

Markdown Content:
info

Interactive output filtering is only available for the [Slack integration](https://docs.botkube.io/installation/slack/) that supports interactive messages.

caution

Using the interactive filter input field causes the command to be re-executed. Be careful when using it for read-write commands. This issue is tracked in [botkube#907](https://github.com/kubeshop/botkube/issues/907).

Sometimes you can get long response from a certain command and you may want to filter that to see a subset of the actual result. For each communication platform, you can use the `--filter` flag at the end of your command. To learn more, see the [Flags](https://docs.botkube.io/usage/executor/#filtering-kubectl-output) section.

If you use the [Slack integration](https://docs.botkube.io/installation/slack/) that supports interactive messages, there is another option to handle that: interactive output filtering. Interactivity is achieved via an input action text box where you can add your filter criteria as text and press the **Enter** button.

### Long response filtering[​](#long-response-filtering "Direct link to Long response filtering")

Output Filter input text box will be attached to your command response if response has more than 15 lines. Let's see an example for this situation.

1.  List pods with `@Botkube kubectl get pods -n kube-system`![Image 1: metrics_pods](https://docs.botkube.io/assets/images/output-filtering-get-pods-metrics-18c746eb2031cc45bc74a63389340b03.png)
2.  Let's check the logs of `metrics-server` with `@Botkube kubectl logs -f metrics-server-7cd5fcb6b7-hzvq8 -n kube-system`![Image 2: metrics_logs](https://docs.botkube.io/assets/images/output-filtering-metrics-logs-b6007e41cbfcc6ef727f848a0cdd5808.png)
3.  Notice that Filter Output is attached to response. Type `RequestHeaderAuthRequestController` to filter and press `Enter`. ![Image 3: metrics_filter_logs](https://docs.botkube.io/assets/images/output-filtering-metrics-logs-filter-a6e6cfc2918f182e1a29d16066d47198.png)

Attachment response filtering[​](#attachment-response-filtering "Direct link to Attachment response filtering")
---------------------------------------------------------------------------------------------------------------

Command response is uploaded as text file once the actual size of response message reaches the limit of messaging platform. Let's take a look how Filter Output behaves for this situation.

1.  List the pods with `@Botkube kubectlc get pods -n kube-system`![Image 4: get_pods](https://docs.botkube.io/assets/images/output-filtering-get-pods-21073bfe8502243fe4b90b83b155b99a.png)
2.  Let's check the logs of Traefik with command `@Botkube kubectl logs -f traefik-df4ff85d6-f2mkz -n kube-system`. Notice that, you might have different pod list, so please select a suitable one for your case. ![Image 5: pod_logs](https://docs.botkube.io/assets/images/output-filtering-get-pods-21073bfe8502243fe4b90b83b155b99a.png)
3.  Since the response is big, it is uploaded as file and you can realize a reply thread. Please expand it to see filter output.
4.  Assume we want to see log lines only containing `Configuration` word. Type `Configuration` in the Filter Output input box and click enter. You will see filtered result as a response. ![Image 6: filter_response](https://docs.botkube.io/assets/images/output-filtering-filter-response-5fed09008578a720a302892f2ab81293.png)
