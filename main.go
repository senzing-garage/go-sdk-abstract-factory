package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
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
	2001: "Testing %s.",
	2002: "Physical cores: %d.",
	2003: "withInfo",
	2004: "License",
	2999: "Cannot retrieve last error message.",
}

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

func demonstrateConfigFunctions(ctx context.Context, g2Config g2config.G2config, g2Configmgr g2configmgr.G2configmgr) error {
	now := time.Now()

	// Using G2Config: Create a default configuration in memory

	configHandle, err := g2Config.Create(ctx)
	if err != nil {
		return logger.Error(5100, err)
	}

	// Using G2Config: Add data source to in-memory configuration.

	for _, testDataSource := range testhelpers.TestDataSources {
		_, err := g2Config.AddDataSource(ctx, configHandle, testDataSource.Data)
		if err != nil {
			return logger.Error(5101, err)
		}
	}

	// Using G2Config: Persist configuration to a string.

	configStr, err := g2Config.Save(ctx, configHandle)
	if err != nil {
		return logger.Error(5102, err)
	}

	// Using G2Configmgr: Persist configuration string to database.

	configComments := fmt.Sprintf("Created by g2diagnostic_test at %s", now.UTC())
	configID, err := g2Configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		return logger.Error(5103, err)
	}

	// Using G2Configmgr: Set new configuration as the default.

	err = g2Configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		return logger.Error(5104, err)
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

	// Using G2Engine: Add record and return "withInfo".

	return g2Engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
}

func demonstrateAdditionalFunctions(ctx context.Context, g2Diagnostic g2diagnostic.G2diagnostic, g2Engine g2engine.G2engine, g2Product g2product.G2product) error {
	var err error = nil

	// Using G2Diagnostic: Check physical cores.

	actual, err := g2Diagnostic.GetPhysicalCores(ctx)
	if err != nil {
		logger.Log(5300, err)
	}
	logger.Log(2002, actual)

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
	logger.Log(2003, withInfo)

	// Using G2Product: Show license metadata.

	license, err := g2Product.License(ctx)
	if err != nil {
		logger.Log(5303, err)
	}
	logger.Log(2004, license)

	return err
}

func destroyObjects(ctx context.Context, g2Config g2config.G2config, g2Configmgr g2configmgr.G2configmgr, g2Diagnostic g2diagnostic.G2diagnostic, g2Engine g2engine.G2engine, g2Product g2product.G2product) error {
	var err error = nil
	errorList := []string{}

	err = g2Config.Destroy(ctx)
	if (err != nil) && (errorId(err) != "senzing-60114001") {
		logger.Log(5401, err)
		errorList = append(errorList, "g2Config")
	}

	err = g2Configmgr.Destroy(ctx)
	if (err != nil) && (errorId(err) != "senzing-60124001") {
		logger.Log(5402, err)
		errorList = append(errorList, "g2Configmgr")
	}

	err = g2Diagnostic.Destroy(ctx)
	if (err != nil) && (errorId(err) != "senzing-60134001") {
		logger.Log(5403, err)
		errorList = append(errorList, "g2Diagnostic")
	}

	err = g2Engine.Destroy(ctx)
	if (err != nil) && (errorId(err) != "senzing-60144001") {
		logger.Log(5404, err)
		errorList = append(errorList, "g2Engine")

	}

	err = g2Product.Destroy(ctx)
	if (err != nil) && (errorId(err) != "senzing-60164001") {
		logger.Log(5405, err)
		errorList = append(errorList, "g2Product")
	}

	if len(errorList) == 0 {
		err = nil
	} else {
		errorListString := strings.Join(errorList, ", ")
		err = fmt.Errorf("errors in %s", errorListString)
	}

	return err
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
// Main
// ----------------------------------------------------------------------------

func main() {
	var err error = nil
	var senzingFactory factory.SdkAbstractFactory
	ctx := context.TODO()

	// Randomize random number generator.

	rand.Seed(time.Now().UnixNano())

	// Configure the "log" standard library.

	log.SetFlags(0)
	logger, err = getLogger(ctx)
	if err != nil {
		logger.Log(5000, err)
	}

	// Create Senzing's Engine Configuration JSON.

	moduleName := "Test module name"
	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		logger.Log(5001, err)
	}
	verboseLogging := 0

	// Iterate through different instantiations of SdkAbstractFactory.

	for _, runNumber := range []int{1, 2} {
		fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

		// Choose different implementations.

		switch runNumber {
		case 1:
			logger.Log(2001, "Local SDK")
			senzingFactory = &factory.SdkAbstractFactoryImpl{}
		case 2:
			logger.Log(2001, "gRPC SDK")
			senzingFactory = &factory.SdkAbstractFactoryImpl{
				GrpcAddress: "localhost:8258",
			}
		default:
			logger.Log(5002)
		}

		// Get Senzing objects for installing a Senzing Engine configuration.

		g2Config, err := senzingFactory.GetG2config(ctx)
		if err != nil {
			logger.Log(5003, err)
		}
		err = g2Config.Init(ctx, moduleName, iniParams, verboseLogging)
		if (err != nil) && (errorId(err) != "senzing-60114002") {
			logger.Log(5004, err)
		}

		g2Configmgr, err := senzingFactory.GetG2configmgr(ctx)
		if err != nil {
			logger.Log(5005, err)
		}
		err = g2Configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
		if (err != nil) && (errorId(err) != "senzing-60124002") {
			logger.Log(5006, err)
		}

		// Persist the Senzing configuration to the Senzing repository.

		err = demonstrateConfigFunctions(ctx, g2Config, g2Configmgr)
		if err != nil {
			logger.Log(5007, err)
		}

		// Now that a Senzing configuration is installed, get the remainder of the Senzing objects.

		g2Diagnostic, err := senzingFactory.GetG2diagnostic(ctx)
		if err != nil {
			logger.Log(5008, err)
		}
		err = g2Diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
		if (err != nil) && (errorId(err) != "senzing-60134002") {
			logger.Log(5009, err)
		}

		g2Engine, err := senzingFactory.GetG2engine(ctx)
		if err != nil {
			logger.Log(5010, err)
		}
		err = g2Engine.Init(ctx, moduleName, iniParams, verboseLogging)
		if (err != nil) && (errorId(err) != "senzing-60144002") {
			logger.Log(5011, err)
		}

		g2Product, err := senzingFactory.GetG2product(ctx)
		if err != nil {
			logger.Log(5012, err)
		}
		err = g2Product.Init(ctx, moduleName, iniParams, verboseLogging)
		if (err != nil) && (errorId(err) != "senzing-60164002") {
			logger.Log(5013, err)
		}

		// Demonstrate tests.

		err = demonstrateAdditionalFunctions(ctx, g2Diagnostic, g2Engine, g2Product)
		if err != nil {
			logger.Log(5014, err)
		}

		// Destroy Senzing objects.

		err = destroyObjects(ctx, g2Config, g2Configmgr, g2Diagnostic, g2Engine, g2Product)
		if err != nil {
			logger.Log(5015, err)
		}

	}
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
}