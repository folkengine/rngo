[![Build and Test](https://github.com/devplaybooks/.baseline/actions/workflows/CI.yml/badge.svg)](https://github.com/folkengine/rngo/blob/main/.github/workflows/CI.yml)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/folkengine/rngo)

---

## Includes

* [Base Template for Dev Playbook](https://github.com/devplaybooks/.baseline)
* GitHub Actions
  * [golangci-lint](https://github.com/golangci/golangci-lint-action)
    * [A Complete Guide to Linting Go Programs](https://freshman.tech/linting-golang/)

## Resources

* [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

## Steps

* Initiate [go module](https://go.dev/doc/modules/gomod-ref).

```shell
go mod init github.com/folkengine/rngo
```

* Updated Build and Test badge
* Updated Go Report Card badge
* Created [/pkg](https://github.com/golang-standards/project-layout/tree/master/pkg) directory

* Added languages repo as git submodule: 

```shell
git submodule add https://github.com/electronicpanopticon/rng_languages.git pkg/languages
```

