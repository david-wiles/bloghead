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
	"errors"
	"github.com/david-wiles/bloghead/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [type] [blueprint name] [page name]",
	Short: "Add a new page to your site",
	Long: `A new page can be added from a bloghead default or a custom default.

The first argument is the type of page to add. Second argument is the name of the
blueprint to copy from, and the final argument is the name of the new page (the .html 
extension will be added automatically.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires type, template, and name arguments")
		} else if len(args) == 1 {
			return errors.New("Requires template and name arguments")
		} else if len(args) == 2 {
			return errors.New("Requires name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		bh := internal.FromEnv()
		if err := bh.Add(args[0], args[1], args[2]); err != nil {
			println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
