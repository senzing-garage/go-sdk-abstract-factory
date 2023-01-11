package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/senzing/g2-sdk-go/g2config"
	"github.com/senzing/g2-sdk-go/g2configmgr"
	"github.com/senzing/g2-sdk-go/g2diagnostic"
	"github.com/senzing/g2-sdk-go/g2engine"
	"github.com/senzing/g2-sdk-go/g2product"
	"github.com/senzing/g2-sdk-go/testhelpers"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelevel"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
	"github.com/senzing/go-sdk-abstract-factory/factory"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var Messages = map[int]string{
	1:    "%s",
	2:    "WithInfo: %s",
	2999: "Cannot retrieve last error message.",
	4001: "Testing %s.",
}

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"
var logger messagelogger.MessageLoggerInterface = nil

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

func getLogger(ctx context.Context) (messagelogger.MessageLoggerInterface, error) {
	messageFormat := &messageformat.MessageFormatJson{}
	messageIdTemplate := &messageid.MessageIdTemplated{
		MessageIdTemplate: MessageIdTemplate,
	}
	messageLevel := &messagelevel.MessageLevelByIdRange{
		IdLevelRanges: messagelevel.IdLevelRanges,
	}
	messageStatus := &messagestatus.MessageStatusByIdRange{
		IdStatusRanges: messagestatus.IdLevelRangesAsString,
	}
	messageText := &messagetext.MessageTextTemplated{
		IdMessages: Messages,
	}
	return messagelogger.New(messageFormat, messageIdTemplate, messageLevel, messageStatus, messageText, messagelogger.LevelInfo)
}

func getConfigString(ctx context.Context, g2Config g2config.G2config) (string, error) {
	configHandle, err := g2Config.Create(ctx)
	if err != nil {
		return "", logger.Error(5100, err)
	}

	for _, testDataSource := range testhelpers.TestDataSources {
		_, err := g2Config.AddDataSource(ctx, configHandle, testDataSource.Data)
		if err != nil {
			logger.Error(5102, err)
		}
	}

	return g2Config.Save(ctx, configHandle)
}

func persistConfig(ctx context.Context, g2Configmgr g2configmgr.G2configmgr, configStr string, configComments string) error {

	configID, err := g2Configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		return logger.Error(5200, err)
	}

	err = g2Configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		return logger.Error(5201, err)
	}

	return err
}

func demonstrateAddRecord(ctx context.Context, g2Engine g2engine.G2engine) (string, error) {
	dataSourceCode := "TEST"
	recordID := strconv.Itoa(rand.Intn(1000000000))
	jsonData := fmt.Sprintf(
		"%s%s%s",
		`{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "`,
		recordID,
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`)
	loadID := dataSourceCode
	var flags int64 = 0

	return g2Engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
}

func demonstrateAdditionalFunctions(ctx context.Context, g2Diagnostic g2diagnostic.G2diagnostic, g2Engine g2engine.G2engine, g2Product g2product.G2product) error {
	var err error = nil

	// Using G2Diagnostic: Check physical cores.

	actual, err := g2Diagnostic.GetPhysicalCores(ctx)
	if err != nil {
		logger.Log(5300, err)
	}
	logger.Log(2300, "Physical cores", actual)

	// Using G2Engine: Purge repository.

	err = g2Engine.PurgeRepository(ctx)
	if err != nil {
		logger.Log(5301, err)
	}

	// Using G2Engine: Add records with information returned.

	withInfo, err := demonstrateAddRecord(ctx, g2Engine)
	if err != nil {
		logger.Log(5302, err)
	}
	logger.Log(2301, "WithInfo", withInfo)

	// Using G2Product: Show license metadata.

	license, err := g2Product.License(ctx)
	if err != nil {
		logger.Log(5303, err)
	}
	logger.Log(2302, "License", license)

	return err
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	var err error = nil
	var senzingFactory factory.SdkAbstractFactory
	ctx := context.TODO()
	now := time.Now()

	// Randomize random number generator.

	rand.Seed(time.Now().UnixNano())

	// Configure the "log" standard library.

	log.SetFlags(0)
	logger, err = getLogger(ctx)
	if err != nil {
		logger.Log(5000, err)
	}

	// Test logger.

	programmMetadataMap := map[string]interface{}{
		"ProgramName":    programName,
		"BuildVersion":   buildVersion,
		"BuildIteration": buildIteration,
	}

	logger.Log(2001, "Just a test of logging", programmMetadataMap)

	engineConfigurationJson, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		logger.Log(5001, err)
	}

	for _, runNumber := range []int{1, 2} {
		fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

		// Choose different implementations.

		switch runNumber {
		case 1:
			logger.Log(4001, "Local SDK")
			senzingFactory = &factory.SdkAbstractFactoryImpl{
				EngineConfigurationJson: engineConfigurationJson,
				ModuleName:              "Test module name",
				VerboseLogging:          0,
			}
		case 2:
			logger.Log(4001, "gRPC SDK")
			senzingFactory = &factory.SdkAbstractFactoryImpl{
				EngineConfigurationJson: engineConfigurationJson,
				GrpcAddress:             "localhost:8258",
				ModuleName:              "Test module name",
				VerboseLogging:          0,
			}
		default:
			logger.Log(5002)
		}

		// Get Senzing objects for installing a Senzing Engine configuration.

		g2Config, err := senzingFactory.GetG2config(ctx)
		if err != nil {
			logger.Log(5003, err)
		}

		g2Configmgr, err := senzingFactory.GetG2configmgr(ctx)
		if err != nil {
			logger.Log(5004, err)
		}

		// Using G2Config: Create a default Senzing Engine Configuration in memory.

		configStr, err := getConfigString(ctx, g2Config)
		if err != nil {
			logger.Log(5005, err)
		}

		// Using G2Configmgr: Persist the Senzing configuration to the Senzing repository.

		configComments := fmt.Sprintf("Created by g2diagnostic_test at %s", now.UTC())
		err = persistConfig(ctx, g2Configmgr, configStr, configComments)
		if err != nil {
			logger.Log(5006, err)
		}

		// Now that a Senzing configuration is installed, get the remainder of the Senzing objects.

		g2Diagnostic, err := senzingFactory.GetG2diagnostic(ctx)
		if err != nil {
			logger.Log(5007, err)
		}

		g2Engine, err := senzingFactory.GetG2engine(ctx)
		if err != nil {
			logger.Log(5008, err)
		}

		g2Product, err := senzingFactory.GetG2product(ctx)
		if err != nil {
			logger.Log(5009, err)
		}

		// Demonstrate tests.

		err = demonstrateAdditionalFunctions(ctx, g2Diagnostic, g2Engine, g2Product)
		if err != nil {
			logger.Log(5010, err)
		}

		// Destroy Senzing objects.

		err = g2Config.Destroy(ctx)
		if err != nil {
			logger.Log(5011, err)
		}

		err = g2Configmgr.Destroy(ctx)
		if err != nil {
			logger.Log(5012, err)
		}

		err = g2Diagnostic.Destroy(ctx)
		if err != nil {
			logger.Log(5013, err)
		}

		err = g2Engine.Destroy(ctx)
		if err != nil {
			logger.Log(5014, err)
		}
	}
}
