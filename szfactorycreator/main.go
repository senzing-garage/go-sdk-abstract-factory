package szfactorycreator

import (
	szabstractfactorycore "github.com/senzing-garage/sz-sdk-go-core/szabstractfactory"
	szabstractfactorygrpc "github.com/senzing-garage/sz-sdk-go-grpc/szabstractfactory"
	"google.golang.org/grpc"

	"github.com/senzing-garage/sz-sdk-go/sz"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the szfactorycreator package found messages having the format "senzing-6041xxxx".
const ComponentId = 6041

// ----------------------------------------------------------------------------
// Factory builders
// ----------------------------------------------------------------------------

func CreateCoreAbstractFactory(instanceName string, settings string, verboseLogging int64, configId int64) (sz.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactorycore.Szabstractfactory{
		ConfigId:       configId,
		InstanceName:   instanceName,
		Settings:       settings,
		VerboseLogging: verboseLogging,
	}
	return szAbstractFactory, nil
}

func CreateGrpcAbstractFactory(grpcConnection *grpc.ClientConn) (sz.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactorygrpc.Szabstractfactory{
		GrpcConnection: grpcConnection,
	}
	return szAbstractFactory, nil
}

func CreateMockAbstractFactory() (sz.SzAbstractFactory, error) {
	return nil, nil
}
