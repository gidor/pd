/*
Copyright © 2021 Gianni Doria  <gianni.doria@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	cj "github.com/gidor/pd/json"
	cy "github.com/gidor/pd/yaml"
	"github.com/sqweek/dialog"
)

var cwd, configName, parName, parValue, basePath, title, message string

const (
	undef   = iota
	jsoncfg = iota
	yamlcfg = iota
)

const (
	none = iota
	win  = iota
	unix = iota
)

var (
	convDelimiter int8   = none
	convPath      bool   = false
	cfgext        int8   = undef
	ver           string = "1.1"
	prod          string = "pd"
	// inputFile     string                 = ""
	// outputFile    string                 = ""
	pretty     bool                   = true
	jyout      bool                   = false
	data       map[string]interface{} = make(map[string]interface{})
	htmlescape bool                   = false
)

func format(value string) string {
	if convDelimiter == win {
		return strings.ReplaceAll(value, "/", "\\")
	}
	if convDelimiter == unix {
		return strings.ReplaceAll(value, "\\", "/")
	}
	return value
}

func init() {
	path, err := os.Getwd()
	check(err)
	cwd = path
}
func iniCfg() {

	if len(configName) == 0 {
		configName = "cfg.json"
		cfgext = jsoncfg
	} else {
		switch ext := filepath.Ext(configName); ext {
		case "":
			configName = configName + ".json"
			cfgext = jsoncfg
		case ".json", ".JSON":
			cfgext = jsoncfg
		case ".yaml", ".yml", ".YML", ".YAML":
			cfgext = yamlcfg
		default:
			cfgext = jsoncfg
		}
	}
	switch cfgext {
	case jsoncfg:
		iniJson(configName)
	case yamlcfg:
		iniYaml(configName)
	}
}

func finalizeCfg() {
	switch cfgext {
	case jsoncfg:
		finJson(configName)
	case yamlcfg:
		finYaml(configName)
	}
	if jyout {
		switch cfgext {
		case jsoncfg:
			cf := strings.ReplaceAll(configName, ".json", ".yaml")
			cf = strings.ReplaceAll(cf, ".JSON", ".yaml")
			finYaml(cf)
		case yamlcfg:
			cf := strings.ReplaceAll(configName, ".yaml", ".json")
			cf = strings.ReplaceAll(cf, ".YAML", ".json")
			cf = strings.ReplaceAll(cf, ".yml", ".json")
			cf = strings.ReplaceAll(cf, ".YML", ".json")
			finJson(cf)
		}

	}
}

func iniJson(confname string) {
	patf := path.Join(cwd, confname)
	if _, err := os.Stat(confname); err == nil {
		in, err := os.OpenFile(patf, os.O_RDONLY, 0755)
		check(err)
		din, err := cj.Load(in)
		in.Close()
		check(err)
		for k, v := range din {
			data[k] = v
		}
	}
}

func iniYaml(confname string) {
	patf := path.Join(cwd, confname)
	if _, err := os.Stat(patf); err == nil {
		in, err := os.OpenFile(patf, os.O_RDONLY, 0755)
		check(err)
		din, err := cy.Load(in)
		in.Close()
		check(err)
		for k, v := range din {
			data[k] = v
		}
	}
}

func finJson(confname string) {
	patf := path.Join(cwd, confname)

	out, err := os.Create(patf)
	check(err)
	err = cj.Save(&data, out, pretty, htmlescape)
	out.Close()
	check(err)
}

func finYaml(confname string) {
	patf := path.Join(cwd, confname)
	out, err := os.Create(patf)
	check(err)
	err = cy.Save(&data, out, pretty, htmlescape)
	out.Close()
	check(err)
}

// dlg show message
func dlgMessage(message string, title string) {
	dlg := dialog.Message("%s", message)
	dlg.Title(title)
	dlg.Info()
}

func dlgYesNo(message string, title string) bool {
	dlg := dialog.Message("%s", message)
	dlg.Title(title)
	return dlg.YesNo()
}

// dlg_ask for file using dialog
func dlgFile(title string) (string, error) {
	dlg := dialog.File()
	dlg.Title(title)
	dlg.Filter("All Files", "*")
	file, err := dlg.Load()
	return file, err
}

func dlgFileF(title string, filter string) (string, error) {
	dlg := dialog.File()
	dlg.Title(title)
	dlg.Filter("All Files", filter)
	file, err := dlg.Save()
	return file, err
}

func dlgFileFP(title string, filter string, path string) (string, error) {
	dlg := dialog.File()
	dlg.Title(title)
	dlg.Filter("All Files", filter)
	dlg.StartDir = path
	file, err := dlg.Save()
	return file, err
}

func dlgFileP(title string, path string) (string, error) {
	dlg := dialog.File()
	dlg.Title(title)
	dlg.Filter("All Files", "*")
	dlg.StartDir = path
	file, err := dlg.Save()
	return file, err
}

func dlgDir(title string) (string, error) {
	dialog.Message("%s", message).Title("Hello world!").Info()
	file, err := dialog.File().Title("Save As").Filter("All Files", "*").Save()
	return file, err
}

func promptMessage(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	return text, err
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func rootPath() string {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// check(err)
	return cwd
}

func setPathParam(name string, value string) {
	base := filepath.Base(value)
	ext := filepath.Ext(value)
	dir := filepath.Dir(value)
	fname := base[0 : len(base)-len(ext)]
	setParam(name, value)
	setParam(name+"_base", base)
	setParam(name+"_bsase_dir", filepath.Base(dir))
	setParam(name+"_ext", ext)
	setParam(name+"_dir", dir)
	setParam(name+"_fname", fname)
}

func setParam(name string, value string) {
	data[name] = format(value)
}

// func setp(path string, value string) {
// 	names := strings.Split(path, ".")
// 	var dest *interface{}
// 	dest = data // interface{}(data)
// 	j := len(names)
// 	for i := 0; i < j; i++ {
// 		name := names[i]

// 		switch dest.(type) {
// 		case map[string]interface{}:
// 			if v, ok = s.(map[string]interface{})[key]; !ok {
// 				err = fmt.Errorf("Key not present. [Key:%s]", key)
// 			}
// 		case []interface{}:
// 			if i, err = strconv.ParseInt(key, 10, 64); err == nil {
// 				array := s.([]interface{})
// 				if int(i) < len(array) {
// 					v = array[i]
// 				} else {
// 					err = fmt.Errorf("Index out of bounds. [Index:%d] [Array:%v]", i, array)
// 				}
// 			}
// 		}

// 		// switch dest.(type) {
// 		// case nil:
// 		// 	return "", nil
// 		// 	// return "null"
// 		// case bool:
// 		// 	if v {
// 		// 		return "true", nil
// 		// 	} else {
// 		// 		return "false", nil
// 		// 	}
// 		// case int:
// 		// 	return strconv.FormatInt(int64(v), 10), nil
// 		// case float64:
// 		// 	return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
// 		// // case *big.Int:
// 		// // 	e.w.Write(v.Append(e.buf[:0], 10))
// 		// case string:
// 		// 	return v, nil
// 		// case []interface{}:
// 		// 	b, err := json.Marshal(v)
// 		// 	return string(b), err
// 		// case map[string]interface{}:
// 		// 	b, err := json.Marshal(v)
// 		// 	return string(b), err
// 		// default:
// 		// 	return "", errors.New(fmt.Sprintf("invalid value: %v", v))
// 		// }

// 	}
// 	if convPath {
// 		value = strings.ReplaceAll(value, "\\", "/")
// 	}
// 	// data[name] = value
// }

// func get(key string, s interface{}) (v interface{}, err error) {
// 	// var (
// 	// 	i  int64
// 	// 	ok bool
// 	// )
// 	switch s.(type) {
// 	case map[string]interface{}:
// 		if v, ok = s.(map[string]interface{})[key]; !ok {
// 			err = fmt.Errorf("Key not present. [Key:%s]", key)
// 		}
// 	case []interface{}:
// 		if i, err = strconv.ParseInt(key, 10, 64); err == nil {
// 			array := s.([]interface{})
// 			if int(i) < len(array) {
// 				v = array[i]
// 			} else {
// 				err = fmt.Errorf("Index out of bounds. [Index:%d] [Array:%v]", i, array)
// 			}
// 		}
// 		// case Signature:
// 		// 	r := reflect.ValueOf(s)
// 		// 	v = reflect.Indirect(r).FieldByName(key)
// 	}
// 	//fmt.Println("Value:", v, " Key:", key, "Error:", err)
// 	return v, err
// }

func getParam(name string) string {
	val, ok := data[name].(string)
	if ok {
		if convPath {
			val = strings.ReplaceAll(val, "\\", "/")
		}
		return val
	}
	return ""
}
