package factory

import (
	"context"

	"github.com/senzing-garage/g2-sdk-go/g2api"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The SdkAbstractFactory interface shows what Senzing objects that can be retrieved from the abstract factory.
type SdkAbstractFactory interface {
	GetG2config(ctx context.Context) (g2api.G2config, error)
	GetG2configmgr(ctx context.Context) (g2api.G2configmgr, error)
	GetG2diagnostic(ctx context.Context) (g2api.G2diagnostic, error)
	GetG2engine(ctx context.Context) (g2api.G2engine, error)
	GetG2product(ctx context.Context) (g2api.G2product, error)
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the factory package found messages having the format "senzing-6041xxxx".
const (
	ComponentId       = 6041
	ImplementedByBase = "base"
	ImplementedByGrpc = "grpc"
	ImplementedByMock = "mock"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for the factory package.
var IdMessages = map[int]string{
	1:    "Enter AddDataSource(%v, %s).",
	2:    "Exit  AddDataSource(%v, %s) returned (%s, %v).",
	4001: "Cannot G2Config.Init()",
	4002: "Cannot G2Configmgr.Init()",
	4003: "Cannot G2Diagnostic.Init()",
	4004: "Cannot G2Engine.Init()",
	4005: "Cannot G2Product.Init()",
	4010: "Did not make a gRPC connection",
}

// Status strings for specific factory messages.
var IdStatuses = map[int]string{}
