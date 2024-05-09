package szfactorycreator

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-helpers/engineconfigurationjson"
	"github.com/senzing-garage/sz-sdk-go/sz"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func testError(test *testing.T, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSzfactorycreator_CreateCoreAbstractFactory(test *testing.T) {
	ctx := context.TODO()
	instanceName := "Test name"
	settings, err := engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingEnvVars()
	testError(test, err)
	verboseLogging := sz.SZ_NO_LOGGING
	configId := sz.SZ_INITIALIZE_WITH_DEFAULT_CONFIGURATION
	szAbstractFactory, err := CreateCoreAbstractFactory(instanceName, settings, verboseLogging, configId)
	testError(test, err)
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	testError(test, err)
	defer szEngine.Destroy(ctx)
}

func TestSzfactorycreator_CreateGrpcAbstractFactory(test *testing.T) {
	ctx := context.TODO()
	grpcAddress := "localhost:8261"
	grpcConnection, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	testError(test, err)
	szAbstractFactory, err := CreateGrpcAbstractFactory(grpcConnection)
	testError(test, err)
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	testError(test, err)
	defer szEngine.Destroy(ctx)
}
