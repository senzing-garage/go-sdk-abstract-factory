# go-sdk-abstract-factory

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: go-sdk-abstract-factory is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `go-sdk-abstract-factory` provides an
[abstract factory](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
creator that construct a
[SzAbstractFactory](https://github.com/senzing-garage/sz-sdk-go/blob/main/sz/main.go)
Senzing objects.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/go-sdk-abstract-factory.svg)](https://pkg.go.dev/github.com/senzing-garage/go-sdk-abstract-factory)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/go-sdk-abstract-factory)](https://goreportcard.com/report/github.com/senzing-garage/go-sdk-abstract-factory)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/go-sdk-abstract-factory/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/go-sdk-abstract-factory/actions/workflows/go-test-windows.yaml)

## Overview

Depending upon the `szfactorycreator.CreateXxxxAbstractFactory(...)` method called,
one of the following AbstractFactory implementations will be returned:

1. `CreateCoreAbstractFactory(...)` returns an AbstractFactory from
   [sz-sdk-go-core](https://github.com/senzing-garage/sz-sdk-go-core) for calling Senzing Go SDK APIs natively
1. `CreateGrpcAbstractFactory(...)` returns an AbstractFactory from
   [sz-sdk-go-grpc](https://github.com/senzing-garage/sz-sdk-go-grpc) for calling Senzing Go SDK APIs via Senzing gRPC server
1. `CreateMockAbstractFactory(...)` returns an AbstractFactory from
   [sz-sdk-go-mock](https://github.com/senzing-garage/sz-sdk-go-mock) for unit testing calls to the Senzing Go SDK APIs

## Use

(TODO:)

## References

1. [API documentation](https://pkg.go.dev/github.com/senzing-garage/go-sdk-abstract-factory)
1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
