package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info    *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning *log.Logger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error   *log.Logger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func doSomething() {
	Info.Println("** begin")
	defer Info.Println("** end")

	Info.Println(">> INFO")
	Warning.Println(">>WARNUNG")
	Error.Println("ERROR")
}
