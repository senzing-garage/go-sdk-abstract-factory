package szfactorycreator

import (
	szabstractfactorycore "github.com/senzing-garage/sz-sdk-go-core/szabstractfactory"
	szabstractfactorygrpc "github.com/senzing-garage/sz-sdk-go-grpc/szabstractfactory"
	"google.golang.org/grpc"

	"github.com/senzing-garage/sz-sdk-go/senzing"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the szfactorycreator package found messages having the format "senzing-6041xxxx".
const ComponentID = 6041

// ----------------------------------------------------------------------------
// Factory builders
// ----------------------------------------------------------------------------

func CreateCoreAbstractFactory(instanceName string, settings string, verboseLogging int64, configID int64) (senzing.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactorycore.Szabstractfactory{
		ConfigID:       configID,
		InstanceName:   instanceName,
		Settings:       settings,
		VerboseLogging: verboseLogging,
	}
	return szAbstractFactory, nil
}

func CreateGrpcAbstractFactory(grpcConnection *grpc.ClientConn) (senzing.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactorygrpc.Szabstractfactory{
		GrpcConnection: grpcConnection,
	}
	return szAbstractFactory, nil
}

func CreateMockAbstractFactory() (senzing.SzAbstractFactory, error) {
	return nil, nil
}
