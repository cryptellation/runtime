package runtime

import (
	"time"

	"github.com/google/uuid"
)

// Run is the context of a run that contains several information about the run.
type Run struct {
	// ID is the ID of the run that corresponds to a backtest, a forwardtest or a live run.
	ID uuid.UUID
	// Mode is the mode of the run (backtests, forwardtest or live run).
	Mode Mode
	// Now is the time of the run.
	// It is used to simulate the time of the run for backtests.
	// On forwardtests and live runs, it is the current time.
	Now time.Time
	// ParentTaskQueue is the task queue of the parent workflow.
	// It's the one that called the callback.
	ParentTaskQueue string
}
