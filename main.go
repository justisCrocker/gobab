package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		println("Error occured while getting working directory")
		os.Exit(1)
	}

	// Get files in directory
	files, err := ioutil.ReadDir(wd)
	if err != nil {
		log.Fatal(err)
		println("Error occured while getting files from directory")
		os.Exit(1)
	}

	// Get Flag values
	// recursiveFlg := flag.Bool("r", false, "a bool")

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			kebabName := strings.ReplaceAll(strings.ToLower(fileName), " ", "-")
			os.Rename(fmt.Sprintf("%s/%s", wd, fileName), fmt.Sprintf("%s/%s", wd, kebabName))
		}
	}
}
