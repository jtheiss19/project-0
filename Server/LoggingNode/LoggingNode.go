package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
)

var ConnSignal chan string = make(chan string)

func main() {

	fmt.Println("Launching Logging Server...")

	ln, _ := net.Listen("tcp", ":8070")

	fmt.Println("Online - Listening on port 8070")

	for {
		go Session(ln, "8070")
		<-ConnSignal
	}
	/*
		for {
			UpdateLog("../..")
			time.Sleep(10 * time.Second)
		}
	*/
}

func Session(ln net.Listener, port string) {
	conn, _ := ln.Accept()
	defer conn.Close()

	logfile, err := os.OpenFile("log_"+conn.RemoteAddr().String()+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("New Connection On")
	Write([]byte("New Connection"), conn.LocalAddr().String(), "N/A", conn.LocalAddr().String(), logfile)
	ConnSignal <- "New Connection"

	for {
		buf := make([]byte, 1024)
		conn.Read(buf)
		for i := 0; i < len(buf); i++ {
			if buf[i] == byte('\u0000') {
				buf = append(buf[0:i])
				break
			}
		}
		if string(buf) != "" {
			logfile.Write(buf)
		}
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

var mu = &sync.Mutex{}

//Write writes to a logfile
func Write(info []byte, In string, Out string, WhoAmI string, file *os.File) {
	//Build complete String
	FullString := WhoAmI + ", " + In + ", " + Out + ", " + string(info) + "\n"
	file.Write([]byte(FullString))
}
