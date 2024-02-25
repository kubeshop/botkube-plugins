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

type kubectlDescribePodArgs struct {
	Namespace string `json:"namespace,omitempty"`
	PodName   string `json:"pod_name,omitempty"`
}

type kubectlGetPodsArgs struct {
	Namespace string `json:"namespace,omitempty"`
	PodName   string `json:"pod_name,omitempty"`
}

type kubectlGetSecretsArgs struct {
	Namespace  string `json:"namespace,omitempty"`
	SecretName string `json:"secret_name,omitempty"`
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

// DescribePod executes kubectl describe pod command.
func (k *KubectlRunner) DescribePod(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlDescribePodArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}

	cmd := fmt.Sprintf("describe pod %s", args.PodName)
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
}

// GetPods executes kubectl get pods command.
func (k *KubectlRunner) GetPods(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlGetPodsArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}
	cmd := fmt.Sprintf("get pod %s", args.PodName)
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
}

// GetSecrets executes kubectl get secrets command.
func (k *KubectlRunner) GetSecrets(ctx context.Context, rawArgs []byte) (string, error) {
	var args kubectlGetSecretsArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}
	cmd := fmt.Sprintf("get secrets %s", args.SecretName)
	return k.runKubectlCommand(ctx, cmd, args.Namespace)
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
