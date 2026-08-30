# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Renamed module from `qumo-lb` to `quic-lb` to match repo name.
- Root package uses re-export pattern; implementation lives in `internal/`.

### Fixed
- CI: `go test ./...` no longer fails on repos with no Go packages (added root `doc.go`).
- CI: combined redundant double test run into single `-v -coverprofile` pass in `go.yml`.
- CI: race job now uses `-short` to skip slow integration tests.
- CI: lint workflow no longer does a full-history clone.
- CI: pinned `golangci_lint_version` to `v2` and use `go_version_file` in `lint.yml`.
- CI: benchmark workflow only installs `benchstat` when Go packages changed.
- CI: require-changelog workflow no longer triggers on `edited` events nor clones full history.
