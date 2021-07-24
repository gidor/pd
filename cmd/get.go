/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get <parameter>",
	Short: "display vaue from cofiguration ",
	Long: `display vaue from cofiguration
		ask question to user and store `,
	Run: func(cmd *cobra.Command, args []string) {
		iniCfg()
		val := getParam(args[0])
		fmt.Print(val)
	},
}
