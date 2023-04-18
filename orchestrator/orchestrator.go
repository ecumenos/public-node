package orchestrator

import (
	"github.com/ecumenos/public-node/clientsledger"
	"github.com/ecumenos/public-node/nodesledger"
	"github.com/ecumenos/public-node/secretmanager"
	"go.uber.org/fx"
)

//go:generate mockery --name=Orchestrator
var _ Orchestrator = (*orchestrator)(nil)

// Orchestrator is interface for orchestrator logic.
type Orchestrator interface{}

// Params is input parameters for creation of orchestrator.
type Params struct {
	fx.In

	ClientsLedger clientsledger.ClientsLedger
	NodesLedger   nodesledger.NodesLedger
	SecretManager secretmanager.SecretManager
}

// Module is fx module.
var Module = fx.Options(
	fx.Provide(func(params Params) Orchestrator {
		return &orchestrator{
			clientsLedger: params.ClientsLedger,
			nodesLedger:   params.NodesLedger,
			secretManager: params.ClientsLedger,
		}
	}),
)

type orchestrator struct {
	clientsLedger clientsledger.ClientsLedger
	nodesLedger   nodesledger.NodesLedger
	secretManager secretmanager.SecretManager
}
