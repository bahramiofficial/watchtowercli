package main

import (
	"fmt"
	"os"

	syncprograms "github.com/bahramiofficial/programs/addprogram"
)

func main() {

	if len(os.Args) < 1 {
		fmt.Println("Please provide a command: syncprograms")
		os.Exit(1)
		return
	}

	// Get the command from the first argument
	command := os.Args[1]

	switch command {
	case "syncprograms":
		syncprograms.SyncProgramWithJson()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: syncprograms ")
	}
}
