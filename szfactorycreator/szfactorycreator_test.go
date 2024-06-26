package szfactorycreator

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-helpers/settings"
	"github.com/senzing-garage/sz-sdk-go/senzing"
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
	settings, err := settings.BuildSimpleSettingsUsingEnvVars()
	testError(test, err)
	verboseLogging := senzing.SzNoLogging
	configId := senzing.SzInitializeWithDefaultConfiguration
	szAbstractFactory, err := CreateCoreAbstractFactory(instanceName, settings, verboseLogging, configId)
	testError(test, err)
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	testError(test, err)
	defer szEngine.Destroy(ctx)
}

func TestSzfactorycreator_CreateGrpcAbstractFactory(test *testing.T) {
	ctx := context.TODO()
	grpcAddress := "localhost:8261"
	grpcConnection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	testError(test, err)
	szAbstractFactory, err := CreateGrpcAbstractFactory(grpcConnection)
	testError(test, err)
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	testError(test, err)
	defer szEngine.Destroy(ctx)
}
