package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const rootAppsDir = "/Applications/"
const macosSysApps = "macos-system-apps"
const macosUserApps = "macos-user-apps"

func getApps(targetDirPath string) (string, error) {
	dirEntries, err := os.ReadDir(targetDirPath)
	if err != nil {
		return "", err
	}

	var entryNames []string
	for _, entry := range dirEntries {
		appName := entry.Name()

		if !strings.HasPrefix(appName, ".") {
			entryNames = append(entryNames, appName)
		}
	}

	return strings.Join(entryNames, "\n"), nil
}

func saveApps(appsPath, baseTargetDir, listType, appsType string) {
	appsList, err := getApps(appsPath)

	if err != nil {
		log.Printf("%v", err)
		// return err
	}

	targetDir := filepath.Join(baseTargetDir, listType)
	filePath := filepath.Join(targetDir, appsType)

	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		log.Printf("%v", err)
		// return err
	}

	err = SaveFileContent(appsList, filePath)
	if err != nil {
		log.Printf("%v", err)
		// return err
	}

	// return nil
}

func saveCurrentAppsList(targetDir string) {
	// system apps @ current
	saveApps(rootAppsDir, targetDir, ListTypeCurrent, macosSysApps)

	// user apps @ current
	appsDir := filepath.Join(GetHomeDir(), rootAppsDir)
	saveApps(appsDir, targetDir, ListTypeCurrent, macosUserApps)
}

func saveAllTimeAppsList(/*targetDir string*/) {
	// TODO
	log.Println("[macos-apps.go::saveAllTimeAppsList] not implemented yet")
}

func doesEnvMatch() bool {
	return runtime.GOOS == "darwin"
}

func SaveMacOsApps(targetDir string) {
	if doesEnvMatch() {
		log.Printf("SaveCurrentMacOsApps: current platform %s is not suported\n", runtime.GOOS)
		return
	}

	saveCurrentAppsList(targetDir)
	saveAllTimeAppsList(/*targetDir*/)
}

// if error := saveUserApps(targetDir); error != nil {
// 	log.Printf("%v", error)
// 	return error
// }
