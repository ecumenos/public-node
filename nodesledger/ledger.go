package nodesledger

import (
	"github.com/ecumenos/public-node/nodesledger/database"
	"go.uber.org/fx"
)

//go:generate mockery --name=NodesLedger
var _ NodesLedger = (*ledger)(nil)

// NodesLedger is interface for nodes ledger logic.
type NodesLedger interface{}

// Module is fx module.
var Module = fx.Options(
	database.Module,
	fx.Provide(func() NodesLedger {
		return &ledger{}
	}),
)

type ledger struct{}
