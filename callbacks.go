package runtime

import (
	"time"

	"github.com/cryptellation/ticks/pkg/tick"
)

// Callbacks is the struct representing callbacks for ans automation through cryptellation API.
type Callbacks struct {
	OnInitCallback      CallbackWorkflow
	OnNewPricesCallback CallbackWorkflow
	OnExitCallback      CallbackWorkflow
}

// CallbackWorkflow is the parameters of a callback workflow.
type CallbackWorkflow struct {
	Name             string
	TaskQueueName    string
	ExecutionTimeout time.Duration
}

// OnInitCallbackWorkflowParams is the parameters of the
// OnInitCallbackWorkflow callback workflow.
type OnInitCallbackWorkflowParams struct {
	RunCtx Context
}

// OnNewPricesCallbackWorkflowParams is the parameters of the
// OnNewPricesCallbackWorkflow callback workflow.
type OnNewPricesCallbackWorkflowParams struct {
	Run   Context
	Ticks []tick.Tick
}

// OnExitCallbackWorkflowParams is the parameters of the
// OnExitCallbackWorkflow callback workflow.
type OnExitCallbackWorkflowParams struct {
	Run Context
}
