package cmd

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/gen2brain/dlgs"
	"github.com/sqweek/dialog"
)

// var title string

func init() {
	// dirCmd.Flags().StringVarP(&title, "title", "t", "", "title for directory")
	rootCmd.AddCommand(dirCmd)
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "manage dir parameter ",
	Long:  `manage dir parameter `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		// check(err)
		// fmt.Println(dir)
		// len := len(os.Args)
		// for i, j := 0, len(os.Args); i < j; i++ {
		// 	fmt.Println(i, os.Args[i])
		// }
		if runtime.GOOS == "windows" {
			if len(parName) > 0 {
				f, err := dialog.Directory().Title(title).Browse()
				check(err)
				setPathParam(parName, f)
			}
		} else if runtime.GOOS == "linux" {
			f, ok, err := dlgs.File(title, "", true)
			check(err)
			if ok {
				setPathParam(parName, f)
			}
		}
	},
}
