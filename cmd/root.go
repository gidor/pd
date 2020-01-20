package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var configName, parName, parValue, basePath, title, message string
var convPath bool

func init() {
	rootCmd.PersistentFlags().StringVarP(&configName, "cfg", "c", "cfg", "configuration name ")
	rootCmd.PersistentFlags().StringVarP(&parName, "name", "n", "", "parameter name")
	rootCmd.PersistentFlags().BoolVarP(&convPath, "path", "p", false, "convert path delimiter")
	rootCmd.PersistentFlags().StringVarP(&parValue, "value", "v", "", "parameter value")

	rootCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "dialog title")
	rootCmd.PersistentFlags().StringVarP(&message, "msg", "m", "", "message to display")

	basePath = rootPath()
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func rootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)
	return dir
}

func setParam(name string, value string) {
	if len(configName) == 0 {
		configName = "cfg"
	}
	if convPath {
		value = strings.ReplaceAll(value, "\\", "/")
	}
	confile := path.Join(basePath, configName+".json")
	fmt.Println(confile)
	var cfg map[string]interface{}

	if _, err := os.Stat(confile); err == nil {
		// path/to/whatever exists
		dat, err := ioutil.ReadFile(confile)
		check(err)
		// fmt.Print(string(dat))
		err = json.Unmarshal(dat, &cfg)
		check(err)
		cfg[name] = value

	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		cfg = make(map[string]interface{})
		cfg[name] = value

	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		fmt.Println(err)
	}
	// dat, err := json.Marshal(cfg)
	dat, err := json.MarshalIndent(cfg, "", "  ")
	check(err)
	ioutil.WriteFile(confile, dat, 0666)
	// ioutil.WriteFile(configName+".json", dat, os.ModePerm)

}

var rootCmd = &cobra.Command{
	Use:   "pd",
	Short: "pd Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(parName) > 0 {
			setParam(parName, parValue)
		}
		// Do Stuff Here
		// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(dir)
		// // len := len(os.Args)
		// for i, j := 0, len(os.Args); i < j; i++ {
		// 	fmt.Println(i, os.Args[i])
		// }
		// fmt.Println(dialog.Message("pipo").YesNo())
		// f, err := dialog.File().Title("sqlite").Filter("Sqlite", "sqlite,sqlite3").Filter("All files", "*").Load()
		// if err != nil {
		// 	fmt.Println("ERROR", err)
		// } else {
		// 	fmt.Println(f)
		// }
		// f, err = dialog.Directory().Title("dir").Browse()
		// if err != nil {
		// 	fmt.Println("ERROR", err)
		// } else {
		// 	fmt.Println(f)
		// }
	},
}

// Execute  execute comamdd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
