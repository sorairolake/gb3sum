# SPDX-FileCopyrightText: 2024 Shun Sakai
#
# SPDX-License-Identifier: GPL-3.0-or-later

alias all := default

# Run default recipe
default: build-debug

# Build `gb3sum` command in debug mode
@build-debug:
    go build

# Build `gb3sum` command in release mode
@build-release:
    go build -buildmode=pie -trimpath -ldflags=-s -mod=readonly -modcacherw

# Remove generated artifacts
@clean:
    go clean

# Run tests
@test:
    go test ./...

# Run `golangci-lint run`
@golangci-lint:
    golangci-lint run -E gofmt,goimports

# Run the formatter
fmt: gofmt goimports

# Run `go fmt`
@gofmt:
    go fmt ./...

# Run `goimports`
@goimports:
    fd -e go -x goimports -w

# Run the linter
lint: vet staticcheck

# Run `go vet`
@vet:
    go vet ./...

# Run `staticcheck`
@staticcheck:
    staticcheck ./...

# Build `gb3sum(1)`
@build-man:
    asciidoctor -b manpage docs/man/man1/gb3sum.1.adoc

# Run the linter for GitHub Actions workflow files
@lint-github-actions:
    actionlint -verbose

# Run the formatter for the README
@fmt-readme:
    npx prettier -w README.md

# Increment the version
@bump part:
    bump-my-version bump {{part}}
