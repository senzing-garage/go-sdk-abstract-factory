package szfactorycreator

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-helpers/engineconfigurationjson"
	"github.com/senzing-garage/sz-sdk-go/sz"
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
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer szEngine.Destroy(ctx)
	// Output:
}
