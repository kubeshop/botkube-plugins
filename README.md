# Botkube Cloud Plugins

This repository shows Botkube Cloud plugins.

## Requirements

- [Go](https://golang.org/doc/install) >= 1.21
- [GoReleaser](https://goreleaser.com/) >= 1.21
- [`golangci-lint`](https://golangci-lint.run/) >= 1.55

## Development

1. Clone the repository.
2. Follow the [local testing guide](https://docs.botkube.io/plugin/local-testing).

## Release

The latest plugins from `main` are released automatically to the bucket `botkube-cloud-plugins-latest`.
To release a new version of the plugins:

1. Navigate to the [`Trigger release` GitHub Actions workflow](https://github.com/kubeshop/botkube-cloud-plugins/actions/workflows/release.yml).
1. Provide the target version in the `version` input. Do not forget about the `v` prefix.

   For example, if you want to release version `1.2.3`, you should provide `version: v1.2.3`.

1. Trigger the pipeline.
