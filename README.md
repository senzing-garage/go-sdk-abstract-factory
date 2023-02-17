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

## Development

### Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

### Install Senzing C library

Since the Senzing library is a prerequisite, it must be installed first.

1. Verify Senzing C shared objects, configuration, and SDK header files are installed.
    1. `/opt/senzing/g2/lib`
    1. `/opt/senzing/g2/sdk/c`
    1. `/etc/opt/senzing`

1. If not installed, see
   [How to Install Senzing for Go Development](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/install-senzing-for-go-development.md).

### Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=go-sdk-abstract-factory
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

### Test locally using SQLite database

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test-base

    ```

1. **Optional:** View the SQLite database.
   Example:

    ```console
    docker run \
        --env SQLITE_DATABASE=G2C.db \
        --interactive \
        --publish 9174:8080 \
        --rm \
        --tty \
        --volume /tmp/sqlite:/data \
        coleifer/sqlite-web

    ```

   Visit [localhost:9174](http://localhost:9174).

### Test using gRPC server

1. Run a Senzing gRPC server, visit
   [Senzing/servegrpc](https://github.com/Senzing/servegrpc).

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test-grpc

    ```

### Run all test cases

These instructions run testcases for both local and gRPC implementations of the Senzing Go SDK.

1. Run a Senzing gRPC server, visit
   [Senzing/servegrpc](https://github.com/Senzing/servegrpc).

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test

    ```
