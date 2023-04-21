# go-sdk-abstract-factory

## :warning: WARNING: go-sdk-abstract-factory is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `go-sdk-abstract-factory` provides an
[abstract factory](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
that constructs G2Config, G2Configmgr, G2Diagnostic, G2Engine, and G2Product Senzing objects.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-sdk-abstract-factory.svg)](https://pkg.go.dev/github.com/senzing/go-sdk-abstract-factory)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/go-sdk-abstract-factory)](https://goreportcard.com/report/github.com/senzing/go-sdk-abstract-factory)

## Overview

The Senzing `go-sdk-abstract-factory` package creates Senzing objects that each adhere to their respective interfaces:

1. [G2config](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2config)
1. [G2configmgr](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2configmgr)
1. [G2diagnostic](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2diagnostic)
1. [G2engine](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2engine)
1. [G2product](https://pkg.go.dev/github.com/senzing/g2-sdk-go/g2api#G2product)

Depending upon the parameters passed to the factory, one of the following implementations will be returned:

1. [g2-sdk-go-base](https://github.com/Senzing/g2-sdk-go-base) - for calling Senzing Go SDK APIs natively
1. [g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc) - for calling Senzing Go SDK APIs via Senzing gRPC server
1. [g2-sdk-go-mock](https://github.com/Senzing/g2-sdk-go-mock) - for unit testing calls to the Senzing Go SDK APIs

## Use

(TODO:)

## References

- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
