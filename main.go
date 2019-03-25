package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	CWD, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while getting the current working directory : %s \n ", err)
	}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s -- $$ ", CWD)
		reader.Scan()
		text := reader.Text()
		commandSlice := strings.Split(text, " ")
		switch commandSlice[0] {
		case "ls":
			list(commandSlice)
		}
	}
}

func list(c []string) {
	var err error
	var files []os.FileInfo
	// check for number of args
	switch len(c) {
	case 1:
		files, err = ioutil.ReadDir("./")
	case 2:
		files, err = ioutil.ReadDir(c[1])
	}
	if err != nil {
		log.Fatalf("Error while listing directory content")
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
