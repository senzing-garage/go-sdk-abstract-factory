/*
Supported abstract factories:
  - [szabstractfactorycore.Szabstractfactory]
  - [szabstractfactorygrpc.Szabstractfactory]
  - [szabstractfactorymock.Szabstractfactory]
*/
package szfactorycreator

import (
	szabstractfactorycore "github.com/senzing-garage/sz-sdk-go-core/szabstractfactory"
	szabstractfactorygrpc "github.com/senzing-garage/sz-sdk-go-grpc/szabstractfactory"
	szabstractfactorymock "github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"

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

/*
Function CreateCoreAbstractFactory returns a Senzing Abstract Factory
that is used to create Senzing objects
which run natively.

Input
  - instanceName: A name for the auditing node, to help identify it within system logs.
  - settings: A JSON string containing configuration parameters.
  - verboseLogging: A flag to enable deeper logging of the Sz processing. 0 for no Senzing logging; 1 for logging.
  - configID: The configuration ID used for the initialization.  0 for current default configuration.

Output
  - A [szabstractfactorycore.Szabstractfactory] implementation conforming to the [senzing.SzAbstractFactory] interface.
*/
func CreateCoreAbstractFactory(instanceName string, settings string, verboseLogging int64, configID int64) (senzing.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactorycore.Szabstractfactory{
		ConfigID:       configID,
		InstanceName:   instanceName,
		Settings:       settings,
		VerboseLogging: verboseLogging,
	}
	return szAbstractFactory, nil
}

/*
Function CreateGrpcAbstractFactory returns a Senzing Abstract Factory
that is used to create Senzing objects
which communicate with a [Senzing gRPC server].

Input
  - grpcConnection: A connection to a Senzing gRPC server.

Output
  - A [szabstractfactorygrpc.Szabstractfactory] implementation conforming to the [senzing.SzAbstractFactory] interface.

[Senzing gRPC server]: https://github.com/senzing-garage/serve-grpc
*/
func CreateGrpcAbstractFactory(grpcConnection *grpc.ClientConn) (senzing.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactorygrpc.Szabstractfactory{
		GrpcConnection: grpcConnection,
	}
	return szAbstractFactory, nil
}

/*
Function CreateMockAbstractFactory returns a Senzing Abstract Factory
that is used to create Senzing [mock objects].

Output
  - A [szabstractfactorymock.Szabstractfactory] implementation conforming to the [senzing.SzAbstractFactory] interface.

[mock objects]: https://en.wikipedia.org/wiki/Mock_object
*/
func CreateMockAbstractFactory() (senzing.SzAbstractFactory, error) {
	szAbstractFactory := &szabstractfactorymock.Szabstractfactory{}
	return szAbstractFactory, nil
}
