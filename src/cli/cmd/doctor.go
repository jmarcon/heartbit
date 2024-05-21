/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"jmarcon/heartbit/hb/specific"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Verify if your environment is setup correctly",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Verify if the dotnet command is available
		_, err := exec.LookPath("dotnet")
		if err != nil {
			fmt.Println("dotnet command not found. Please install .NET Core SDK")
			return
		} else {
			// Verify dotnet version
			out, err := exec.Command("dotnet", "--version").Output()
			if err != nil {
				fmt.Println("Error getting dotnet version")
				return
			}
			// Verify if dotnet version is greater than 8.0
			if string(out) < "8.0" {
				fmt.Println("Please install .NET Core SDK 8.0 or greater")
				//print os version
				return
			}
		}

		//Get current os platform
		fmt.Println("Your OS is: " + runtime.GOOS + " [" + runtime.GOARCH + "]")

		prompt := promptui.Select{
			Label: `Do you want to install all the dependencies?
for your OS ` + runtime.GOOS + " [" + runtime.GOARCH + `]
[Yes/No]`,
			Items: []string{"Yes", "No"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		if result == "Yes" {
			err1 := specific.InstallLinux()
			if err1 != nil {
				fmt.Println("Error installing Linux dependencies")
				return
			}

			err2 := specific.InstallMac()
			if err2 != nil {
				fmt.Println("Error installing MacOS dependencies")
				return
			}

			err3 := specific.InstallWindows()
			if err3 != nil {
				fmt.Println("Error installing Windows dependencies")
				return
			}
		}

		fmt.Println("Your environment is setup correctly")
	},
	DisableAutoGenTag: true,
}

func init() {
	rootCmd.AddCommand(doctorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doctorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doctorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
