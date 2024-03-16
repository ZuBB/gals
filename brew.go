package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const brewPath = "/opt/homebrew/bin/brew"
const brewCommandEnvVar1 = "HOMEBREW_NO_AUTO_UPDATE=1"
const filename = "Brewfile"

func SaveCurrentBrewApps(basePath string) {
	outputPath := filepath.Join(basePath, ListTypeCurrent, filename)
	cmd := exec.Command("bash", brewPath, "bundle", "dump", "-qf", "--brews", "--cask", "--tap", "--mas", "--file", outputPath)
	cmd.Env = append(os.Environ(), brewCommandEnvVar1)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Printf("%#v", err)
		fmt.Printf("%v", err)
	}
}
