// Copyright © 2022 Weald Technology Trading
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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	chainqueues "github.com/wealdtech/ethdo/cmd/chain/queues"
)

var chainQueuesCmd = &cobra.Command{
	Use:   "queues",
	Short: "Show chain queues",
	Long: `Show beacon chain activation and exit queues.  For example:

    ethdo chain queues

In quiet mode this will return 0 if the entry and exit queues are 0, otherwise 1.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		res, err := chainqueues.Run(cmd)
		if err != nil {
			return err
		}
		if viper.GetBool("quiet") {
			return nil
		}
		if res != "" {
			fmt.Println(res)
		}
		return nil
	},
}

func init() {
	chainCmd.AddCommand(chainQueuesCmd)
	chainFlags(chainQueuesCmd)
	chainQueuesCmd.Flags().String("epoch", "", "epoch for which to fetch the queues")
}

func chainQueuesBindings(cmd *cobra.Command) {
	if err := viper.BindPFlag("epoch", cmd.Flags().Lookup("epoch")); err != nil {
		panic(err)
	}
}
