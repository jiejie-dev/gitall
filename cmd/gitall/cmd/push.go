/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"strings"

	"github.com/jerloo/gitall"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Perform git push command of multiple repositories in batch.",
	Run: func(cmd *cobra.Command, args []string) {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		workspace = strings.ReplaceAll(workspace, "$HOME", home)

		client, err := gitall.NewGitAllClient(workspace, gitall.WithVerbose(verbose))
		cobra.CheckErr(err)

		err = client.Push()
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
