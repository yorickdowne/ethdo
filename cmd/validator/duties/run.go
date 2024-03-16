// Copyright © 2019 - 2024 Weald Technology Trading.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validatorduties

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run runs the wallet create data command.
func Run(cmd *cobra.Command) (string, error) {
	ctx := context.Background()
	dataIn, err := input(ctx)
	if err != nil {
		return "", errors.Join(errors.New("failed to set up command"), err)
	}

	// Further errors do not need a usage report.
	cmd.SilenceUsage = true

	dataOut, err := process(ctx, dataIn)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			return "", errors.New("operation timed out; try increasing with --timeout option")
		default:
			return "", errors.Join(errors.New("failed to process"), err)
		}
	}

	if viper.GetBool("quiet") {
		return "", nil
	}

	results, err := output(ctx, dataOut)
	if err != nil {
		return "", errors.Join(errors.New("failed to obtain output"), err)
	}

	return results, nil
}
