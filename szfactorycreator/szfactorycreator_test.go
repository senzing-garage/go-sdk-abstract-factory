package szfactorycreator

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-helpers/settings"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSzfactorycreator_CreateCoreAbstractFactory(test *testing.T) {
	ctx := context.TODO()
	instanceName := "Test name"
	settings, err := settings.BuildSimpleSettingsUsingEnvVars()
	require.NoError(test, err)
	verboseLogging := senzing.SzNoLogging
	configID := senzing.SzInitializeWithDefaultConfiguration
	szAbstractFactory, err := CreateCoreAbstractFactory(instanceName, settings, verboseLogging, configID)
	require.NoError(test, err)
	_, err = szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
}

func TestSzfactorycreator_CreateGrpcAbstractFactory(test *testing.T) {
	ctx := context.TODO()
	grpcAddress := "localhost:8261"
	grpcConnection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(test, err)
	szAbstractFactory, err := CreateGrpcAbstractFactory(grpcConnection)
	require.NoError(test, err)
	_, err = szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
}

func TestSzfactorycreator_CreateMockAbstractFactory(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory, err := CreateMockAbstractFactory()
	require.NoError(test, err)
	_, err = szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
