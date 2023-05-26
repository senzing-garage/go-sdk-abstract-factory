package factory

import (
	"context"
	"sync"

	g2configbase "github.com/senzing/g2-sdk-go-base/g2config"
	g2configmgrbase "github.com/senzing/g2-sdk-go-base/g2configmgr"
	g2diagnosticbase "github.com/senzing/g2-sdk-go-base/g2diagnostic"
	g2enginebase "github.com/senzing/g2-sdk-go-base/g2engine"
	g2productbase "github.com/senzing/g2-sdk-go-base/g2product"
	g2configgrpc "github.com/senzing/g2-sdk-go-grpc/g2config"
	g2configmgrgrpc "github.com/senzing/g2-sdk-go-grpc/g2configmgr"
	g2diagnosticgrpc "github.com/senzing/g2-sdk-go-grpc/g2diagnostic"
	g2enginegrpc "github.com/senzing/g2-sdk-go-grpc/g2engine"
	g2productgrpc "github.com/senzing/g2-sdk-go-grpc/g2product"
	"github.com/senzing/g2-sdk-go/g2api"
	g2configpb "github.com/senzing/g2-sdk-proto/go/g2config"
	g2configmgrpb "github.com/senzing/g2-sdk-proto/go/g2configmgr"
	g2diagnosticpb "github.com/senzing/g2-sdk-proto/go/g2diagnostic"
	g2enginepb "github.com/senzing/g2-sdk-proto/go/g2engine"
	g2productpb "github.com/senzing/g2-sdk-proto/go/g2product"
	"github.com/senzing/go-logging/logging"
	"github.com/senzing/go-observing/observer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// SdkAbstractFactoryImpl is the default implementation of the SdkAbstractFactory interface.
type SdkAbstractFactoryImpl struct {
	g2configmgrSingleton  g2api.G2configmgr
	g2configmgrSyncOnce   sync.Once
	g2configSingleton     g2api.G2config
	g2configSyncOnce      sync.Once
	g2diagnosticSingleton g2api.G2diagnostic
	g2diagnosticSyncOnce  sync.Once
	g2engineSingleton     g2api.G2engine
	g2engineSyncOnce      sync.Once
	g2productSingleton    g2api.G2product
	g2productSyncOnce     sync.Once
	GrpcDialOptions       []grpc.DialOption
	GrpcTarget            string
	logger                logging.LoggingInterface
	ObserverOrigin        string
	Observers             []observer.Observer
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Get the gRPC connection.
func (factory *SdkAbstractFactoryImpl) getGrpcConnection(ctx context.Context) *grpc.ClientConn {
	if factory.GrpcDialOptions == nil {
		factory.GrpcDialOptions = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}
	result, err := grpc.DialContext(ctx, factory.GrpcTarget, factory.GrpcDialOptions...)
	if err != nil {
		factory.getLogger().Log(4010, err)
	}
	return result
}

// Get the Logger singleton.
func (factory *SdkAbstractFactoryImpl) getLogger() logging.LoggingInterface {
	var err error = nil
	if factory.logger == nil {
		loggerOptions := []interface{}{
			&logging.OptionCallerSkip{Value: 3},
		}
		factory.logger, err = logging.NewSenzingToolsLogger(ComponentId, IdMessages, loggerOptions...)
		if err != nil {
			panic(err)
		}
	}
	return factory.logger
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetG2config method returns a G2config object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcTarget is spectified, an implementation that communicates over gRPC will be returned.
If GrpcTarget is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2config object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2config(ctx context.Context) (g2api.G2config, error) {
	var err error = nil
	factory.g2configSyncOnce.Do(func() {
		if len(factory.GrpcTarget) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2configSingleton = &g2configgrpc.G2config{
				GrpcClient: g2configpb.NewG2ConfigClient(grpcConnection),
			}
		} else {
			factory.g2configSingleton = &g2configbase.G2config{}
		}
		factory.g2configSingleton.SetObserverOrigin(ctx, factory.ObserverOrigin)
		for _, observer := range factory.Observers {
			factory.g2configSingleton.RegisterObserver(ctx, observer)
		}
	})
	return factory.g2configSingleton, err
}

/*
The GetG2configmgr method returns a G2configmgr object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcTarget is spectified, an implementation that communicates over gRPC will be returned.
If GrpcTarget is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2configmgr object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2configmgr(ctx context.Context) (g2api.G2configmgr, error) {
	var err error = nil
	factory.g2configmgrSyncOnce.Do(func() {
		if len(factory.GrpcTarget) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2configmgrSingleton = &g2configmgrgrpc.G2configmgr{
				GrpcClient: g2configmgrpb.NewG2ConfigMgrClient(grpcConnection),
			}
		} else {
			factory.g2configmgrSingleton = &g2configmgrbase.G2configmgr{}
		}
		factory.g2configmgrSingleton.SetObserverOrigin(ctx, factory.ObserverOrigin)
		for _, observer := range factory.Observers {
			factory.g2configmgrSingleton.RegisterObserver(ctx, observer)
		}
	})
	return factory.g2configmgrSingleton, err
}

/*
The GetG2diagnostic method returns a G2diagnostic object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcTarget is spectified, an implementation that communicates over gRPC will be returned.
If GrpcTarget is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2diagnostic object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2diagnostic(ctx context.Context) (g2api.G2diagnostic, error) {
	var err error = nil
	factory.g2diagnosticSyncOnce.Do(func() {
		if len(factory.GrpcTarget) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2diagnosticSingleton = &g2diagnosticgrpc.G2diagnostic{
				GrpcClient: g2diagnosticpb.NewG2DiagnosticClient(grpcConnection),
			}
		} else {
			factory.g2diagnosticSingleton = &g2diagnosticbase.G2diagnostic{}
		}
		factory.g2diagnosticSingleton.SetObserverOrigin(ctx, factory.ObserverOrigin)
		for _, observer := range factory.Observers {
			factory.g2diagnosticSingleton.RegisterObserver(ctx, observer)
		}
	})
	return factory.g2diagnosticSingleton, err
}

/*
The GetG2engine method returns a G2engine object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcTarget is spectified, an implementation that communicates over gRPC will be returned.
If GrpcTarget is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2engine object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2engine(ctx context.Context) (g2api.G2engine, error) {
	var err error = nil
	factory.g2engineSyncOnce.Do(func() {
		if len(factory.GrpcTarget) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2engineSingleton = &g2enginegrpc.G2engine{
				GrpcClient: g2enginepb.NewG2EngineClient(grpcConnection),
			}
		} else {
			factory.g2engineSingleton = &g2enginebase.G2engine{}
		}
		factory.g2engineSingleton.SetObserverOrigin(ctx, factory.ObserverOrigin)
		for _, observer := range factory.Observers {
			factory.g2engineSingleton.RegisterObserver(ctx, observer)
		}
	})
	return factory.g2engineSingleton, err
}

/*
The GetG2product method returns a G2product object based on the
information passed in the SdkAbstractFactoryImpl structure.
If GrpcTarget is spectified, an implementation that communicates over gRPC will be returned.
If GrpcTarget is empty, an implementation that uses a local Senzing Go SDK will be returned.

Input
  - ctx: A context to control lifecycle.

Output
  - An initialized G2product object.
    See the example output.
*/
func (factory *SdkAbstractFactoryImpl) GetG2product(ctx context.Context) (g2api.G2product, error) {
	var err error = nil
	factory.g2productSyncOnce.Do(func() {
		if len(factory.GrpcTarget) > 0 {
			grpcConnection := factory.getGrpcConnection(ctx)
			factory.g2productSingleton = &g2productgrpc.G2product{
				GrpcClient: g2productpb.NewG2ProductClient(grpcConnection),
			}
		} else {
			factory.g2productSingleton = &g2productbase.G2product{}
		}
		factory.g2productSingleton.SetObserverOrigin(ctx, factory.ObserverOrigin)
		for _, observer := range factory.Observers {
			factory.g2productSingleton.RegisterObserver(ctx, observer)
		}
	})
	return factory.g2productSingleton, err
}
