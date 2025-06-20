package runtime

import (
	"fmt"

	"github.com/cryptellation/ticks/pkg/tick"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// OnInitCallbackWorkflowParams is the parameters of the
// OnInitCallbackWorkflow callback workflow.
type OnInitCallbackWorkflowParams struct {
	Context Context
}

// OnNewPricesCallbackWorkflowParams is the parameters of the
// OnNewPricesCallbackWorkflow callback workflow.
type OnNewPricesCallbackWorkflowParams struct {
	Context Context
	Ticks   []tick.Tick
}

// OnExitCallbackWorkflowParams is the parameters of the
// OnExitCallbackWorkflow callback workflow.
type OnExitCallbackWorkflowParams struct {
	Context Context
}

// Runnable is the interface for a struct that can be run on Cryptellation.
type Runnable interface {
	Name() string
	OnInit(ctx workflow.Context, params OnInitCallbackWorkflowParams) error
	OnNewPrices(ctx workflow.Context, params OnNewPricesCallbackWorkflowParams) error
	OnExit(ctx workflow.Context, params OnExitCallbackWorkflowParams) error
}

// RegisterRunnable registers a runnable to a worker and returns the callbacks.
func RegisterRunnable(w worker.Worker, taskQueue string, name string, r Runnable) Callbacks {
	// Register OnInitCallback
	onInitCallbackWorkflowName := fmt.Sprintf("%s-OnInit", name)
	w.RegisterWorkflowWithOptions(r.OnInit, workflow.RegisterOptions{
		Name: onInitCallbackWorkflowName,
	})

	// Register OnNewPricesCallback
	onNewPricesCallbackWorkflowName := fmt.Sprintf("%s-OnNewPrices", name)
	w.RegisterWorkflowWithOptions(r.OnNewPrices, workflow.RegisterOptions{
		Name: onNewPricesCallbackWorkflowName,
	})

	// Register OnExitCallback
	onExitCallbackWorkflowName := fmt.Sprintf("%s-OnExit", name)
	w.RegisterWorkflowWithOptions(r.OnExit, workflow.RegisterOptions{
		Name: onExitCallbackWorkflowName,
	})

	return Callbacks{
		OnInitCallback: CallbackWorkflow{
			Name:          onInitCallbackWorkflowName,
			TaskQueueName: taskQueue,
		},
		OnNewPricesCallback: CallbackWorkflow{
			Name:          onNewPricesCallbackWorkflowName,
			TaskQueueName: taskQueue,
		},
		OnExitCallback: CallbackWorkflow{
			Name:          onExitCallbackWorkflowName,
			TaskQueueName: taskQueue,
		},
	}
}
