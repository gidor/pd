package main

import "github.com/gidor/pd/cmd"

// "github.com/sqweek/dialog"

func main() {
	cmd.Execute()
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
}
