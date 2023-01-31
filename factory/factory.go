package factory

import (
	"context"
	"sync"

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

// SdkAbstractFactoryImpl is the default implementation of the SdkAbstractFactory interface.
type SdkAbstractFactoryImpl struct {
	g2configmgrSingleton  g2configmgr.G2configmgr
	g2configmgrSyncOnce   sync.Once
	g2configSingleton     g2config.G2config
	g2configSyncOnce      sync.Once
	g2diagnosticSingleton g2diagnostic.G2diagnostic
	g2diagnosticSyncOnce  sync.Once
	g2engineSingleton     g2engine.G2engine
	g2engineSyncOnce      sync.Once
	g2productSingleton    g2product.G2product
	g2productSyncOnce     sync.Once
	GrpcAddress           string
	GrpcOptions           []grpc.DialOption
	logger                messagelogger.MessageLoggerInterface
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the gRPC connection.
func (factory *SdkAbstractFactoryImpl) getGrpcConnection(ctx context.Context) *grpc.ClientConn {
	if factory.GrpcOptions == nil {
		factory.GrpcOptions = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}
	result, err := grpc.DialContext(ctx, factory.GrpcAddress, factory.GrpcOptions...)
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
The GetG2config method returns a G2config object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcAddress is spectified, an implementation that communicates over gRPC will be returned.
If GrpcAddress is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2config object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2config(ctx context.Context) (g2config.G2config, error) {
	var err error = nil
	factory.g2configSyncOnce.Do(func() {
		if len(factory.GrpcAddress) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2configSingleton = &g2configclient.G2configClient{
				GrpcClient: pbg2config.NewG2ConfigClient(grpcConnection),
			}
		} else {
			factory.g2configSingleton = &g2config.G2configImpl{}
		}
	})
	return factory.g2configSingleton, err
}

/*
The GetG2configmgr method returns a G2configmgr object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcAddress is spectified, an implementation that communicates over gRPC will be returned.
If GrpcAddress is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2configmgr object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2configmgr(ctx context.Context) (g2configmgr.G2configmgr, error) {
	var err error = nil
	factory.g2configmgrSyncOnce.Do(func() {
		if len(factory.GrpcAddress) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2configmgrSingleton = &g2configmgrclient.G2configmgrClient{
				GrpcClient: pbg2configmgr.NewG2ConfigMgrClient(grpcConnection),
			}
		} else {
			factory.g2configmgrSingleton = &g2configmgr.G2configmgrImpl{}
		}
	})
	return factory.g2configmgrSingleton, err
}

/*
The GetG2diagnostic method returns a G2diagnostic object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcAddress is spectified, an implementation that communicates over gRPC will be returned.
If GrpcAddress is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2diagnostic object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2diagnostic(ctx context.Context) (g2diagnostic.G2diagnostic, error) {
	var err error = nil
	factory.g2diagnosticSyncOnce.Do(func() {
		if len(factory.GrpcAddress) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2diagnosticSingleton = &g2diagnosticclient.G2diagnosticClient{
				GrpcClient: pbg2diagnostic.NewG2DiagnosticClient(grpcConnection),
			}
		} else {
			factory.g2diagnosticSingleton = &g2diagnostic.G2diagnosticImpl{}
		}
	})
	return factory.g2diagnosticSingleton, err
}

/*
The GetG2engine method returns a G2engine object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcAddress is spectified, an implementation that communicates over gRPC will be returned.
If GrpcAddress is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2engine object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2engine(ctx context.Context) (g2engine.G2engine, error) {
	var err error = nil
	factory.g2engineSyncOnce.Do(func() {
		if len(factory.GrpcAddress) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2engineSingleton = &g2engineclient.G2engineClient{
				GrpcClient: pbg2engine.NewG2EngineClient(grpcConnection),
			}
		} else {
			factory.g2engineSingleton = &g2engine.G2engineImpl{}
		}
	})
	return factory.g2engineSingleton, err
}

/*
The GetG2product method returns a G2product object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcAddress is spectified, an implementation that communicates over gRPC will be returned.
If GrpcAddress is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2product object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2product(ctx context.Context) (g2product.G2product, error) {
	var err error = nil
	factory.g2productSyncOnce.Do(func() {
		if len(factory.GrpcAddress) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2productSingleton = &g2productclient.G2productClient{
				GrpcClient: pbg2product.NewG2ProductClient(grpcConnection),
			}
		} else {
			factory.g2productSingleton = &g2product.G2productImpl{}
		}
	})
	return factory.g2productSingleton, err
}
