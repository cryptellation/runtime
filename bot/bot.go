package bot

import (
	"github.com/cryptellation/runtime"
	"go.temporal.io/sdk/workflow"
)

// Bot is the interface for a trading robot executed on Cryptellation.
type Bot interface {
	OnInit(ctx workflow.Context, params runtime.OnInitCallbackWorkflowParams) error
	OnNewPrices(ctx workflow.Context, params runtime.OnNewPricesCallbackWorkflowParams) error
	OnExit(ctx workflow.Context, params runtime.OnExitCallbackWorkflowParams) error
}
