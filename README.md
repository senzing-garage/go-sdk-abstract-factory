# go-sdk-abstract-factory

## :warning: WARNING: go-sdk-abstract-factory is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing go-sdk-abstract-factory provides an
[abstract factory](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
that constructs G2Config, G2Configmgr, G2Diagnostic, G2Engine, and G2Product Senzing objects.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-sdk-abstract-factory.svg)](https://pkg.go.dev/github.com/senzing/go-sdk-abstract-factory)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/go-sdk-abstract-factory)](https://goreportcard.com/report/github.com/senzing/go-sdk-abstract-factory)
[![go-test.yaml](https://github.com/Senzing/go-sdk-abstract-factory/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/go-sdk-abstract-factory/actions/workflows/go-test.yaml)

## Overview

The Senzing go-sdk-abstract-factory package creates Senzing objects that each adhere to their respective interfaces:

1. [G2config interface](https://github.com/Senzing/g2-sdk-go/blob/main/g2config/main.go)
1. [G2configmgr interface](https://github.com/Senzing/g2-sdk-go/blob/main/g2configmgr/main.go)
1. [G2diagnostic interface](https://github.com/Senzing/g2-sdk-go/blob/main/g2diagnostic/main.go)
1. [G2engine interface](https://github.com/Senzing/g2-sdk-go/blob/main/g2engine/main.go)
1. [G2Product interface](https://github.com/Senzing/g2-sdk-go/blob/main/g2product/main.go)

Depending upon the parameters passed to the factory, the underlying implementation may be:

1. [g2-sdk-go](https://github.com/Senzing/g2-sdk-go) which talks to the native Senzing SDK C API
1. [g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc) which communicates with a Senzing gRPC server

## Developing with go-sdk-abstract-factory

### Install Senzing library

If using an existing Senzing SDK gRPC server,
[g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc),
this step is not required.

If configuring the go-sdk-abstract factory to use the local
[g2-sdk-go](https://github.com/Senzing/g2-sdk-go)
implementation,  the Senzing library is a prerequisite, it must be installed first.
This can be done by installing the Senzing package using `apt`, `yum`,
or a technique using Docker containers.
Once complete, the Senzing library will be installed in the `/opt/senzing` directory.
This is important as the compiling of the code expects Senzing to be in `/opt/senzing`.

- Using `apt`:

    ```console
    wget https://senzing-production-apt.s3.amazonaws.com/senzingrepo_1.0.0-1_amd64.deb
    sudo apt install ./senzingrepo_1.0.0-1_amd64.deb
    sudo apt update
    sudo apt install senzingapi

    ```

- Using `yum`:

    ```console
    sudo yum install https://senzing-production-yum.s3.amazonaws.com/senzingrepo-1.0.0-1.x86_64.rpm
    sudo yum install senzingapi

    ```

- Using Docker:

  This technique can be handy if you are using MacOS or Windows and cross-compiling.

    1. Build Senzing installer.

        ```console
        curl -X GET \
            --output /tmp/senzing-versions-stable.sh \
            https://raw.githubusercontent.com/Senzing/knowledge-base/main/lists/senzing-versions-stable.sh
        source /tmp/senzing-versions-stable.sh

        sudo docker build \
            --build-arg SENZING_ACCEPT_EULA=I_ACCEPT_THE_SENZING_EULA \
            --build-arg SENZING_APT_INSTALL_PACKAGE=senzingapi=${SENZING_VERSION_SENZINGAPI_BUILD} \
            --build-arg SENZING_DATA_VERSION=${SENZING_VERSION_SENZINGDATA} \
            --no-cache \
            --tag senzing/installer:${SENZING_VERSION_SENZINGAPI} \
            https://github.com/senzing/docker-installer.git#main

        ```

    1. Install Senzing.

        ```console
            curl -X GET \
                --output /tmp/senzing-versions-stable.sh \
                https://raw.githubusercontent.com/Senzing/knowledge-base/main/lists/senzing-versions-stable.sh
            source /tmp/senzing-versions-stable.sh

            sudo rm -rf /opt/senzing
            sudo mkdir -p /opt/senzing

            sudo docker run \
                --rm \
                --user 0 \
                --volume /opt/senzing:/opt/senzing \
                senzing/installer:${SENZING_VERSION_SENZINGAPI}

        ```

### Configure Senzing

1. Move the "versioned" Senzing data to the system location.
   Example:

    ```console
      sudo mv /opt/senzing/data/3.0.0/* /opt/senzing/data/

    ```

1. Create initial configuration.
   Example:

    ```console
    export SENZING_ETC_FILES=( \
        "cfgVariant.json" \
        "customGn.txt" \
        "customOn.txt" \
        "customSn.txt" \
        "defaultGNRCP.config" \
        "g2config.json" \
        "G2Module.ini" \
        "stb.config" \
    )

    sudo mkdir /etc/opt/senzing
    for SENZING_ETC_FILE in ${SENZING_ETC_FILES[@]}; \
    do \
        sudo --preserve-env cp /opt/senzing/g2/resources/templates/${SENZING_ETC_FILE} /etc/opt/senzing/${SENZING_ETC_FILE}
    done
    ```

### Test locally using SQLite database

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=go-sdk-abstract-factory
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test-local

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

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=go-sdk-abstract-factory
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test-grpc

    ```

### Run all test cases

These instructions run testcases for both local and gRPC implementations of the Senzing Go SDK.
A Senzing SDK gRPC server,
[g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc),
listening on `localhost:8258` is required.

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=go-sdk-abstract-factory
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test

    ```
