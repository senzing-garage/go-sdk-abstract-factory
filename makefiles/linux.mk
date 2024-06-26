# Makefile extensions for linux.

# -----------------------------------------------------------------------------
# Variables
# -----------------------------------------------------------------------------


# -----------------------------------------------------------------------------
# OS specific targets
# -----------------------------------------------------------------------------

.PHONY: build-osarch-specific
build-osarch-specific: linux/amd64


.PHONY: clean-osarch-specific
clean-osarch-specific:
	@docker rm --force senzing-serve-grpc || true
	@rm -f  $(GOPATH)/bin/$(PROGRAM_NAME) || true
	@rm -f  $(MAKEFILE_DIRECTORY)/coverage.html || true
	@rm -f  $(MAKEFILE_DIRECTORY)/coverage.out || true
	@rm -fr /tmp/sqlite || true
	@rm -fr $(TARGET_DIRECTORY) || true


.PHONY: coverage-osarch-specific
coverage-osarch-specific: export SENZING_LOG_LEVEL=TRACE
coverage-osarch-specific:
	@go test -v -coverprofile=coverage.out -p 1 ./...
	@go tool cover -html="coverage.out" -o coverage.html
	@xdg-open $(MAKEFILE_DIRECTORY)/coverage.html


.PHONY: hello-world-osarch-specific
hello-world-osarch-specific:
	@echo "Hello World, from linux."


.PHONY: run-osarch-specific
run-osarch-specific:
	@go run main.go


.PHONY: setup-osarch-specific
setup-osarch-specific:
	@mkdir /tmp/sqlite
	@cp testdata/sqlite/G2C.db /tmp/sqlite/G2C.db
	@mkdir -p $(TARGET_DIRECTORY)/$(GO_OS)-$(GO_ARCH) || true
	@docker run \
		--detach \
		--env SENZING_TOOLS_DATABASE_URL=sqlite3://na:na@/tmp/sqlite/G2C.db \
		--env SENZING_TOOLS_ENABLE_ALL=true \
		--name senzing-serve-grpc \
		--publish 8261:8261 \
		--rm \
		senzing/serve-grpc
	@echo "senzing/senzing-tools server-grpc running in background."


.PHONY: test-osarch-specific
test-osarch-specific:
	@go test -json -v -p 1 ./... 2>&1 | tee /tmp/gotest.log | gotestfmt

# -----------------------------------------------------------------------------
# Makefile targets supported only by this platform.
# -----------------------------------------------------------------------------

.PHONY: only-linux
only-linux:
	@echo "Only linux has this Makefile target."
