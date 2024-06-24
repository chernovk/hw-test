package main

import (
	"log"
	"os"
)

func main() {
	arguments := os.Args
	envDirectory := arguments[1]
	environments, err := ReadDir(envDirectory)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	exitCode := RunCmd(arguments[2:], environments)
	os.Exit(exitCode)
}
