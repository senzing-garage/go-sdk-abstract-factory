package factory

import (
	"context"
	"fmt"

	"github.com/senzing/g2-sdk-go-grpc/g2configclient"
	"github.com/senzing/g2-sdk-go-grpc/g2configmgrclient"
	"github.com/senzing/g2-sdk-go-grpc/g2diagnosticclient"
	"github.com/senzing/g2-sdk-go-grpc/g2engineclient"
	"github.com/senzing/g2-sdk-go-grpc/g2productclient"
	"github.com/senzing/g2-sdk-go/g2config"
	"github.com/senzing/g2-sdk-go/g2configmgr"
	"github.com/senzing/g2-sdk-go/g2diagnostic"
	"github.com/senzing/g2-sdk-go/g2engine"
	"github.com/senzing/g2-sdk-go/g2product"
	pbg2config "github.com/senzing/g2-sdk-proto/go/g2config"
	pbg2configmgr "github.com/senzing/g2-sdk-proto/go/g2configmgr"
	pbg2diagnostic "github.com/senzing/g2-sdk-proto/go/g2diagnostic"
	pbg2engine "github.com/senzing/g2-sdk-proto/go/g2engine"
	pbg2product "github.com/senzing/g2-sdk-proto/go/g2product"
	"github.com/senzing/go-logging/messagelogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// G2configImpl is the default implementation of the G2config interface.
type SdkAbstractFactoryImpl struct {
	EngineConfigurationJson string
	GrpcAddress             string
	ModuleName              string
	VerboseLogging          int
	logger                  messagelogger.MessageLoggerInterface
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (factory *SdkAbstractFactoryImpl) getGrpcConnection() *grpc.ClientConn {
	result, err := grpc.Dial(factory.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		factory.getLogger().Log(4010, err)

	}
	return result
}

// Get the Logger singleton.
func (factory *SdkAbstractFactoryImpl) getLogger() messagelogger.MessageLoggerInterface {
	if factory.logger == nil {
		factory.logger, _ = messagelogger.NewSenzingApiLogger(ProductId, IdMessages, IdStatuses, messagelogger.LevelInfo)
	}
	return factory.logger
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetG2Config method returns a G2config object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcAddress is spectified, an implementation that communicates over gRPC will be returned.
If GrpcAddress is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2config object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2Config(ctx context.Context) (g2config.G2config, error) {
	var result g2config.G2config
	var err error = nil

	// Determine which instantiation of the G2config interface to create.

	if len(factory.GrpcAddress) > 0 {
		grpcConnection := factory.getGrpcConnection()
		result = &g2configclient.G2configClient{
			GrpcClient: pbg2config.NewG2ConfigClient(grpcConnection),
		}
	} else {
		result = &g2config.G2configImpl{}
	}

	// Initialize the object.

	err = result.Init(ctx, factory.ModuleName, factory.EngineConfigurationJson, factory.VerboseLogging)
	if err != nil {
		factory.getLogger().Log(4001, err)
	}
	return result, err
}

func (factory *SdkAbstractFactoryImpl) GetG2ConfigMgr(ctx context.Context) (g2configmgr.G2configmgr, error) {
	var result g2configmgr.G2configmgr
	var err error = nil

	// Determine which instantiation of the G2configmgr interface to create.

	if len(factory.GrpcAddress) > 0 {
		grpcConnection := factory.getGrpcConnection()
		result = &g2configmgrclient.G2configmgrClient{
			GrpcClient: pbg2configmgr.NewG2ConfigMgrClient(grpcConnection),
		}
	} else {
		result = &g2configmgr.G2configmgrImpl{}
	}

	// Initialize the object.

	err = result.Init(ctx, factory.ModuleName, factory.EngineConfigurationJson, factory.VerboseLogging)
	if err != nil {
		factory.getLogger().Log(4002, err)
	}
	return result, err
}

func (factory *SdkAbstractFactoryImpl) GetG2Diagnostic(ctx context.Context) (g2diagnostic.G2diagnostic, error) {
	var result g2diagnostic.G2diagnostic
	var err error = nil

	// Determine which instantiation of the G2diagnostic interface to create.

	if len(factory.GrpcAddress) > 0 {
		grpcConnection := factory.getGrpcConnection()
		result = &g2diagnosticclient.G2diagnosticClient{
			GrpcClient: pbg2diagnostic.NewG2DiagnosticClient(grpcConnection),
		}
	} else {
		result = &g2diagnostic.G2diagnosticImpl{}
	}

	// Initialize the object.

	err = result.Init(ctx, factory.ModuleName, factory.EngineConfigurationJson, factory.VerboseLogging)
	if err != nil {
		factory.getLogger().Log(4003, err)
	}
	return result, err
}

func (factory *SdkAbstractFactoryImpl) GetG2Engine(ctx context.Context) (g2engine.G2engine, error) {
	var result g2engine.G2engine
	var err error = nil

	// Determine which instantiation of the G2engine interface to create.

	if len(factory.GrpcAddress) > 0 {
		grpcConnection := factory.getGrpcConnection()
		result = &g2engineclient.G2engineClient{
			GrpcClient: pbg2engine.NewG2EngineClient(grpcConnection),
		}
	} else {
		result = &g2engine.G2engineImpl{}
	}

	// Initialize the object.

	err = result.Init(ctx, factory.ModuleName, factory.EngineConfigurationJson, factory.VerboseLogging)
	if err != nil {
		factory.getLogger().Log(4004, err)
	}
	return result, err

}

func (factory *SdkAbstractFactoryImpl) GetG2Product(ctx context.Context) (g2product.G2product, error) {
	var result g2product.G2product
	var err error = nil

	// Determine which instantiation of the G2product interface to create.

	if len(factory.GrpcAddress) > 0 {
		grpcConnection := factory.getGrpcConnection()
		result = &g2productclient.G2productClient{
			GrpcClient: pbg2product.NewG2ProductClient(grpcConnection),
		}
	} else {
		result = &g2product.G2productImpl{}
	}

	// Initialize the object.

	err = result.Init(ctx, factory.ModuleName, factory.EngineConfigurationJson, factory.VerboseLogging)
	if err != nil {
		factory.getLogger().Log(4005, err)
	}
	return result, err
}
