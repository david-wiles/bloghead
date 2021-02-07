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

var watch bool

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Builds all pages for the current project",
	Run: func(cmd *cobra.Command, args []string) {
		bh := internal.FromEnv()

		if watch {
			if err := bh.Watch(); err != nil {
				println(err.Error())
			}
		} else {
			if err := bh.Start(); err != nil {
				println(err.Error())
			}
		}
	},
}

func init() {
	publishCmd.Flags().BoolVarP(&watch, "watch", "w", false, "--watch, -w. Watch files for changes")

	rootCmd.AddCommand(publishCmd)
}
