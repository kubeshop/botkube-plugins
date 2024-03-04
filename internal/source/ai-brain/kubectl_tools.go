package aibrain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gookit/color"
	"github.com/kubeshop/botkube/pkg/plugin"
)

const (
	kubectlVersion       = "v1.28.1"
	kubectlBinaryName    = "kubectl"
	kubeconfigEnvVarName = "KUBECONFIG"
)

type kubectlDescribeResourceArgs struct {
	ResourceType  string `json:"resource_type"`
	ResourceName  string `json:"resource_name"`
	AllNamespaces bool   `json:"all-namespaces"`
	Namespace     string `json:"namespace,omitempty"`
}

type kubectlGetResourceArgs struct {
	ResourceType  string `json:"resource_type"`
	ResourceName  string `json:"resource_name"`
	AllNamespaces bool   `json:"all-namespaces"`
	Namespace     string `json:"namespace,omitempty"`
}

type kubectlTopPodsArgs struct {
	ResourceName string `json:"resource_name"`
	Namespace    string `json:"namespace,omitempty"`
}
type kubectlTopNodesArgs struct {
	ResourceName string `json:"resource_name"`
}
type kubectlGetEventsArgs struct {
	Types         string `json:"types"`
	For           string `json:"for"`
	AllNamespaces bool   `json:"all-namespaces"`
	Namespace     string `json:"namespace,omitempty"`
}

type kubectlLogsArgs struct {
	Namespace     string `json:"namespace,omitempty"`
	PodName       string `json:"pod_name,omitempty"`
	ContainerName string `json:"container_name,omitempty"`
}

// KubectlRunner is a runner that executes kubectl commands using a specific kubeconfig file.
type KubectlRunner struct {
	kubeconfigPath string
}

// NewKubectlRunner creates new runner instance.
func NewKubectlRunner(kubeconfigPath string) *KubectlRunner {
	return &KubectlRunner{kubeconfigPath: kubeconfigPath}
}

// DescribeResource executes kubectl describe resource command.
func (k *KubectlRunner) DescribeResource(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlDescribeResourceArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}

	cmd := fmt.Sprintf("describe %s %s %s", args.ResourceType, args.ResourceName, allNsIfPresent(args.AllNamespaces))
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
}

func allNsIfPresent(namespaces bool) string {
	if namespaces {
		return "--all-namespaces"
	}
	return ""
}

// GetResource executes kubectl get resource command.
func (k *KubectlRunner) GetResource(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlGetResourceArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}
	cmd := fmt.Sprintf("get %s %s %s", args.ResourceType, args.ResourceName, allNsIfPresent(args.AllNamespaces))
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
}

// GetEvents executes kubectl get events command.
func (k *KubectlRunner) GetEvents(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlGetEventsArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}
	cmd := fmt.Sprintf("events %s", allNsIfPresent(args.AllNamespaces))

	if args.Types != "" {
		cmd = fmt.Sprintf("%s --type %s", cmd, args.Types)
	}
	if args.For != "" {
		cmd = fmt.Sprintf("%s --for %s", cmd, args.For)
	}
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
}

// TopPods executes kubectl top pods command.
func (k *KubectlRunner) TopPods(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlTopPodsArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}
	cmd := fmt.Sprintf("top pods %s", args.ResourceName)
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
}

// TopNodes executes kubectl top nodes command.
func (k *KubectlRunner) TopNodes(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlTopNodesArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}
	cmd := fmt.Sprintf("top nodes %s", args.ResourceName)
	return k.runKubectlCommand(ctx, cmd, "")
}

// Logs executes kubectl logs command.
func (k *KubectlRunner) Logs(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlLogsArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}
	cmd := fmt.Sprintf("logs %s", args.PodName)
	if args.ContainerName != "" {
		cmd = fmt.Sprintf("%s -c %s", cmd, args.ContainerName)
	}
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
}

func (k *KubectlRunner) runKubectlCommand(ctx context.Context, cmd, ns string) (string, error) {
	envs := map[string]string{
		kubeconfigEnvVarName: k.kubeconfigPath,
	}

	if ns != "" {
		cmd = fmt.Sprintf("-n %s %s ", ns, cmd)
	}

	cmd = fmt.Sprintf("%s %s", kubectlBinaryName, cmd)
	out, err := plugin.ExecuteCommand(ctx, cmd, plugin.ExecuteCommandEnvs(envs))
	if err != nil {
		return "", err
	}

	return color.ClearCode(out.CombinedOutput()), nil
}