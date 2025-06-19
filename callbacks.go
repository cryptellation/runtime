package runtime

import (
	"errors"
	"fmt"
	"time"

	"github.com/cryptellation/ticks/pkg/tick"
)

var (
	// ErrEmptyWorkflowName is returned when the workflow name is empty.
	ErrEmptyWorkflowName = errors.New("workflow name is empty")
	// ErrEmptyTaskQueueName is returned when the task queue name is empty.
	ErrEmptyTaskQueueName = errors.New("task queue name is empty")
)

// Callbacks is the struct representing callbacks for ans automation through cryptellation API.
type Callbacks struct {
	OnInitCallback      CallbackWorkflow
	OnNewPricesCallback CallbackWorkflow
	OnExitCallback      CallbackWorkflow
}

// Validate validates the callbacks.
func (c Callbacks) Validate() error {
	if err := c.OnInitCallback.Validate(); err != nil {
		return fmt.Errorf("onInitCallback validation failed: %w", err)
	}
	if err := c.OnNewPricesCallback.Validate(); err != nil {
		return fmt.Errorf("onNewPricesCallback validation failed: %w", err)
	}
	if err := c.OnExitCallback.Validate(); err != nil {
		return fmt.Errorf("onExitCallback validation failed: %w", err)
	}
	return nil
}

// CallbackWorkflow is the parameters of a callback workflow.
type CallbackWorkflow struct {
	Name             string
	TaskQueueName    string
	ExecutionTimeout time.Duration
}

// Validate validates the callback workflow.
func (cw CallbackWorkflow) Validate() error {
	if cw.Name == "" {
		return ErrEmptyWorkflowName
	}
	if cw.TaskQueueName == "" {
		return ErrEmptyTaskQueueName
	}
	return nil
}

// OnInitCallbackWorkflowParams is the parameters of the
// OnInitCallbackWorkflow callback workflow.
type OnInitCallbackWorkflowParams struct {
	Run Run
}

// OnNewPricesCallbackWorkflowParams is the parameters of the
// OnNewPricesCallbackWorkflow callback workflow.
type OnNewPricesCallbackWorkflowParams struct {
	Run   Run
	Ticks []tick.Tick
}

// OnExitCallbackWorkflowParams is the parameters of the
// OnExitCallbackWorkflow callback workflow.
type OnExitCallbackWorkflowParams struct {
	Run Run
}
