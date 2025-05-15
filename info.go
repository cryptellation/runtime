package runtime

import (
	"time"

	"github.com/google/uuid"
)

// RunInfo is the context of a run that contains several information about the run.
type RunInfo struct {
	ID        uuid.UUID
	Mode      Mode
	Now       time.Time
	TaskQueue string
}
