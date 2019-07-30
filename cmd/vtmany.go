// Copyright © 2019 Nicolas Morin <nicolas.morin@gmail.com>
//
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
)

// vtmanyCmd represents the vtmany command
var vtmanyCmd = &cobra.Command{
	Use:   "vtmany [input file]",
	Short: "Translate many source IDs to VIAF IDs.",
	Long: `
	Translate a bunch of Source IDs 
	(identifier for an original source record at a specific institution,
	e.g. DNB, Sudoc, etc)
	to VIAF URIs.

	Input should be a file with 1 url encoded string per line, e.g. 
	SUDOC%7c033522448
	DNB%7c123456
	etc.
	
	Result is a csv with 1 result per line, like so:
	SUDOC%7c033522448,VIAFID#1
	DNB%7c123456,VIAFID#2
	etc.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vtmany called")
	},
}

func init() {
	rootCmd.AddCommand(vtmanyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vtmanyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vtmanyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// vtmanyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
