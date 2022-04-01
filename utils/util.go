package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	projectFolder = "."
)

func LogAndExit(v ...interface{}) {
	v = append([]interface{}{"O:"}, v...)
	log.Println(v...)
	os.Exit(15)
}

func ListFile(folder string, fun func(string)) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			d := folder + "/" + file.Name()
			fun(d)
			ListFile(d, fun)
		}
	}
}

func LitDirs(d string, dirs *[]string) bool {
	d += "/"
	for _, v := range *dirs {
		if strings.HasPrefix(d, projectFolder+"/"+v+"/") {
			return true
		}
	}
	return false
}

func DirParse2Array(s string) []string {
	a := strings.Split(s, ",")
	r := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if ss := strings.Trim(a[i], " "); ss != "" {
			r = append(r, ss)
		}
	}
	return r
}
