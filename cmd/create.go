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

var createCmd = &cobra.Command{
	Use:   "create [type] [name]",
	Short: "Create a new custom template page",
	Long: `Create a new template of the predefined type and assign it a name. The type
determines how the template will be used, and the assigned name should be
used on the command line to identify the template. Possible types include:

blueprint - used to initialize pages. Creates an html file in the 
			templates/blueprints directory.
template  - a blank template in the templates directory. Simply creates an
			html file in the correct location to be found by the compiler
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
		if err := bh.Create(args[0], args[1]); err != nil {
			println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
