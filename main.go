package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/senzing-garage/go-helpers/settings"
	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-sdk-abstract-factory/szfactorycreator"
	"github.com/senzing-garage/sz-sdk-go/senzing"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var IDMessages = map[int]string{
	1:    "%s",
	2:    "WithInfo: %s",
	2001: "Testing %s.",
	2002: "Physical cores: %d.",
	2003: "withInfo",
	2004: "License",
	2005: "Implementation: %s",
	2999: "Cannot retrieve last error message.",
}

var logger logging.Logging

var errMain = errors.New("main")

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	var err error
	var testcaseList []int
	ctx := context.TODO()

	// Configure the "log" standard library.

	log.SetFlags(0)
	logger, err = getLogger(ctx)
	failOnError(5000, err)

	outputf("\n-------------------------------------------------------------------------------\n\n")
	logger.Log(2001, "Just a test of logging")

	// Determine if specific testcase is requested.

	testcaseNumber := os.Getenv("SENZING_TOOLS_TESTCASE_NUMBER")
	if len(testcaseNumber) > 0 {
		testcaseInt, err := strconv.Atoi(testcaseNumber)
		failOnError(5002, err)
		testcaseList = append(testcaseList, testcaseInt)
	} else {
		testcaseList = append(testcaseList, 1)
	}

	// Iterate through different instantiations of SdkAbstractFactory.

	testCases(ctx, testcaseList)

	outputf("\n-------------------------------------------------------------------------------\n\n")
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func demonstrateAddRecord(ctx context.Context, szEngine senzing.SzEngine) (string, error) {
	dataSourceCode := "TEST"
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000000))
	if err != nil {
		panic(err)
	}
	recordID := randomNumber.String()
	recordDefinition := fmt.Sprintf(
		"%s%s%s",
		`{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "`,
		recordID,
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
	)
	flags := senzing.SzWithInfo

	// Using SzEngine: Add record and return "withInfo".

	return szEngine.AddRecord(ctx, dataSourceCode, recordID, recordDefinition, flags)
}

func demonstrateConfigFunctions(ctx context.Context, szAbstractFactory senzing.SzAbstractFactory) error {
	now := time.Now()

	// Create Senzing objects.

	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		return logger.NewError(5100, err)
	}

	szConfig, err := szConfigManager.CreateConfigFromTemplate(ctx)
	if err != nil {
		return logger.NewError(5101, err)
	}

	// Using SzConfig: Add data sources to Senzing configuration.

	for dataSourceCode := range truthset.TruthsetDataSources {
		_, err := szConfig.AddDataSource(ctx, dataSourceCode)
		if err != nil {
			return logger.NewError(5102, err)
		}
	}

	// Using SzConfig: Persist configuration to a string.

	configDefinition, err := szConfig.Export(ctx)
	if err != nil {
		return logger.NewError(5103, err)
	}

	// Using SzConfigManager: Persist configuration string to database.

	configComment := fmt.Sprintf("Created by go-sdk-abstract_factory_test at %s", now.UTC())
	_, err = szConfigManager.SetDefaultConfig(ctx, configDefinition, configComment)
	if err != nil {
		return logger.NewError(5104, err)
	}

	return err
}

func demonstrateAdditionalFunctions(
	ctx context.Context,
	szAbstractFactory senzing.SzAbstractFactory,
) error {
	var err error

	// Create Senzing objects.

	szDiagnostic, err := szAbstractFactory.CreateDiagnostic(ctx)
	if err != nil {
		return logger.NewError(5300, err)
	}

	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		return logger.NewError(5301, err)
	}

	szProduct, err := szAbstractFactory.CreateProduct(ctx)
	if err != nil {
		return logger.NewError(5302, err)
	}

	// Using SzDiagnostic: Purge repository.

	err = szDiagnostic.PurgeRepository(ctx)
	failOnError(5303, err)

	// Using SzEngine: Add records with information returned.

	withInfo, err := demonstrateAddRecord(ctx, szEngine)
	failOnError(5304, err)
	logger.Log(2003, withInfo)

	// Using szProduct: Show license metadata.

	license, err := szProduct.GetLicense(ctx)
	failOnError(5305, err)
	logger.Log(2004, license)

	// Using SzDiagnostic: Purge repository again.

	err = szDiagnostic.PurgeRepository(ctx)
	failOnError(5306, err)

	return err
}

func failOnError(msgID int, err error) {
	if err != nil {
		logger.Log(msgID, err)
		panic(err)
	}
}

func outputf(format string, message ...any) {
	fmt.Printf(format, message...) //nolint
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func testCases(ctx context.Context, testcaseList []int) {
	var err error
	var szAbstractFactory senzing.SzAbstractFactory

	// Create Senzing's Engine Configuration JSON.

	instanceName := "Test name"
	settings, err := settings.BuildSimpleSettingsUsingEnvVars()
	failOnError(5001, err)
	verboseLogging := senzing.SzNoLogging
	configID := senzing.SzInitializeWithDefaultConfiguration

	for _, runNumber := range testcaseList {
		outputf("\n-------------------------------------------------------------------------------\n\n")

		// Choose different implementations.

		switch runNumber {
		case 1:
			logger.Log(2001, "Local SDK")
			szAbstractFactory, err = szfactorycreator.CreateCoreAbstractFactory(
				instanceName,
				settings,
				verboseLogging,
				configID,
			)
			failOnError(9999, err)
		default:
			failOnError(5003, wraperror.Errorf(errMain, "unknown testcase number"))
		}
		defer func() { panicOnError(szAbstractFactory.Destroy(ctx)) }()

		// Get Senzing objects for installing a Senzing Engine configuration.

		_, err = szAbstractFactory.CreateConfigManager(ctx)
		failOnError(5005, err)

		_, err = szAbstractFactory.CreateDiagnostic(ctx)
		failOnError(5006, err)

		_, err = szAbstractFactory.CreateEngine(ctx)
		failOnError(5007, err)

		_, err = szAbstractFactory.CreateProduct(ctx)
		failOnError(5008, err)

		// Persist the Senzing configuration to the Senzing repository.

		err = demonstrateConfigFunctions(ctx, szAbstractFactory)
		failOnError(5009, err)

		// Demonstrate tests.

		err = demonstrateAdditionalFunctions(ctx, szAbstractFactory)
		failOnError(5010, err)

	}
}

func getLogger(ctx context.Context) (logging.Logging, error) {
	_ = ctx
	loggerOptions := []interface{}{
		logging.OptionCallerSkip{Value: 3},
		logging.OptionMessageFields{Value: []string{"id", "text", "reason", "errors"}},
	}
	return logging.NewSenzingLogger(9999, IDMessages, loggerOptions...)
}
