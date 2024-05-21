/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"
	"log"
	
	"github.com/spf13/cobra/doc"

	"jmarcon/heartbit/hb/cmd"
)

func main() {
	// If the .docs folder exists, do not generate the documentation
	if _, err := os.Stat("./.docs/"); err == nil {
		cmd.Execute()
		return
	}

	if err := os.Mkdir("./.docs/", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	err := doc.GenMarkdownTree(*cmd.GetCommand(), "./.docs/")
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
