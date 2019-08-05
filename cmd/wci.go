/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/nicomo/oclcapis"

	"github.com/spf13/cobra"
)

// wciCmd represents the wci command
var wciCmd = &cobra.Command{
	Use:   "wci [input file] [OPTIONS]",
	Short: "Retrieve Worldcat Identities infos for a bunch of LCCNs",
	Long: `Retrieve a Worldcat Identities infos 
	from LCCNs.

	Input should be a file with 1 LCCN per line, prefixed with "lccn-", e.g.
	lccn-n00000400
	lccn-n00000522
	lccn-n00000608
	etc.
	
	Result is a json file.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wci called")

		// open input file
		f, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("could not read input file: %v", err)
		}
		defer f.Close()

		// read line by line and add to input slice
		var input []string
		fScanner := bufio.NewScanner(f)
		for fScanner.Scan() {
			s := fScanner.Text()

			// checks for empty lines and whitespaces
			if len(s) == 0 {
				log.Fatalf("file should not have empty lines")
			}
			for _, runeValue := range s {
				if unicode.IsSpace(runeValue) {
					log.Fatalf("lines should not have spaces, see %s\n", s)
				}
			}
			input = append(input, s)
		}

		// call WS
		res, err := oclcapis.WCIBatchRead(input)
		if err != nil {
			log.Fatalln(err)
		}

		// create output file
		ofile, err := os.Create(ofname + ".json")
		if err != nil {
			log.Fatalf("coud not create output file: %v", err)
		}
		defer ofile.Close()

		// json encode result to file
		encoder := json.NewEncoder(ofile)
		encoder.SetIndent("", "	")
		err = encoder.Encode(res)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(wciCmd)
}
