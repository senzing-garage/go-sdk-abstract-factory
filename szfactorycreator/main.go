package szfactorycreator

import (
	"github.com/senzing-garage/sz-sdk-go-core/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go/sz"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the szfactorycreator package found messages having the format "senzing-6001xxxx".
// TODO:  Assign ComponentId
const ComponentId = 9999

// ----------------------------------------------------------------------------
// Factory builders
// ----------------------------------------------------------------------------

func CreateCoreAbstractFactory(instanceName string, settings string, verboseLogging int64, configId int64) (sz.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactory.Szabstractfactory{
		ConfigId:       configId,
		InstanceName:   instanceName,
		Settings:       settings,
		VerboseLogging: verboseLogging,
	}
	return szAbstractFactory, nil
}

func CreateGrpcAbstractFactory() (sz.SzAbstractFactory, error) {
	return nil, nil
}

func CreateMockAbstractFactory() (sz.SzAbstractFactory, error) {
	return nil, nil
}
