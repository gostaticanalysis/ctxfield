# ctxfield [![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godoc] [![Travis](https://img.shields.io/travis/gostaticanalysis/ctxfield.svg?style=flat-square)][travis] [![Go Report Card](https://goreportcard.com/badge/github.com/gostaticanalysis/ctxfield)](https://goreportcard.com/report/github.com/gostaticanalysis/ctxfield) [![codecov](https://codecov.io/gh/gostaticanalysis/ctxfield/branch/master/graph/badge.svg)](https://codecov.io/gh/gostaticanalysis/ctxfield)

`ctxfield` reports `context.Context` which belongs to a struct as a field.
`ctxfield` ignores a struct which implements an interface.

## Install

```sh
$ go get github.com/gostaticanalysis/ctxfield/cmd/ctxfield
```

## Usage

```sh
$ go vet -vettool=`which ctxfield` pkgname
```

<!-- links -->
[godoc]: http://godoc.org/github.com/gostaticanalysis/ctxfield
[travis]: https://travis-ci.org/gostaticanalysis/ctxfield
