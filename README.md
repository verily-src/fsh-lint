# FSH Lint

[![CI](https://github.com/verily-src/fsh-lint/actions/workflows/ci.yaml/badge.svg)][ci-workflow]
[![GitHub Release](https://img.shields.io/github/v/release/verily-src/fsh-lint)][release]
[![GitHub License](https://img.shields.io/github/license/verily-src/fsh-lint)][license]

A Style and lint checker tool for validating [FHIR Shorthand] (FSH) files.

[FHIR Shorthand]: https://hl7.org/fhir/uv/shorthand/
[license]: https://github.com/verily-src/fsh-lint/blob/main/LICENSE
[release]: https://github.com/verily-src/fsh-lint/releases/latest
[ci-workflow]: https://github.com/verily-src/fsh-lint/actions/workflows/ci.yaml

## Quick Start

* [📦 Installation](#installation)
* [🚀 Usage](#usage)
* [📜 Rules](#rules)
* [🤝 Contributing](#contributing)

## Installation

Prebuilt binaries are available through [GitHub Releases][release], or you can
build from source if you have [Go] installed, with:

```bash
go install github.com/verily-src/fsh-lint@latest
```

[Go]: https://golang.org

## Usage

The linter can be run from the command line. The basic usage is:

```bash
fsh-lint --paths path/to/YourFile.fsh
```

Or over all fsh files in the `input/fsh` directory:
```bash
find input/fsh -name "*.fsh" -exec fsh-lint --paths {} \;
```

Automatic fixes are available for some rules as well, which can be applied with
the `--fix` flag:

```bash
fsh-lint --paths path/to/YourFile.fsh --fix
```

## Rules

Below is the complete list of rules by their rule-id grouped by their category.

### Special Rules

* [required-field-present](docs/rules.md#required-field-present)
* [profile-assignment-present](docs/rules.md#profile-assignment-present)

### Code System Rules

* [code-system-name-format](docs/rules.md#code-system-name-format)
* [code-system-name-matches-filename](docs/rules.md#code-system-name-matches-filename)
* [code-system-name-matches-id](docs/rules.md#code-system-name-matches-id)
* [code-system-name-matches-title](docs/rules.md#code-system-name-matches-title)

### Extension Rules

* [binding-strength-present](docs/rules.md#binding-strength-present)

### Profile Rules

* [binding-strength-present](docs/rules.md#binding-strength-present)
* [profile-name-matches-filename](docs/rules.md#profile-name-matches-filename)
* [profile-name-matches-id](docs/rules.md#profile-name-matches-id)
* [profile-name-matches-title](docs/rules.md#profile-name-matches-title)
* [profile-parent-not-retired](docs/rules.md#profile-parent-not-retired)

### Search Parameter Rules

* [search-parameter-notice](docs/rules.md#search-parameter-notice)

### Value Set Rules

* [value-set-name-format](docs/rules.md#value-set-name-format)
* [value-set-name-matches-filename](docs/rules.md#value-set-name-matches-filename)
* [value-set-name-matches-id](docs/rules.md#value-set-name-matches-id)
* [value-set-name-matches-title](docs/rules.md#value-set-name-matches-title)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for more information.
