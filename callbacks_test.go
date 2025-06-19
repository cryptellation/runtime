package runtime

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestCallbacksSuite(t *testing.T) {
	suite.Run(t, new(CallbacksSuite))
}

type CallbacksSuite struct {
	suite.Suite
}

func (suite *CallbacksSuite) TestCallbackWorkflowValidate() {
	// Test valid callback workflow
	validCallback := CallbackWorkflow{
		Name:          "test-workflow",
		TaskQueueName: "test-queue",
	}
	err := validCallback.Validate()
	suite.Require().NoError(err)

	// Test empty workflow name
	emptyNameCallback := CallbackWorkflow{
		Name:          "",
		TaskQueueName: "test-queue",
	}
	err = emptyNameCallback.Validate()
	suite.Require().Error(err)
	suite.Require().Equal(ErrEmptyWorkflowName, err)

	// Test empty task queue name
	emptyQueueCallback := CallbackWorkflow{
		Name:          "test-workflow",
		TaskQueueName: "",
	}
	err = emptyQueueCallback.Validate()
	suite.Require().Error(err)
	suite.Require().Equal(ErrEmptyTaskQueueName, err)

	// Test both empty
	emptyBothCallback := CallbackWorkflow{
		Name:          "",
		TaskQueueName: "",
	}
	err = emptyBothCallback.Validate()
	suite.Require().Error(err)
	suite.Require().Equal(ErrEmptyWorkflowName, err) // First error encountered
}

func (suite *CallbacksSuite) TestCallbacksValidate() {
	// Test valid callbacks
	validCallbacks := Callbacks{
		OnInitCallback: CallbackWorkflow{
			Name:          "init-workflow",
			TaskQueueName: "test-queue",
		},
		OnNewPricesCallback: CallbackWorkflow{
			Name:          "prices-workflow",
			TaskQueueName: "test-queue",
		},
		OnExitCallback: CallbackWorkflow{
			Name:          "exit-workflow",
			TaskQueueName: "test-queue",
		},
	}
	err := validCallbacks.Validate()
	suite.Require().NoError(err)

	// Test invalid OnInitCallback
	invalidInitCallbacks := Callbacks{
		OnInitCallback: CallbackWorkflow{
			Name:          "", // Invalid
			TaskQueueName: "test-queue",
		},
		OnNewPricesCallback: CallbackWorkflow{
			Name:          "prices-workflow",
			TaskQueueName: "test-queue",
		},
		OnExitCallback: CallbackWorkflow{
			Name:          "exit-workflow",
			TaskQueueName: "test-queue",
		},
	}
	err = invalidInitCallbacks.Validate()
	suite.Require().Error(err)
	suite.Require().Equal(ErrEmptyWorkflowName, err)

	// Test invalid OnNewPricesCallback
	invalidPricesCallbacks := Callbacks{
		OnInitCallback: CallbackWorkflow{
			Name:          "init-workflow",
			TaskQueueName: "test-queue",
		},
		OnNewPricesCallback: CallbackWorkflow{
			Name:          "prices-workflow",
			TaskQueueName: "", // Invalid
		},
		OnExitCallback: CallbackWorkflow{
			Name:          "exit-workflow",
			TaskQueueName: "test-queue",
		},
	}
	err = invalidPricesCallbacks.Validate()
	suite.Require().Error(err)
	suite.Require().Equal(ErrEmptyTaskQueueName, err)

	// Test invalid OnExitCallback
	invalidExitCallbacks := Callbacks{
		OnInitCallback: CallbackWorkflow{
			Name:          "init-workflow",
			TaskQueueName: "test-queue",
		},
		OnNewPricesCallback: CallbackWorkflow{
			Name:          "prices-workflow",
			TaskQueueName: "test-queue",
		},
		OnExitCallback: CallbackWorkflow{
			Name:          "", // Invalid
			TaskQueueName: "test-queue",
		},
	}
	err = invalidExitCallbacks.Validate()
	suite.Require().Error(err)
	suite.Require().Equal(ErrEmptyWorkflowName, err)
}
