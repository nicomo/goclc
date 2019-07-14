// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/nicomo/oclcapis"
	"github.com/spf13/cobra"
)

// viafGetidsCmd represents the viafGetids command
var viafGetidsCmd = &cobra.Command{
	Use:   "viafGetids",
	Short: "Finds other IDs from a VIAF ID",
	Long: `Finds all other available sources IDs, e.g. LC, DNB, WKP, etc.
	from a VIAF ID provided as a string.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("viafGetids called")
		res, err := oclcapis.ViafGetIDs(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("input: %s, result: %v\n", args[0], res)
	},
}

func init() {
	rootCmd.AddCommand(viafGetidsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viafGetidsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viafGetidsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
