package otelx

import (
	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"
)

// Init initialises opentelemetry with honeycomb exporter. Returns a shutdown func.
func Init(apiKey, serviceName string, sampleRate int, version string) (func(), error) {
	return otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithSpanProcessor(honeycomb.NewBaggageSpanProcessor()),
		otelconfig.WithHeaders(map[string]string{"X-Honeycomb-Team": apiKey}),
		otelconfig.WithServiceName(serviceName),
		otelconfig.WithSampler(honeycomb.NewDeterministicSampler(sampleRate)),
		otelconfig.WithServiceVersion(version),
	)
}
