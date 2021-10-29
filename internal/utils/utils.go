package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

// Creates a small startup delay
// For example helpful for waiting for DB container
// (e.g. in the case of Docker Compose)
func StartupDelay(duration time.Duration) {
	message := fmt.Sprintf("Starting in %d seconds...", int(duration.Seconds()))
	color.Set(color.FgYellow, color.Bold)
	log.Println(message)
	color.Unset()
	time.Sleep(duration)
}

// Prints some ASCII art
// TODO: Print version as well
func PrintStartupBanner() {
	log.Println(`
 ____  _                _   _             _             
/ ___|| |__   ___  _ __| |_(_)_ __   __ _| |_ ___  _ __ 
\___ \| '_ \ / _ \| '__| __| | '_ \ / _  | __/ _ \| '__|
 ___) | | | | (_) | |  | |_| | | | | (_| | || (_) | |   
|____/|_| |_|\___/|_|   \__|_|_| |_|\__,_|\__\___/|_|   
	`)
}
