package aibrain

import (
	"os/exec"
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

func kubectlDescribePod(args *kubectlDescribePodArgs) ([]byte, error) {
	cmd := exec.Command("kubectl")
	cmd.Args = append(cmd.Args, []string{"-n", args.Namespace, "describe", "pod", args.PodName}...)
	return cmd.Output()
}

func kubectlGetPods(args *kubectlGetPodsArgs) ([]byte, error) {
	cmd := exec.Command("kubectl")
	cmd.Args = append(cmd.Args, []string{"-n", args.Namespace, "get", "pods"}...)

	if args.PodName != "" {
		cmd.Args = append(cmd.Args, args.PodName)
	}

	return cmd.Output()
}

func kubectlGetSecrets(args *kubectlGetSecretsArgs) ([]byte, error) {
	cmd := exec.Command("kubectl")
	cmd.Args = append(cmd.Args, []string{"-n", args.Namespace, "get", "secrets"}...)

	if args.SecretName != "" {
		cmd.Args = append(cmd.Args, args.SecretName)
	}

	return cmd.Output()
}

func kubectlLogs(args *kubectlLogsArgs) ([]byte, error) {
	cmd := exec.Command("kubectl")
	cmd.Args = append(cmd.Args, []string{"-n", args.Namespace, "logs", args.PodName}...)

	if args.ContainerName != "" {
		cmd.Args = append(cmd.Args, "-c", args.ContainerName)
	}

	return cmd.Output()
}
