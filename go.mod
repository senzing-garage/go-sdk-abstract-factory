module github.com/senzing/go-sdk-abstract-factory

go 1.19

require (
	github.com/aquilax/truncate v1.0.0
	github.com/senzing/g2-sdk-go v0.3.1
	github.com/senzing/g2-sdk-go-base v0.0.0-20230209201723-0d99b7147739
	github.com/senzing/g2-sdk-go-grpc v0.1.0
	github.com/senzing/g2-sdk-proto/go v0.0.0-20230126140313-273e96bc7dbd
	github.com/senzing/go-common v0.1.1
	github.com/senzing/go-logging v1.1.3
	github.com/stretchr/testify v1.8.1
	google.golang.org/grpc v1.53.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/senzing/go-observing v0.1.1 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/senzing/g2-sdk-go v0.3.1 => /home/senzing/senzing.git/g2-sdk-go
	github.com/senzing/g2-sdk-go-base v0.0.0-20230209201723-0d99b7147739 => /home/senzing/senzing.git/g2-sdk-go-base
	github.com/senzing/g2-sdk-go-grpc v0.1.0 => /home/senzing/senzing.git/g2-sdk-go-grpc
)
