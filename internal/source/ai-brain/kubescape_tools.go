package aibrain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gookit/color"
	"github.com/kubeshop/botkube/pkg/plugin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	kubescapeBinaryName = "kubescape"
	kubescapeVersion    = "v3.0.11"
)

type KubescapeRunner struct {
	tracer         trace.Tracer
	kubeconfigPath string
}

func NewKubescapeRunner(kubeconfigPath string, tracer trace.Tracer) *KubescapeRunner {
	return &KubescapeRunner{
		kubeconfigPath: kubeconfigPath,
		tracer:         tracer,
	}
}

func (k *KubescapeRunner) ScanCluster(ctx context.Context, _ []byte, _ *Payload) (string, error) {
	ctx, span := k.tracer.Start(ctx, "aibrain.KubescapeRunner.ScanCluster")
	defer span.End()

	cmd := "scan --verbose"
	return k.runKubescapeCommand(ctx, cmd)
}

type scanWorkloadArgs struct {
	ResourceKind string `json:"resource_kind"`
	ResourceName string `json:"resource_name"`
	Namespace    string `json:"namespace,omitempty"`
}

func (k *KubescapeRunner) ScanWorkload(ctx context.Context, rawArgs []byte, _ *Payload) (string, error) {
	ctx, span := k.tracer.Start(ctx, "aibrain.KubescapeRunner.ScanWorkload")
	defer span.End()

	span.SetAttributes(attribute.String("kubescapeRunner.args", string(rawArgs)))

	var args scanWorkloadArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}

	cmd := fmt.Sprintf("scan -v workload %s/%s", args.ResourceKind, args.ResourceName)
	if args.Namespace != "" {
		cmd += fmt.Sprintf(" --namespace %s", args.Namespace)
	}
	return k.runKubescapeCommand(ctx, cmd)
}

type scanImageArgs struct {
	Image string `json:"image"`
}

func (k *KubescapeRunner) ScanImage(ctx context.Context, rawArgs []byte, _ *Payload) (string, error) {
	ctx, span := k.tracer.Start(ctx, "aibrain.KubescapeRunner.ScanImage")
	defer span.End()

	span.SetAttributes(attribute.String("kubescapeRunner.args", string(rawArgs)))

	var args scanImageArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}

	cmd := fmt.Sprintf("scan -v image %s", args.Image)
	return k.runKubescapeCommand(ctx, cmd)
}

type scanControlArgs struct {
	Control string `json:"control"`
}

func (k *KubescapeRunner) ScanControl(ctx context.Context, rawArgs []byte, _ *Payload) (string, error) {
	ctx, span := k.tracer.Start(ctx, "aibrain.KubescapeRunner.ScanControl")
	defer span.End()

	span.SetAttributes(attribute.String("kubescapeRunner.args", string(rawArgs)))

	var args scanControlArgs
	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return "", fmt.Errorf("invalid arguments: %w", err)
	}

	cmd := fmt.Sprintf("scan -v control %s", args.Control)
	return k.runKubescapeCommand(ctx, cmd)
}

func (k *KubescapeRunner) runKubescapeCommand(ctx context.Context, cmd string) (string, error) {
	ctx, span := k.tracer.Start(ctx, "aibrain.KubectlRunner.runKubescapeCommand")
	defer span.End()

	envs := map[string]string{
		kubeconfigEnvVarName: k.kubeconfigPath,
	}

	cmd = fmt.Sprintf("%s %s", kubescapeBinaryName, cmd)
	span.SetAttributes(attribute.String("kubectlRunner.cmd", cmd))

	out, err := plugin.ExecuteCommand(ctx, cmd, plugin.ExecuteCommandEnvs(envs))

	if out.ExitCode != 0 {
		return fmt.Sprintf("kubectl command failed: %s", color.ClearCode(out.CombinedOutput())), nil
	}

	if err != nil {
		return "", err
	}

	return color.ClearCode(out.CombinedOutput()), nil
}
