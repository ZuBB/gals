package main

import (
	"log"
	"os"
)

func GetHomeDir() string {
	homedir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Failed to get user's home dir")
		os.Exit(1)
	}

	return homedir
}
