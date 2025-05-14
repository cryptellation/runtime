package bot

import (
	"fmt"

	"github.com/cryptellation/runtime"
	"github.com/google/uuid"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

// RegisterWorkflows registers a bot workflows to a worker.
func RegisterWorkflows(w worker.Worker, taskQueue string, id uuid.UUID, bot Bot) runtime.Callbacks {
	// Register OnInitCallback
	onInitCallbackWorkflowName := fmt.Sprintf("OnInit-%s", id.String())
	w.RegisterWorkflowWithOptions(bot.OnInit, workflow.RegisterOptions{
		Name: onInitCallbackWorkflowName,
	})

	// Register OnNewPricesCallback
	onNewPricesCallbackWorkflowName := fmt.Sprintf("OnNewPrices-%s", id.String())
	w.RegisterWorkflowWithOptions(bot.OnNewPrices, workflow.RegisterOptions{
		Name: onNewPricesCallbackWorkflowName,
	})

	// Register OnExitCallback
	onExitCallbackWorkflowName := fmt.Sprintf("OnExit-%s", id.String())
	w.RegisterWorkflowWithOptions(bot.OnExit, workflow.RegisterOptions{
		Name: onExitCallbackWorkflowName,
	})

	return runtime.Callbacks{
		OnInitCallback: runtime.CallbackWorkflow{
			Name:          onInitCallbackWorkflowName,
			TaskQueueName: taskQueue,
		},
		OnNewPricesCallback: runtime.CallbackWorkflow{
			Name:          onNewPricesCallbackWorkflowName,
			TaskQueueName: taskQueue,
		},
		OnExitCallback: runtime.CallbackWorkflow{
			Name:          onExitCallbackWorkflowName,
			TaskQueueName: taskQueue,
		},
	}
}
