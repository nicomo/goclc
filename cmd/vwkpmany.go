/*
Copyright © 2019 Nicolas Morin <nicolas.morin@gmail.com>

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
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/nicomo/oclcapis"
	"github.com/spf13/cobra"
)

// vwkpmanyCmd represents the vwkpmany command
var vwkpmanyCmd = &cobra.Command{
	Use:   "vwkpmany [input file] [OPTIONS]",
	Short: "Retrieve Wikidata IDs for a bunch of VIAF IDs.",
	Long: `Retrieve a bunch of wikidata IDs 
	from Viaf IDs.

	Input should be a file with 1 viaf ID per line, e.g. 
	10017193
	10020421
	10011312
	etc.
	
	Result is a csv with 1 result per line, like so:
	10017193,could not find  10017193
	10020421,Q30084598
	10011312,
	etc.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vwkpmany called")

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
		res, err := oclcapis.ViafGetWKPs(input)
		if err != nil {
			log.Fatalln(err)
		}

		// prepare result for csv
		output := [][]string{}
		for k := range res {
			output = append(output, []string{k, res[k]})
		}

		// write result to file
		file, err := os.Create(ofname + ".csv")
		if err != nil {
			log.Fatalf("cannot create file: %v\n", err)
		}
		defer file.Close()

		w := csv.NewWriter(file)

		for _, o := range output {
			if err := w.Write(o); err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
		}

		// Write any buffered data to the underlying writer
		w.Flush()

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(vwkpmanyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vwkpmanyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vwkpmanyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}