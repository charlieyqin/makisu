package step

import "fmt"

// CopyStep is similar to add, so they depend on a common base.
type CopyStep struct {
	*addCopyStep
}

// NewCopyStep creates a new CopyStep.
func NewCopyStep(args, chown, fromStage string, fromPaths []string, toPath string, commit bool) (*CopyStep, error) {
	s, err := newAddCopyStep(Copy, args, chown, fromStage, fromPaths, toPath, commit)
	if err != nil {
		return nil, fmt.Errorf("new add/copy step: %s", err)
	}
	return &CopyStep{s}, nil
}
