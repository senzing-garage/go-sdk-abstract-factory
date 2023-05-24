package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/senzing/g2-sdk-go/g2api"
	"github.com/senzing/go-common/g2engineconfigurationjson"
	"github.com/senzing/go-common/truthset"
	"github.com/senzing/go-logging/logging"

	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-sdk-abstract-factory/factory"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IdMessages = map[int]string{
	1:    "%s",
	2:    "WithInfo: %s",
	2001: "Testing %s.",
	2002: "Physical cores: %d.",
	2003: "withInfo",
	2004: "License",
	2005: "Implementation: %s",
	2999: "Cannot retrieve last error message.",
}

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"
var logger logging.LoggingInterface = nil

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getLogger(ctx context.Context) (logging.LoggingInterface, error) {
	loggerOptions := []interface{}{
		&logging.OptionCallerSkip{Value: 3},
	}
	return logging.NewSenzingToolsLogger(9999, IdMessages, loggerOptions...)
}

func demonstrateConfigFunctions(ctx context.Context, g2Config g2api.G2config, g2Configmgr g2api.G2configmgr) error {
	now := time.Now()

	// Print the SDK implementation.

	sdkId := g2Config.GetSdkId(ctx)
	logger.Log(2005, sdkId)

	// Using G2Config: Create a default configuration in memory.

	configHandle, err := g2Config.Create(ctx)
	if err != nil {
		return logger.NewError(5100, err)
	}

	// Using G2Config: Add data source to in-memory configuration.

	for _, testDataSource := range truthset.TruthsetDataSources {
		_, err := g2Config.AddDataSource(ctx, configHandle, testDataSource.Json)
		if err != nil {
			return logger.NewError(5101, err)
		}
	}

	// Using G2Config: Persist configuration to a string.

	configStr, err := g2Config.Save(ctx, configHandle)
	if err != nil {
		return logger.NewError(5102, err)
	}

	// Using G2Configmgr: Persist configuration string to database.

	configComments := fmt.Sprintf("Created by g2diagnostic_test at %s", now.UTC())
	configID, err := g2Configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		return logger.NewError(5103, err)
	}

	// Using G2Configmgr: Set new configuration as the default.

	err = g2Configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		return logger.NewError(5104, err)
	}

	return err
}

func demonstrateAddRecord(ctx context.Context, g2Engine g2api.G2engine) (string, error) {
	dataSourceCode := "TEST"
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000000))
	if err != nil {
		panic(err)
	}
	recordID := randomNumber.String()
	jsonData := fmt.Sprintf(
		"%s%s%s",
		`{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "`,
		recordID,
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`)
	loadID := dataSourceCode
	var flags int64 = 0

	// Using G2Engine: Add record and return "withInfo".

	return g2Engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
}

func demonstrateAdditionalFunctions(ctx context.Context, g2Diagnostic g2api.G2diagnostic, g2Engine g2api.G2engine, g2Product g2api.G2product) error {
	// Using G2Diagnostic: Check physical cores.

	actual, err := g2Diagnostic.GetPhysicalCores(ctx)
	if err != nil {
		failOnError(5300, err)
	}
	logger.Log(2002, actual)

	// Using G2Engine: Purge repository.

	err = g2Engine.PurgeRepository(ctx)
	if err != nil {
		failOnError(5301, err)
	}

	// Using G2Engine: Add records with information returned.

	withInfo, err := demonstrateAddRecord(ctx, g2Engine)
	if err != nil {
		failOnError(5302, err)
	}
	logger.Log(2003, withInfo)

	// Using G2Product: Show license metadata.

	license, err := g2Product.License(ctx)
	if err != nil {
		failOnError(5303, err)
	}
	logger.Log(2004, license)

	// Using G2Engine: Purge repository again.

	err = g2Engine.PurgeRepository(ctx)
	if err != nil {
		failOnError(5304, err)
	}

	return err
}

func destroyObjects(ctx context.Context, g2Config g2api.G2config, g2Configmgr g2api.G2configmgr, g2Diagnostic g2api.G2diagnostic, g2Engine g2api.G2engine, g2Product g2api.G2product) error {
	var err error = nil

	// Destroy is only needed for "base" implementation.

	if g2Config.GetSdkId(ctx) == "base" {
		err = g2Config.Destroy(ctx)
		if err != nil {
			failOnError(5401, err)
		}

		err = g2Configmgr.Destroy(ctx)
		if err != nil {
			failOnError(5402, err)
		}

		err = g2Diagnostic.Destroy(ctx)
		if err != nil {
			failOnError(5403, err)
		}

		err = g2Engine.Destroy(ctx)
		if err != nil {
			failOnError(5404, err)
		}

		err = g2Product.Destroy(ctx)
		if err != nil {
			failOnError(5405, err)
		}
	}

	return err
}

func failOnError(msgId int, err error) {
	logger.Log(msgId, err)
	panic(err.Error())
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	var err error = nil
	var senzingFactory factory.SdkAbstractFactory
	var testcaseList []int
	ctx := context.TODO()

	// Configure the "log" standard library.

	log.SetFlags(0)
	logger, err = getLogger(ctx)
	if err != nil {
		failOnError(5000, err)
	}

	// Test logger.

	programmMetadataMap := map[string]interface{}{
		"ProgramName":    programName,
		"BuildVersion":   buildVersion,
		"BuildIteration": buildIteration,
	}

	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
	logger.Log(2001, "Just a test of logging", programmMetadataMap)

	// Create 2 observers.

	observer1 := &observer.ObserverNull{
		Id: "Observer 1",
	}
	observer2 := &observer.ObserverNull{
		Id: "Observer 2",
	}

	// Create Senzing's Engine Configuration JSON.

	moduleName := "Test module name"
	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		failOnError(5001, err)
	}
	verboseLogging := 0

	// Determine if specific testcase is requested.

	testcaseNumber := os.Getenv("SENZING_TOOLS_TESTCASE_NUMBER")
	if len(testcaseNumber) > 0 {
		testcaseInt, err := strconv.Atoi(testcaseNumber)
		if err != nil {
			failOnError(5002, err)
		}
		testcaseList = append(testcaseList, testcaseInt)

	} else {
		testcaseList = append(testcaseList, 1, 2)
	}

	// Iterate through different instantiations of SdkAbstractFactory.

	for _, runNumber := range testcaseList {
		fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

		// Choose different implementations.

		switch runNumber {
		case 1:
			logger.Log(2001, "Local SDK")
			senzingFactory = &factory.SdkAbstractFactoryImpl{}
		case 2:
			logger.Log(2001, "gRPC SDK")
			senzingFactory = &factory.SdkAbstractFactoryImpl{
				GrpcTarget: "localhost:8258",
			}
		default:
			failOnError(5003, fmt.Errorf("unknown testcase number"))
		}

		// Get Senzing objects for installing a Senzing Engine configuration.

		g2Config, err := senzingFactory.GetG2config(ctx)
		if err != nil {
			failOnError(5004, err)
		}
		err = g2Config.RegisterObserver(ctx, observer1)
		if err != nil {
			failOnError(5005, err)
		}
		err = g2Config.RegisterObserver(ctx, observer2)
		if err != nil {
			failOnError(5006, err)
		}

		g2Configmgr, err := senzingFactory.GetG2configmgr(ctx)
		if err != nil {
			failOnError(5007, err)
		}
		err = g2Configmgr.RegisterObserver(ctx, observer1)
		if err != nil {
			failOnError(5008, err)
		}

		// Initialize "base" implementations.

		if g2Config.GetSdkId(ctx) == "base" {
			err = g2Config.Init(ctx, moduleName, iniParams, verboseLogging)
			if err != nil {
				failOnError(5009, err)
			}
			err = g2Configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
			if err != nil {
				failOnError(5010, err)
			}
		}

		// Persist the Senzing configuration to the Senzing repository.

		err = demonstrateConfigFunctions(ctx, g2Config, g2Configmgr)
		if err != nil {
			failOnError(5011, err)
		}

		// Now that a Senzing configuration is installed, get the remainder of the Senzing objects.

		g2Diagnostic, err := senzingFactory.GetG2diagnostic(ctx)
		if err != nil {
			failOnError(5012, err)
		}
		err = g2Diagnostic.RegisterObserver(ctx, observer1)
		if err != nil {
			failOnError(5013, err)
		}

		g2Engine, err := senzingFactory.GetG2engine(ctx)
		if err != nil {
			failOnError(5014, err)
		}
		err = g2Engine.RegisterObserver(ctx, observer1)
		if err != nil {
			failOnError(5015, err)
		}

		g2Product, err := senzingFactory.GetG2product(ctx)
		if err != nil {
			failOnError(5016, err)
		}
		err = g2Product.RegisterObserver(ctx, observer1)
		if err != nil {
			failOnError(5017, err)
		}

		// Initialize "base" implementations.

		if g2Diagnostic.GetSdkId(ctx) == "base" {
			err = g2Diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
			if err != nil {
				failOnError(5018, err)
			}
			err = g2Engine.Init(ctx, moduleName, iniParams, verboseLogging)
			if err != nil {
				failOnError(5019, err)
			}
			err = g2Product.Init(ctx, moduleName, iniParams, verboseLogging)
			if err != nil {
				failOnError(5020, err)
			}
		}

		// Demonstrate tests.

		err = demonstrateAdditionalFunctions(ctx, g2Diagnostic, g2Engine, g2Product)
		if err != nil {
			failOnError(5021, err)
		}

		// Destroy Senzing objects.

		err = destroyObjects(ctx, g2Config, g2Configmgr, g2Diagnostic, g2Engine, g2Product)
		if err != nil {
			failOnError(5022, err)
		}
	}
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
}
