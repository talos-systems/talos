// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config

import (
	"log"
	"os"

	"github.com/talos-systems/talos/internal/app/machined/internal/phase"
	"github.com/talos-systems/talos/internal/pkg/runtime"
)

// ExtraEnvVars represents the ExtraEnvVars task.
type ExtraEnvVars struct{}

// NewExtraEnvVarsTask initializes and returns an ExtraEnvVars task.
func NewExtraEnvVarsTask() phase.Task {
	return &ExtraEnvVars{}
}

// TaskFunc returns the runtime function.
func (task *ExtraEnvVars) TaskFunc(mode runtime.Mode) phase.TaskFunc {
	return task.runtime
}

func (task *ExtraEnvVars) runtime(r runtime.Runtime) (err error) {
	for key, val := range r.Config().Machine().Env() {
		if err = os.Setenv(key, val); err != nil {
			log.Printf("WARNING failed to set enivronment variable: %v", err)
		}
	}

	return nil
}
