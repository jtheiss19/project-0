package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		UpdateLog("../..")
		time.Sleep(10 * time.Second)
	}
}

//RecursiveCompile serches through all non hidden files
//and grads all log.txt documents to compile them into
//one large log.txt
func RecursiveCompile(file string) []string {
	var Textfiles []string
	files, err := ioutil.ReadDir(file)
	if err != nil {
		return nil
	}
	for _, v := range files {
		if strings.HasPrefix(v.Name(), ".") {
		} else {
			NewFile := RecursiveCompile(file + "/" + v.Name())
			for _, v := range NewFile {
				Textfiles = append(Textfiles, v)
			}
		}
		if strings.Contains(v.Name(), "log.txt") {
			Textfiles = append(Textfiles, file+"/"+v.Name())
		}
	}
	return Textfiles
}

//UpdateLog takes in a top directory and serches through all
//sub directories. It then saves all these files into Log.txt
func UpdateLog(FilePath string) {
	Textfiles := RecursiveCompile(FilePath)

	File, Error := os.OpenFile("Log.txt", os.O_RDWR|os.O_CREATE, 7777)

	if Error != nil {
		fmt.Println(Error)
	}

	w := bufio.NewWriter(File)

	for _, f := range Textfiles {
		logdata, err := os.OpenFile(f, os.O_RDWR, 7777)
		if err != nil {
			fmt.Println(err)
		}

		Reader := bufio.NewScanner(logdata)
		for i := 0; i >= 0; i++ {
			if Reader.Scan() == false {
				break
			}
			w.Write(append([]byte(Reader.Text()), byte('\n')))
		}
		logdata.Truncate(0)
	}

	w.Flush()

}
