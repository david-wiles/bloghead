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
	"fmt"
	"github.com/david-wiles/bloghead/internal"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new site in the specified directory",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := ".bloghead"
		if len(args) > 0 {
			configFile = path.Join(args[0], ".bloghead")
		}

		_, err := os.Stat(configFile)
		if err != nil {
			if os.IsNotExist(err) {
				if err := internal.Init(configFile); err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		} else {
			_, _ = fmt.Fprintf(os.Stdout, "Configuration already exists at %v. Delete it to create a new one, or use "+
				"init in a different directory\n", configFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
