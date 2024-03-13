package main

import (
	// "fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// const listTypeCurrent = "current"
// const listTypeAllTime = "all-time"

func getArchBasedPath() string {
	return filepath.Join("arches", runtime.GOARCH)
}

func getOsBasedPath() string {
	return filepath.Join("oses", runtime.GOOS)
}

func getHostBasedPath() (string, error) {
	fqdn, err := os.Hostname()

	if err != nil {
		return "", err
	}

	hostname := strings.Split(fqdn, ".")[0]
	return filepath.Join("hosts", hostname), nil
}

func getBaseDir() string {
	const shareDir = ".local/share/" // TODO
	const appDir = "gals/"
	return filepath.Join(GetHomeDir(), shareDir, appDir)
}

func main() {
	baseDir := getBaseDir()
	// log.Println(baseDir)
	hostnamePathSegment, err := getHostBasedPath()

	if err != nil {
		log.Fatal("Failed to get user's home dir")
		os.Exit(1)
	}

	dynamicPathSegment := filepath.Join(getArchBasedPath(), getOsBasedPath(), hostnamePathSegment)
	targetPath := filepath.Join(baseDir, dynamicPathSegment)
	// fmt.Println(targetPath)

	SaveCurrentSystemApps(targetPath)
	SaveCurrentUserApps(targetPath)
	SaveCurrentBrewApps(targetPath)
}
