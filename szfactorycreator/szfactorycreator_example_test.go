package szfactorycreator

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-helpers/settings"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleCreateCoreAbstractFactory() {
	// For more information, visit https://github.com/senzing-garage/go-sdk-abstract-factory/blob/main/szfactorycreator/szfactorycreator_examples_test.go
	ctx := context.TODO()
	instanceName := "Test name"
	verboseLogging := senzing.SzNoLogging
	configID := senzing.SzInitializeWithDefaultConfiguration
	settings, err := settings.BuildSimpleSettingsUsingEnvVars()
	if err != nil {
		fmt.Println(err)
	}
	szAbstractFactory, err := CreateCoreAbstractFactory(instanceName, settings, verboseLogging, configID)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_ = szEngine // Use szEngine.
	// Output:
}

func ExampleCreateGrpcAbstractFactory() {
	// For more information, visit https://github.com/senzing-garage/go-sdk-abstract-factory/blob/main/szfactorycreator/szfactorycreator_examples_test.go
	ctx := context.TODO()
	grpcAddress := "localhost:8261"
	grpcConnection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Did not connect: %v\n", err)
	}
	szAbstractFactory, err := CreateGrpcAbstractFactory(grpcConnection)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_ = szEngine // Use szEngine.
	// Output:
}
