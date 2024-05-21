package specific

import (
	"fmt"
	"runtime"
)

func InstallMac() error {
	if runtime.GOOS != "darwin" {
		return nil
	}

	if runtime.GOARCH == "amd64" {
		fmt.Println("MacOS AMD/Intel specific code")
	}

	if runtime.GOARCH == "arm64" {
		fmt.Println("MacOS ARM specific code | Apple Silicon")
	}

	return nil
}
