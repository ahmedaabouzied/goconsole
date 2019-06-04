package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "ls",
	Short: "List the contents of the a directory",
	Long:  "Lists the contents of the specified directory. If no directory is specified it lists the contents of the current direcotry.",
	Run: func(cmd *cobra.Command, args []string) {
		list(args)
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
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
	case 0:
		files, err = ioutil.ReadDir("./")
	// if a directory is provided
	case 1:
		files, err = ioutil.ReadDir(c[0])
	default:
		err = errors.New(" ls : Too many arguments")
	}
	if err != nil {
		log.Fatal(err)
	}
	// print the files array contents each at a time
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
