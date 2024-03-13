package main

import (
	"os"
	"path/filepath"
	"strings"
)

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

func saveApps(appsPath, baseTargetDir, listType, appsType string) error {
	appsList, err := getApps(appsPath)

	if err != nil {
		return err
	}

	targetDir := filepath.Join(baseTargetDir, listType)
	filePath := filepath.Join(targetDir, appsType)

	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		return err
	}

	err = SaveFileContent(appsList, filePath)
	if err != nil {
		return err
	}

	return nil
}

func SaveCurrentSystemApps(targetDir string) error {
	return saveApps("/Applications/", targetDir, ListTypeCurrent, macosSysApps)
}

func SaveCurrentUserApps(targetDir string) error {
	appsDir := filepath.Join(GetHomeDir(), "Applications/")
	return saveApps(appsDir, targetDir, ListTypeCurrent, macosUserApps)
}
