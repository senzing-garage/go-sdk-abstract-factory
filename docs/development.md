# go-sdk-abstract-factory development

## Install Go

1. See Go's [Download and install](https://go.dev/doc/install)

## Install Senzing C library

Since the Senzing library is a prerequisite, it must be installed first.

1. Verify Senzing C shared objects, configuration, and SDK header files are installed.
    1. `/opt/senzing/g2/lib`
    1. `/opt/senzing/g2/sdk/c`
    1. `/etc/opt/senzing`

1. If not installed, see
   [How to Install Senzing for Go Development](https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/install-senzing-for-go-development.md).

## Install Git repository

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=go-sdk-abstract-factory
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/senzing-garage/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

## Test using SQLite database

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make clean test

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
   [Senzing/servegrpc](https://github.com/senzing-garage/servegrpc).

    1. Optionally, initialize a SqlLite database.
       Example:

        ```console
        export SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db
        senzing-tools init-database

        ```

    1. Start Senzing gRPC server.
       Example:

        ```console
        senzing-tools serve-grpc --enable-all

        ```

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test-grpc

    ```

### Run all test cases

These instructions run testcases for both local and gRPC implementations of the Senzing Go SDK.

1. Run a Senzing gRPC server, visit
   [Senzing/servegrpc](https://github.com/senzing-garage/servegrpc).

    1. Optionally, initialize a SqlLite database.
       Example:

        ```console
        export SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db
        senzing-tools init-database

        ```

    1. Start Senzing gRPC server.
       Example:

        ```console
        senzing-tools serve-grpc --enable-all

        ```

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test

    ```
