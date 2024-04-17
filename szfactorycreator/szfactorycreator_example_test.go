package szfactorycreator

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-helpers/engineconfigurationjson"
	"github.com/senzing-garage/sz-sdk-go/sz"
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
	verboseLogging := sz.SZ_NO_LOGGING
	configId := sz.SZ_INITIALIZE_WITH_DEFAULT_CONFIGURATION
	settings, err := engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingEnvVars()
	if err != nil {
		fmt.Println(err)
	}
	szAbstractFactory, err := CreateCoreAbstractFactory(instanceName, settings, verboseLogging, configId)
	if err != nil {
		fmt.Println(err)
	}
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer szEngine.Destroy(ctx)
	// Output:
}

func ExampleCreateGrpcAbstractFactory() {
	// For more information, visit https://github.com/senzing-garage/go-sdk-abstract-factory/blob/main/szfactorycreator/szfactorycreator_examples_test.go
	ctx := context.TODO()
	grpcAddress := "localhost:8261"
	grpcConnection, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Did not connect: %v\n", err)
	}
	szAbstractFactory, err := CreateGrpcAbstractFactory(grpcConnection)
	if err != nil {
		fmt.Println(err)
	}
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer szEngine.Destroy(ctx)
	// Output:
}
