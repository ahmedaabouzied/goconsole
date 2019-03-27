// Copyright (c) 2019 Ahmed Abouzied <ahmedaabouzied44@gmail.com>.
// This file should not be used for any purposes that contradict the LICENCE.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Get the current working directory.
	CWD, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while getting the current working directory : %s \n ", err)
	}
	// Read text from standard input
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s -- $$ ", CWD)
		reader.Scan()
		text := reader.Text()
		commandSlice := strings.Split(text, " ")
		switch commandSlice[0] {
		case "ls":
			list(commandSlice)
		case "cd":
			changeDir(commandSlice, &CWD)
		case "touch":
			createFile(commandSlice)
		case "del":
			deleteFile(commandSlice, &CWD)
		case "mkdir":
			makeDir(commandSlice, &CWD)
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
	default:
		err = errors.New(" ls : Too many arguments")
	}
	if err != nil {
		errorLog(err)
	}
	// print the files array contents each at a time
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// Change the current working directory.
func changeDir(c []string, wd *string) {
	var err error
	var path string
	switch len(c) {
	// exactly one argument to cd
	case 2:
		path = c[1]
		err = os.Chdir(path)
		if err == nil {
			*wd, err = os.Getwd()
		}
	default:
		err = errors.New("cd : Too many arguments")
	}
	if err != nil {
		errorLog(err)
	}
}

// Create New File
func createFile(l []string) {
	// Create New files
	for i := 1; i < len(l); i++ {
		os.Create(l[i])
	}
}

// Create a new directory with the 777 permissions
func makeDir(l []string, wd *string) {
	for i := 1; i < len(l); i++ {
		var mode os.FileMode
		mode = 777
		path := *wd + "/" + l[i]
		err := os.Mkdir(path, mode)
		if err != nil {
			errorLog(err)
		} else {
			fmt.Printf("Create file %s", l[i])
		}
	}
}

// Delete file in the current directory
func deleteFile(l []string, wd *string) {
	// Delete each file in the params
	for i := 1; i < len(l); i++ {
		filePath := *wd + "/" + l[i]
		err := os.Remove(filePath)
		if err != nil {
			errorLog(err)
		} else {
			fmt.Printf("Deleted File %s \n", l[i])
		}
	}
}

// Prints an error message to the console
func errorLog(err error) {
	fmt.Printf("Error %s \n", err)
}
