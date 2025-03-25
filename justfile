# SPDX-FileCopyrightText: 2024 Shun Sakai
#
# SPDX-License-Identifier: GPL-3.0-or-later

alias build := build-debug
alias fmt := golangci-lint-fmt
alias lint := golangci-lint-run

# Run default recipe
_default:
    just -l

# Build `gb3sum` command in debug mode
build-debug $CGO_ENABLED="0":
    go build

# Build `gb3sum` command in release mode
build-release $CGO_ENABLED="0":
    go build -ldflags="-s -w" -trimpath

# Remove generated artifacts
clean:
    go clean

# Run tests
test:
    go test ./...

# Run `golangci-lint`
golangci-lint: golangci-lint-fmt golangci-lint-run

# Run the formatter
golangci-lint-fmt:
    go tool golangci-lint fmt

# Run the linter
golangci-lint-run:
    go tool golangci-lint run

# Build `gb3sum(1)`
build-man:
    asciidoctor -b manpage docs/man/man1/gb3sum.1.adoc

# Run the linter for GitHub Actions workflow files
lint-github-actions:
    actionlint -verbose

# Run the formatter for the README
fmt-readme:
    npx prettier -w README.md

# Increment the version
bump part:
    bump-my-version bump {{ part }}
