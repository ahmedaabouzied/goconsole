// Copyright (c) 2019 Ahmed Abouzied <ahmedaabouzied44@gmail.com>.
// This file should not be used for any purposes that contradict the LICENCE.

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
		case "exit":
			fmt.Println("Bye")
			os.Exit(0)
		}
	}
}

// list the contents of current directory.
// if a directory is provided it will
// list that directory instead
func list(c []string) {
	var err error
	var files []os.FileInfo
	// check for number of args
	switch len(c) {
	// if no directory is provided
	case 1:
		files, err = ioutil.ReadDir("./")
	// if a directory is provided
	case 2:
		files, err = ioutil.ReadDir(c[1])
	}
	if err != nil {
		log.Fatalf("Error while listing directory content")
	}
	// print the files array contents each at a time
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
