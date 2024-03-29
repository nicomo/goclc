// Copyright © 2019 Nicolas Morin <nicolas.morin@gmail.com>
/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vlcnmanyCmd represents the vlcnmany command
var vlcnmanyCmd = &cobra.Command{
	Use:   "vlcnmany [input file] [OPTIONS]",
	Short: "Retrieve Library of Congress IDs (LCN) for a bunch of VIAF IDs.",
	Long: `Retrieve a bunch of Library of Congress IDs 
	from Viaf IDs.

	Input should be a file with 1 viaf ID per line, e.g. 
	10017193
	10020421
	10011312
	etc.
	
	Result is a csv with 1 result per line, like so:
	10017193,could not find a LC Number for 10017193
	10020421,n2001113638
	10011312,n00037379
	etc.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vlcnmany called")
		getMany("vlcnmany", args[0])
	},
}

func init() {
	//	vtmanyCmd.Flags().StringVarP(&vlmany, "output", "o", "vlcnmany-output", "output filename, without extension, e.g. my-output-file")
	rootCmd.AddCommand(vlcnmanyCmd)
}
