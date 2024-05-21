package specific

import (
	"fmt"
	"runtime"
)

func InstallLinux() error {
	if runtime.GOOS != "linux" {
		return nil
	}

	// Add repos for Ubuntu based systems
	ubuntuRepositories := []string{
		"ppa:boltgolt/howdy",
		"ppa:deadsnakes/ppa",
		"ppa:papirus/papirus",
		"ppa:neovim-ppa/stable",
		"deb [arch=amd64] https://packages.microsoft.com/repos/edge stable main",
	}

	for _, repo := range ubuntuRepositories {
		fmt.Println("Adding repository: ", repo)
	}

	// Update and upgrade
	fmt.Println("Updating and upgrading system")

	linuxPackages := []string{
		"build-essential",
		"bzip2",
		"ca-certificates",
		"cmake",
		"cpanminus",
		"curl",
		"docker-buildx-plugin",
		"docker-ce",
		"docker-ce-cli",
		"docker-compose-plugin",
		"eza",
		"fd-find",
		"ffmpeg",
		"gcc",
		"git",
		"gnome-tweaks",
		"howdy",
		"kubectl",
		"make",
		"mpv",
		"onedriver",
		"openssl",
		"papirus-icon-theme",
		"perl",
		"php-cli",
		"python-software-properties",
		"python3-dev",
		"python3-neovim",
		"python3-pip",
		"python3-setuptools",
		"python3-wheel",
		"rbenv",
		"solaar",
		"software-properties-common",
		"warp-terminal",
		"wget",
		"x11-utils",
		"xautomation",
		"xorg-dev",
		"zsh",
	}

	for _, pkg := range linuxPackages {
		fmt.Println("Installing package: ", pkg)
	}

	if runtime.GOARCH == "amd64" || runtime.GOARCH == "386" {
		fmt.Println("Linux AMD/Intel specific code")
	}

	if runtime.GOARCH == "arm" || runtime.GOARCH == "arm64" {
		fmt.Println("Linux ARM specific code | Container, Raspberry PI or Graviton")
	}

	return nil
}
