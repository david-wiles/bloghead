/*
Copyright © 2021 David Wiles <davd@wiles.fyi>

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
	"fmt"
	"github.com/david-wiles/bloghead/internal"
	"github.com/spf13/cobra"
)

// metaCmd represents the meta command
var metaCmd = &cobra.Command{
	Use:   "meta [key] [value]",
	Short: "Update the metadata for the site",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Must provide two arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		bh := internal.FromEnv()
		if err := bh.SetMetaValue(args[0], args[1]); err != nil {
			_, _ = fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(metaCmd)
}
