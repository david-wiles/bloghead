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
)

var createCmd = &cobra.Command{
	Use:   "create [type] [name]",
	Short: "Create a new custom template page",
	Long: `New templates must be of a predefined type. 

The type argument specifies the type of template and the name argument
specifies the name of the template. When using the template to create
a new page, the name will be used to specify the template.`,
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
