# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], [markdownlint],
and this project adheres to [Semantic Versioning].

## [Unreleased]

-

## [0.9.11] - 2025-06-18

### Changed in 0.9.11

- Update dependencies

## [0.9.10] - 2025-06-03

### Changed in 0.9.10

- Update dependencies

## [0.9.9] - 2025-05-09

### Changed in 0.9.9

- Update dependencies

## [0.9.8] - 2025-04-18

### Changed in 0.9.8

- Update dependencies

## [0.9.7] - 2025-04-15

### Changed in 0.9.7

- Update dependencies

## [0.9.6] - 2025-04-09

### Changed in 0.9.6

- Update dependencies

## [0.9.5] - 2025-02-27

### Changed in 0.9.5

- Update dependencies

## [0.9.4] - 2024-12-12

### Changed in 0.9.4

- Update dependencies

## [0.9.3] - 2024-10-30

### Changed in 0.9.3

- Migrate to using `SzAbstractFactory.Destroy()`
- Update dependencies

## [0.9.2] - 2024-09-12

### Changed in 0.9.2

- Update dependencies

## [0.9.1] - 2024-08-27

### Changed in 0.9.1

- Modify method calls to match Senzing API 4.0.0-24237

## [0.9.0] - 2024-08-26

### Changed in 0.9.0

- Change from `g2` to `sz`/`er`

## [0.8.1] - 2024-06-27

### Changed in 0.8.1

- Update dependencies

## [0.8.0] - 2024-05-09

### Changed in 0.8.0

- Migrated to improved FactoryCreator
- Update dependencies

## [0.7.0] - 2024-03-14

### Changed in 0.7.0

- Updated dependencies
- Deleted methods not used in V4

## [0.6.0] - 2024-01-29

### Changed in 0.6.0

- Renamed module to `github.com/senzing-garage/go-sdk-abstract-factory`
- Refactor to [template-go](https://github.com/senzing-garage-garage/template-go)
- Update dependencies
  - google.golang.org/grpc v1.61.0
  - github.com/senzing-garage/g2-sdk-go v0.9.0
  - github.com/senzing-garage/g2-sdk-go-base v0.5.0
  - github.com/senzing-garage/g2-sdk-go-grpc v0.6.0
  - github.com/senzing-garage/g2-sdk-proto/go v0.0.0-20240126210601-d02d3beb81d4

## [0.5.0] - 2024-01-03

### Changed in 0.5.0

- Refactor to [template-go](https://github.com/senzing-garage-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-common v0.4.0
  - github.com/senzing-garage/go-logging v1.4.0
  - github.com/senzing-garage/go-observing v0.3.0
  - github.com/senzing/g2-sdk-go v0.8.0
  - github.com/senzing/g2-sdk-go-base v0.4.0
  - github.com/senzing/g2-sdk-go-grpc v0.5.0
  - google.golang.org/grpc v1.60.1

## [0.4.3] - 2023-11-01

### Changed in 0.4.3

- Update dependencies
  - github.com/senzing/g2-sdk-go-base v0.3.3
  - github.com/senzing-garage/go-common v0.3.2-0.20231018174900-c1895fb44c30

## [0.4.2] - 2023-10-23

### Changed in 0.4.2

- Update to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.7.4
  - github.com/senzing/g2-sdk-go-base v0.3.2
  - github.com/senzing/g2-sdk-go-grpc v0.4.3
  - github.com/senzing/g2-sdk-proto/go v0.0.0-20231016131354-0d0fba649357
  - github.com/senzing-garage/go-common v0.3.1
  - github.com/senzing-garage/go-logging v1.3.3
  - github.com/senzing-garage/go-observing v0.2.8
  - google.golang.org/grpc v1.59.0

## [0.4.1] - 2023-10-13

### Changed in 0.4.1

- Changed from int to int64 where required by the SenzingAPI
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.7.3
  - google.golang.org/grpc v1.58.3

## [0.4.0] - 2023-10-03

### Changed in 0.4.0

- Supports SenzingAPI 3.8.0
- Deprecated functions have been removed
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.7.0
  - github.com/senzing/g2-sdk-go-base v0.3.0
  - github.com/senzing/g2-sdk-go-grpc v0.4.1
  - github.com/senzing/g2-sdk-proto/go v0.0.0-20230925212041-8259762ae97e
  - google.golang.org/grpc v1.58.2

## [0.3.3] - 2023-09-01

### Changed in 0.3.3

- Last version before SenzingAPI 3.8.0

## [0.3.2] - 2023-08-09

### Changed in 0.3.2

- Refactor to `template-go`
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.8
  - github.com/senzing/g2-sdk-go-base v0.2.4
  - github.com/senzing/g2-sdk-go-grpc v0.3.2
  - github.com/senzing-garage/go-common v0.2.11
  - github.com/senzing-garage/go-logging v1.3.2
  - github.com/senzing-garage/go-observing v0.2.7
  - google.golang.org/grpc v1.57.0

## [0.3.1] - 2023-06-16

### Changed in 0.3.1

- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.6
  - github.com/senzing/g2-sdk-go-base v0.2.1
  - github.com/senzing/g2-sdk-go-grpc v0.3.1
  - github.com/senzing/g2-sdk-proto/go v0.0.0-20230608182106-25c8cdc02e3c
  - github.com/senzing-garage/go-common v0.1.4
  - github.com/senzing-garage/go-logging v1.2.6
  - github.com/senzing-garage/go-observing v0.2.6
  - github.com/stretchr/testify v1.8.4
  - google.golang.org/grpc v1.56.0

## [0.3.0] - 2023-05-26

### Changed in 0.3.0

- Support Observers
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.4
  - github.com/senzing/g2-sdk-go-base v0.2.0
  - github.com/senzing/g2-sdk-proto/go v0.0.0-20230526140633-b44eb0f20e1b

## [0.2.4] - 2023-05-24

### Changed in 0.2.4

- Renamed `GrpcAddress` to `GrpcTarget`
- Renamed `GrpcOptions` to `GrpcDialOptions`

## [0.2.3] - 2023-05-11

### Changed in 0.2.3

- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.2
  - github.com/senzing/g2-sdk-go-base v0.1.10
  - github.com/senzing/g2-sdk-go-grpc v0.2.6
  - github.com/senzing-garage/go-common v0.1.3
  - github.com/senzing-garage/go-logging v1.2.3
  - github.com/senzing-garage/go-observing v0.2.2
  - google.golang.org/grpc v1.55.0

## [0.2.2] - 2023-04-21

### Changed in 0.2.2

- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.1
  - github.com/senzing/g2-sdk-go-base v0.1.8
  - github.com/senzing/g2-sdk-go-grpc v0.2.4

## [0.2.1] - 2023-03-02

### Added to 0.2.1

## [0.2.0] - 2023-02-16

### Added to 0.2.0

- Use updated interfaces
- Use Truth Set data in tests

## [0.1.0] - 2023-01-31

### Added to 0.1.0

- Initial implementation

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[markdownlint]: https://dlaa.me/markdownlint/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
