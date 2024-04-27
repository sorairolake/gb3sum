<!--
SPDX-FileCopyrightText: 2024 Shun Sakai

SPDX-License-Identifier: GPL-3.0-or-later
-->

# gb3sum

[![CI][ci-badge]][ci-url]
[![Go Reference][reference-badge]][reference-url]
![Go version][go-version-badge]

**gb3sum** is a command-line utility for printing and checking [BLAKE3]
checksums.

## Installation

### From source

```sh
go install github.com/sorairolake/gb3sum@latest
```

### Via a package manager

| OS    | Package manager | Command                               |
| ----- | --------------- | ------------------------------------- |
| _Any_ | [Homebrew]      | `brew install sorairolake/tap/gb3sum` |

### From binaries

The [release page] contains pre-built binaries for Linux, macOS and Windows.

### How to build

Please see [BUILD.adoc].

## Usage

### Basic usage

#### Print BLAKE3 checksums

```sh
echo "Hello, world!" > foo.txt
gb3sum foo.txt | tee b3sums.txt
```

Output:

```text
94f1675bac4f8bc3c593c63dbf5fe78a0bfda01082af85d5b41a65096db56bff  foo.txt
```

#### Check BLAKE3 checksums

```sh
gb3sum -c b3sums.txt
```

Output:

```text
foo.txt: OK
```

### Generate shell completion

`completion` command generates shell completions to stdout.

The following shells are supported:

- `bash`
- `fish`
- `powershell`
- `zsh`

Example:

```sh
gb3sum completion bash > gb3sum.bash
```

## Command-line options

Please see the following:

- [`gb3sum(1)`]

## Changelog

Please see [CHANGELOG.adoc].

## Contributing

Please see [CONTRIBUTING.adoc].

## License

Copyright &copy; 2024 Shun Sakai (see [AUTHORS.adoc])

1. This program is distributed under the terms of the _GNU General Public
   License v3.0 or later_.
2. Some files are distributed under the terms of the _Creative Commons
   Attribution 4.0 International Public License_.

This project is compliant with version 3.0 of the [_REUSE Specification_]. See
copyright notices of individual files for more details on copyright and
licensing information.

[ci-badge]: https://img.shields.io/github/actions/workflow/status/sorairolake/gb3sum/CI.yaml?branch=develop&style=for-the-badge&logo=github&label=CI
[ci-url]: https://github.com/sorairolake/gb3sum/actions?query=branch%3Adevelop+workflow%3ACI++
[reference-badge]: https://img.shields.io/badge/Go-Reference-steelblue?style=for-the-badge&logo=go
[reference-url]: https://pkg.go.dev/github.com/sorairolake/gb3sum
[go-version-badge]: https://img.shields.io/github/go-mod/go-version/sorairolake/gb3sum?style=for-the-badge&logo=go
[BLAKE3]: https://github.com/BLAKE3-team/BLAKE3
[Homebrew]: https://brew.sh/
[release page]: https://github.com/sorairolake/gb3sum/releases
[BUILD.adoc]: BUILD.adoc
[`gb3sum(1)`]: docs/man/man1/gb3sum.1.adoc
[CHANGELOG.adoc]: CHANGELOG.adoc
[CONTRIBUTING.adoc]: CONTRIBUTING.adoc
[AUTHORS.adoc]: AUTHORS.adoc
[_REUSE Specification_]: https://reuse.software/spec/
