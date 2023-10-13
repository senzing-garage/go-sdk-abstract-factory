package factory

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/go-common/g2engineconfigurationjson"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
)

var (
	iniParams                        string
	moduleName                       string = "Test module name"
	sdkAbstractFactoryGrpcSingleton  SdkAbstractFactory
	sdkAbstractFactoryLocalSingleton SdkAbstractFactory
	verboseLogging                   int64 = 0
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObjectLocal(ctx context.Context, test *testing.T) SdkAbstractFactory {
	if sdkAbstractFactoryLocalSingleton == nil {
		sdkAbstractFactoryLocalSingleton = &SdkAbstractFactoryImpl{}
	}
	return sdkAbstractFactoryLocalSingleton
}

func getTestObjectGrpc(ctx context.Context, test *testing.T) SdkAbstractFactory {
	if sdkAbstractFactoryGrpcSingleton == nil {
		sdkAbstractFactoryGrpcSingleton = &SdkAbstractFactoryImpl{
			GrpcTarget: "localhost:8261",
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

func errorId(err error) string {
	var result string = ""
	if err != nil {
		errorMessage := err.Error()[strings.Index(err.Error(), "{"):]
		var dictionary map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(errorMessage), &dictionary)
		if unmarshalErr != nil {
			fmt.Print("Unmarshal Error:", unmarshalErr.Error())
		}
		result = dictionary["id"].(string)
	}
	return result
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
	iniParams, err = g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingEnvVars()
	return err
}

func teardown() error {
	var err error = nil
	return err
}

func TestBuildSimpleSystemConfigurationJsonUsingEnvVars(test *testing.T) {
	actual, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingEnvVars()
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test helper functions
// ----------------------------------------------------------------------------

func helperSdkAbstractFactoryImpl_GetG2config(test *testing.T, ctx context.Context, testObject SdkAbstractFactory) {
	g2config, err := testObject.GetG2config(ctx)
	testError(test, ctx, err)
	err = g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	if errorId(err) != "senzing-60114002" {
		testError(test, ctx, err)
	}
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, err)
	actual, err := g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, err)
	printActual(test, actual)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, err)
}

func helperSdkAbstractFactoryImpl_GetG2configmgr(test *testing.T, ctx context.Context, testObject SdkAbstractFactory) {
	g2config, err := testObject.GetG2config(ctx)
	testError(test, ctx, err)
	err = g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	if errorId(err) != "senzing-60114002" {
		testError(test, ctx, err)
	}
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, err)
	configStr, err := g2config.Save(ctx, configHandle)
	testError(test, ctx, err)

	g2configmgr, err := testObject.GetG2configmgr(ctx)
	testError(test, ctx, err)
	err = g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	if errorId(err) != "senzing-60124002" {
		testError(test, ctx, err)
	}
	now := time.Now()
	configComments := fmt.Sprintf("Created by g2diagnostic_test at %s", now.UTC())
	configID, err := g2configmgr.AddConfig(ctx, configStr, configComments)
	testError(test, ctx, err)
	err = g2configmgr.SetDefaultConfigID(ctx, configID)
	testError(test, ctx, err)
}

func helperSdkAbstractFactoryImpl_GetG2diagnostic(test *testing.T, ctx context.Context, testObject SdkAbstractFactory) {
	g2diagnostic, err := testObject.GetG2diagnostic(ctx)
	testError(test, ctx, err)
	err = g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	if errorId(err) != "senzing-60134002" {
		testError(test, ctx, err)
	}
	actual, err := g2diagnostic.GetTotalSystemMemory(ctx)
	testError(test, ctx, err)
	printActual(test, actual)
}

func helperSdkAbstractFactoryImpl_GetG2engine(test *testing.T, ctx context.Context, testObject SdkAbstractFactory) {
	g2engine, err := testObject.GetG2engine(ctx)
	testError(test, ctx, err)
	err = g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
	if errorId(err) != "senzing-60144002" {
		testError(test, ctx, err)
	}
	actual, err := g2engine.Stats(ctx)
	testError(test, ctx, err)
	printActual(test, actual)
}

func helperSdkAbstractFactoryImpl_GetG2product(test *testing.T, ctx context.Context, testObject SdkAbstractFactory) {
	g2product, err := testObject.GetG2product(ctx)
	testError(test, ctx, err)
	err = g2product.Init(ctx, moduleName, iniParams, verboseLogging)
	if errorId(err) != "senzing-60164002" {
		testError(test, ctx, err)
	}
	actual, err := g2product.License(ctx)
	testError(test, ctx, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSdkAbstractFactoryImpl_GetG2config_local(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectLocal(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2config(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2config_gRPC(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectGrpc(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2config(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2configmgr_local(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectLocal(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2configmgr(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2configmgr_gRPC(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectGrpc(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2configmgr(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2diagnostic_local(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectLocal(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2diagnostic(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2diagnostic_gRPC(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectGrpc(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2diagnostic(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2engine_local(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectLocal(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2engine(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2engine_gRPC(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectGrpc(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2engine(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2product_local(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectLocal(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2product(test, ctx, testObject)
}

func TestSdkAbstractFactoryImpl_GetG2product_gRPC(test *testing.T) {
	ctx := context.TODO()
	testObject := getTestObjectGrpc(ctx, test)
	helperSdkAbstractFactoryImpl_GetG2product(test, ctx, testObject)
}
