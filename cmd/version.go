/*
Copyright © 2021 Gianni Doria (gianni.doria@gmail.com)
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display product version",
	Long:  `Display product version.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print(fmt.Sprintf("%s version  %s\n", prod, ver))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
