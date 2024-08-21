package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gleich/lumber/v3"
)

func main() {
	demos := []func(){
		func() {
			lumber.Done("Loaded up the program!")
			time.Sleep(2 * time.Second)
			lumber.Done("Waited 2 seconds")
		},
		func() {
			lumber.Info("Getting the current year")
			now := time.Now()
			lumber.Info("Current year is", now.Year())
		},
		func() {
			homeDir, _ := os.UserHomeDir()
			lumber.Debug("User's home dir is", homeDir)
		},
		func() {
			now := time.Now()
			if now.Year() != 2004 {
				lumber.Warning("Current year isn't 2004")
			}
		},
		func() {
			fname := "invisible-file.txt"
			_, err := os.ReadFile(fname)
			if err != nil {
				lumber.Error(err, "Failed to read from", fname)
			}
		},
		func() {
			lumber.FatalMsg("Ahhh stuff broke")
		},
	}

	for _, demo := range demos {
		fmt.Println()
		demo()
		fmt.Println()
	}
}
