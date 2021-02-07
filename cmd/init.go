/*
Copyright © 2021 David Wiles david@wiles.fyi

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
	"github.com/david-wiles/bloghead/internal"
	"github.com/spf13/cobra"
	"path"
)

var root string
var output string

var initCmd = &cobra.Command{
	Use:   "init [folder?]",
	Short: "Initialize a new site in the specified directory",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := ".bloghead"
		if len(args) > 0 {
			configFile = path.Join(args[0], ".bloghead")
		}

		if err := internal.Init(root, output, configFile); err != nil {
			panic(err)
		}
	},
}

func init() {
	initCmd.Flags().StringVarP(&root, "root", "r", "html", "--root [directory], -r [directory]. Root directory for html files")
	initCmd.Flags().StringVarP(&output, "output", "o", "www", "--output [directory], -o [directory]. Output directory for geneated files")

	rootCmd.AddCommand(initCmd)
}
