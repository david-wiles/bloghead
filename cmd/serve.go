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
	"net/http"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a file server using the output directory as root, listening on localhost:8081",
	Run: func(cmd *cobra.Command, args []string) {
		bh := internal.FromEnv()
		if err := http.ListenAndServe(":8081", http.FileServer(http.Dir(bh.Output))); err != nil {
			println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
