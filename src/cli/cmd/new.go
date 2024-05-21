/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var ProjectName string
var OutputDirectory string
var Path string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new application using ❤️️bit standard",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// If outputdirectory is empty, not provided or equaul current folder, set the output directory to the project name
		if OutputDirectory == "./" || OutputDirectory == "" || OutputDirectory == "." {
			OutputDirectory = "./src"
		}

		// Set the Output path from variables
		OutputDirectory = filepath.Clean(OutputDirectory)
		Path = filepath.Clean(filepath.Join(OutputDirectory, ProjectName))

		// Create the project directory
		if _, err := os.Stat(Path); err == nil {
			fmt.Println("Project already exists")
			return
		}

		// Create the dotnet project using exec
		out, err := exec.Command("dotnet", "new", "webapi", "-o", filepath.Clean(OutputDirectory), "-n", ProjectName).Output()
		if err != nil {
			fmt.Println("Error creating project")
			return
		}

		fmt.Println(string(out))
	},
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().StringVarP(&ProjectName, "project-name", "p", "Scaffold", "Project Name")
	newCmd.PersistentFlags().StringVarP(&OutputDirectory, "output", "o", "./src", "Output directory")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
