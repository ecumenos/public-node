package clientsledger

import (
	"github.com/ecumenos/public-node/clientsledger/database"
	"go.uber.org/fx"
)

//go:generate mockery --name=ClientsLedger
var _ ClientsLedger = (*ledger)(nil)

// ClientsLedger is interface for clients ledger logic.
type ClientsLedger interface{}

// Module is fx module.
var Module = fx.Options(
	database.Module,
	fx.Provide(func() ClientsLedger {
		return &ledger{}
	}),
)

type ledger struct{}
