// main
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func grep(number int, line, searchString string, numberFlag bool) {
	if strings.Contains(line, searchString) {
		if numberFlag {
			fmt.Printf("%v: %v", number, line)
		} else {
			fmt.Printf("%v", line)
		}
	}
}

func grepFile(file *os.File, searchString string, numberFlag bool) {

	reader := bufio.NewReader(file)

	number := 0

	for {
		number += 1
		line, err := reader.ReadString('\n')

		switch {
		case err == io.EOF:
			grep(number, line, searchString, numberFlag) // last line without ending '/n'
			return
		case err != nil:
			log.Fatal(err)
		}

		grep(number, line, searchString, numberFlag)
	}
}

func main() {
	// suppress date and time in front of log messages
	log.SetFlags(log.Lshortfile)

	// parse cmd. line
	var searchString string
	flag.StringVar(&searchString, "search", "", "string to be searched in file")

	var numberFlag bool
	flag.BoolVar(&numberFlag, "number", false, "print line number in front of found lines")

	flag.Parse()

	if searchString == "" {
		log.Fatal("missing cmd. line option '--search'")
	}

	// open file
	file, err := os.Open("./datei.txt")

	if err != nil {
		log.Fatal(err) // calls os.Exit(1)
	}

	// close file
	defer file.Close()

	grepFile(file, searchString, numberFlag)
}
