package task

import (
	"errors"
	"fmt"
)

var (
	// ErrTaskfileAlreadyExists is returned on creating a Taskfile if one already exists
	ErrTaskfileAlreadyExists = errors.New("task: A Taskfile already exists")
)

type taskNotFoundError struct {
	taskName string
}

func (err *taskNotFoundError) Error() string {
	return fmt.Sprintf(`task: Task "%s" not found`, err.taskName)
}

type taskRunError struct {
	taskName string
	err      error
}

func (err *taskRunError) Error() string {
	return fmt.Sprintf(`task: Failed to run task "%s": %v`, err.taskName, err.err)
}

// MaximumTaskCallExceededError is returned when a task is called too
// many times. In this case you probably have a cyclic dependendy or
// infinite loop
type MaximumTaskCallExceededError struct {
	task string
}

func (e *MaximumTaskCallExceededError) Error() string {
	return fmt.Sprintf(
		`task: maximum task call exceeded (%d) for task "%s": probably an cyclic dep or infinite loop`,
		MaximumTaskCall,
		e.task,
	)
}

// DepCycleError is used when for example A depends on B depends on A.
type DepCycleError struct {
	task1 string
	task2 string
}

func (e DepCycleError) Error() string {
	return fmt.Sprintf("cyclic dependency detected: task %s depends on %s, which depends on %s", e.task1, e.task2, e.task1)
}
