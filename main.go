package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/strebul/parallelCmd/cmd"
	"time"
)

func main() {
	var command string
	var argsRaw string
	var args []string
	var directory string
	var subDirsRaw string
	var subDirs []string

	command = os.Args[1]
	argsRaw = os.Args[2]
	directory = os.Args[3]
	subDirsRaw = os.Args[4]

	// Checking root directory
	if _, err := os.Stat(directory); err != nil {
		panic("Undefined directory " + directory)
	}

	subDirs = strings.Fields(subDirsRaw)
	args = strings.Fields(argsRaw)

	// Checking sub directories
	for _, sub := range subDirs {
		if _, err := os.Stat(directory + "/" + sub); err != nil {
			panic("Undefined sub directory " + sub)
		}
	}

	chanel := make(chan cmd.Result)

	for _, sub := range subDirs {
		go cmd.Exec(chanel, directory, sub, command, args)
	}

	for {
		result := <- chanel
		fmt.Println(result)
		time.Sleep(time.Second * 1)
	}

	close(chanel)

	fmt.Println(directory)
	fmt.Println(subDirs)
}

