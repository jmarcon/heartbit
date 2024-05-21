package specific

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func InstallWindows() error {
	if runtime.GOOS != "windows" {
		return nil
	}

	// Verify if powershell is available
	_, err := exec.LookPath("powershell")
	if err == nil {
		cmd := exec.Command("powershell", "-nologo", "-noprofile")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		// Installing Scoop
		go func() {
			defer stdin.Close()
			fmt.Fprintln(stdin, "Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser")
			fmt.Fprintln(stdin, "Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression")
		}()

		// Installing Wsl
		go func() {
			defer stdin.Close()
			fmt.Fprintln(stdin, "wsl --install")
		}()

		// Installing Packages
		go func() {
			defer stdin.Close()

			// Add buckets
			buckets := []string{
				"extras",
				"nerd-fonts",
				"versions",
			}

			for _, bucket := range buckets {
				fmt.Fprintf(stdin, "scoop bucket add %s\n", bucket)
			}

			packages := []string{
				"7zip",
				"aws",
				"base64",
				"bat",
				"cacert",
				"clink",
				"clipboard",
				"cloudflared",
				"ctop",
				"curl",
				"curlie",
				"dark",
				"duf",
				"fzf",
				"gawk",
				"gcc",
				"gh",
				"git",
				"go",
				"graphviz",
				"helm",
				"innounp",
				"jid",
				"jq",
				"k6",
				"k9s",
				"kubectl",
				"krew",
				"kubectx",
				"kubens",
				"less",
				"neovim",
				"nmap",
				"nodejs",
				"nvm",
				"oh-my-posh",
				"openssl",
				"pwsh",
				"python",
				"ripgrep",
				"sed",
				"sudo",
				"touch",
				"wget",
				"which",
				"wttop",
				"yq",
			}

			for _, pkg := range packages {
				fmt.Fprintf(stdin, "scoop install %s\n", pkg)
			}

			// extras
			extras := []string{
				"eartrumpet",
				"insomnia",
				"kubebox",
				"lazy-posh-git",
				"lazygit",
				"plantuml",
				"posh-docker",
				"posh-git",
				"powertoys",
				"psfzf",
				"psreadline",
				"scoop-completion",
				"sfsu",
				"Terminal-Icons",
				"vscode",
			}

			for _, extra := range extras {
				fmt.Fprintf(stdin, "scoop install extras/%s\n", extra)
			}
		}()

		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}

	if runtime.GOARCH == "arm" || runtime.GOARCH == "arm64" {
		fmt.Println("Windows ARM specific code! WOW!!!")
	}

	if runtime.GOARCH == "amd64" || runtime.GOARCH == "386" {
		fmt.Println("Windows AMD/Intel specific code")
	}

	return nil
}
