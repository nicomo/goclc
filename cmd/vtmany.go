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
	Use:   "vtmany [input file] [OPTIONS]",
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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vtmany called")
		getMany("vtmany", args[0])
	},
}

func init() {

	//	vtmanyCmd.Flags().StringVarP(&vtmof, "output", "o", "vtmany-output", "output filename, without extension, e.g. my-output-file")
	rootCmd.AddCommand(vtmanyCmd)
}
