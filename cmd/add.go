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

var blueprint string = ""

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [type] [page name]",
	Short: "Add a new page to your site",
	Long: `Add a new page to the site using a predetermined type by creating necessary 
html and data files. The type will determine how the page is compiled, 
and the name should be an internal name for the page. Possible types are:

page    - a standalone page
article - a blog post used in a sequence of posts
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires type and name arguments")
		} else if len(args) == 1 {
			return errors.New("Requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		bh := internal.FromEnv()
		if err := bh.Add(args[0], args[1], blueprint); err != nil {
			println(err.Error())
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&blueprint, "blueprint", "b", "", "Specify the blueprint to initialize the page with")
	rootCmd.AddCommand(addCmd)
}
