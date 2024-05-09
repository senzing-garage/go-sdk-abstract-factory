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

	"github.com/senzing-garage/go-helpers/engineconfigurationjson"
	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-sdk-abstract-factory/szfactorycreator"
	"github.com/senzing-garage/sz-sdk-go/sz"
)

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

var logger logging.LoggingInterface = nil

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getLogger(ctx context.Context) (logging.LoggingInterface, error) {
	_ = ctx
	loggerOptions := []interface{}{
		&logging.OptionCallerSkip{Value: 3},
	}
	return logging.NewSenzingToolsLogger(9999, IdMessages, loggerOptions...)
}

func demonstrateConfigFunctions(ctx context.Context, szConfig sz.SzConfig, szConfigmgr sz.SzConfigManager) error {
	now := time.Now()

	// Using SzConfig: Create a default configuration in memory.

	configHandle, err := szConfig.CreateConfig(ctx)
	if err != nil {
		return logger.NewError(5100, err)
	}

	// Using SzConfig: Add data source to in-memory configuration.

	for dataSourceCode := range truthset.TruthsetDataSources {
		_, err := szConfig.AddDataSource(ctx, configHandle, dataSourceCode)
		if err != nil {
			return logger.NewError(5101, err)
		}
	}

	// Using SzConfig: Persist configuration to a string.

	configStr, err := szConfig.ExportConfig(ctx, configHandle)
	if err != nil {
		return logger.NewError(5102, err)
	}

	// Using SzConfigmgr: Persist configuration string to database.

	configComments := fmt.Sprintf("Created by go-sdk-abstract_factory_test at %s", now.UTC())
	configID, err := szConfigmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		return logger.NewError(5103, err)
	}

	// Using SzConfigmgr: Set new configuration as the default.

	err = szConfigmgr.SetDefaultConfigId(ctx, configID)
	if err != nil {
		return logger.NewError(5104, err)
	}

	return err
}

func demonstrateAddRecord(ctx context.Context, szEngine sz.SzEngine) (string, error) {
	dataSourceCode := "TEST"
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000000))
	if err != nil {
		panic(err)
	}
	recordId := randomNumber.String()
	recordDefinition := fmt.Sprintf(
		"%s%s%s",
		`{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "`,
		recordId,
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`)
	flags := sz.SZ_WITH_INFO

	// Using SzEngine: Add record and return "withInfo".

	return szEngine.AddRecord(ctx, dataSourceCode, recordId, recordDefinition, flags)
}

func demonstrateAdditionalFunctions(ctx context.Context, szDiagnostic sz.SzDiagnostic, szEngine sz.SzEngine, szProduct sz.SzProduct) error {

	// Using SzDiagnostic: Purge repository.

	err := szDiagnostic.PurgeRepository(ctx)
	if err != nil {
		failOnError(5301, err)
	}

	// Using SzEngine: Add records with information returned.

	withInfo, err := demonstrateAddRecord(ctx, szEngine)
	if err != nil {
		failOnError(5302, err)
	}
	logger.Log(2003, withInfo)

	// Using szProduct: Show license metadata.

	license, err := szProduct.GetLicense(ctx)
	if err != nil {
		failOnError(5303, err)
	}
	logger.Log(2004, license)

	// Using SzDiagnostic: Purge repository again.

	err = szDiagnostic.PurgeRepository(ctx)
	if err != nil {
		failOnError(5304, err)
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
	var szAbstractFactory sz.SzAbstractFactory
	var testcaseList []int
	ctx := context.TODO()

	// Configure the "log" standard library.

	log.SetFlags(0)
	logger, err = getLogger(ctx)
	if err != nil {
		failOnError(5000, err)
	}

	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
	logger.Log(2001, "Just a test of logging")

	// Create Senzing's Engine Configuration JSON.

	instanceName := "Test name"
	settings, err := engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingEnvVars()
	if err != nil {
		failOnError(5001, err)
	}
	verboseLogging := sz.SZ_NO_LOGGING
	configId := sz.SZ_INITIALIZE_WITH_DEFAULT_CONFIGURATION

	// Determine if specific testcase is requested.

	testcaseNumber := os.Getenv("SENZING_TOOLS_TESTCASE_NUMBER")
	if len(testcaseNumber) > 0 {
		testcaseInt, err := strconv.Atoi(testcaseNumber)
		if err != nil {
			failOnError(5002, err)
		}
		testcaseList = append(testcaseList, testcaseInt)
	} else {
		testcaseList = append(testcaseList, 1)
	}

	// Iterate through different instantiations of SdkAbstractFactory.

	for _, runNumber := range testcaseList {
		fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

		// Choose different implementations.

		switch runNumber {
		case 1:
			logger.Log(2001, "Local SDK")
			szAbstractFactory, err = szfactorycreator.CreateCoreAbstractFactory(instanceName, settings, verboseLogging, configId)
			if err != nil {
				failOnError(9999, err)
			}
		default:
			failOnError(5003, fmt.Errorf("unknown testcase number"))
		}

		// Get Senzing objects for installing a Senzing Engine configuration.

		szConfig, err := szAbstractFactory.CreateSzConfig(ctx)
		if err != nil {
			failOnError(5004, err)
		}
		defer szConfig.Destroy(ctx)

		szConfigManager, err := szAbstractFactory.CreateSzConfigManager(ctx)
		if err != nil {
			failOnError(5007, err)
		}
		defer szConfigManager.Destroy(ctx)

		// Persist the Senzing configuration to the Senzing repository.

		err = demonstrateConfigFunctions(ctx, szConfig, szConfigManager)
		if err != nil {
			failOnError(5011, err)
		}

		// Now that a Senzing configuration is installed, get the remainder of the Senzing objects.

		szDiagnostic, err := szAbstractFactory.CreateSzDiagnostic(ctx)
		if err != nil {
			failOnError(5012, err)
		}
		defer szDiagnostic.Destroy(ctx)

		szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
		if err != nil {
			failOnError(5014, err)
		}
		defer szEngine.Destroy(ctx)

		szProduct, err := szAbstractFactory.CreateSzProduct(ctx)
		if err != nil {
			failOnError(5016, err)
		}
		defer szProduct.Destroy(ctx)

		// Demonstrate tests.

		err = demonstrateAdditionalFunctions(ctx, szDiagnostic, szEngine, szProduct)
		if err != nil {
			failOnError(5021, err)
		}

	}
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
}
