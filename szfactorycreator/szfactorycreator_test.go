package szfactorycreator_test

import (
	"testing"

	"github.com/senzing-garage/go-helpers/settings"
	"github.com/senzing-garage/go-sdk-abstract-factory/szfactorycreator"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSzfactorycreator_CreateCoreAbstractFactory(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	instanceName := "Test name"
	settings, err := settings.BuildSimpleSettingsUsingEnvVars()
	require.NoError(test, err)

	verboseLogging := senzing.SzNoLogging
	configID := senzing.SzInitializeWithDefaultConfiguration
	szAbstractFactory, err := szfactorycreator.CreateCoreAbstractFactory(
		instanceName,
		settings,
		verboseLogging,
		configID,
	)
	require.NoError(test, err)
	_, err = szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
}

func TestSzfactorycreator_CreateGrpcAbstractFactory(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	grpcAddress := "localhost:8261"
	grpcConnection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(test, err)
	szAbstractFactory, err := szfactorycreator.CreateGrpcAbstractFactory(grpcConnection)
	require.NoError(test, err)
	_, err = szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
}

func TestSzfactorycreator_CreateMockAbstractFactory(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szAbstractFactory, err := szfactorycreator.CreateMockAbstractFactory()
	require.NoError(test, err)
	_, err = szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
