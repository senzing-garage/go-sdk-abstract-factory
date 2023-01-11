package factory

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
)

var (
	sdkAbstractFactoryLocalSingleton SdkAbstractFactory
	sdkAbstractFactoryGrpcSingleton  SdkAbstractFactory
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObjectLocal(ctx context.Context, test *testing.T) SdkAbstractFactory {
	if sdkAbstractFactoryLocalSingleton == nil {
		engineConfigurationJson, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if err != nil {
			test.Logf("Could not construct engineConfigurationJson. Error: %v", err)
		}

		sdkAbstractFactoryLocalSingleton = &SdkAbstractFactoryImpl{
			EngineConfigurationJson: engineConfigurationJson,
			ModuleName:              "Test module name",
			VerboseLogging:          0,
		}
		log.SetFlags(0)
	}
	return sdkAbstractFactoryLocalSingleton
}

func getTestObjectGrpc(ctx context.Context, test *testing.T) SdkAbstractFactory {
	if sdkAbstractFactoryGrpcSingleton == nil {
		engineConfigurationJson, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if err != nil {
			test.Logf("Could not construct engineConfigurationJson. Error: %v", err)
		}

		if sdkAbstractFactoryGrpcSingleton == nil {
			sdkAbstractFactoryGrpcSingleton = &SdkAbstractFactoryImpl{
				EngineConfigurationJson: engineConfigurationJson,
				GrpcAddress:             "localhost:8258",
				ModuleName:              "Test module name",
				VerboseLogging:          0,
			}
			log.SetFlags(0)
		}
	}
	return sdkAbstractFactoryGrpcSingleton
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if 1 == 0 {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

func testErrorNoFail(test *testing.T, ctx context.Context, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSdkAbstractFactoryImpl_GetG2config_local(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectLocal(ctx, test)
	g2config, err := testObject.GetG2config(ctx)
	testError(test, ctx, err)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, err)
	actual, err := g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, err)
	printActual(test, actual)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, err)
}

func TestSdkAbstractFactoryImpl_GetG2config_gRPC(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectGrpc(ctx, test)
	g2config, err := testObject.GetG2config(ctx)
	testError(test, ctx, err)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, err)
	actual, err := g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, err)
	printActual(test, actual)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, err)
}
