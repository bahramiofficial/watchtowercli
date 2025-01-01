package syncprograms

import (
	"log"
	"path/filepath"
	"runtime"

	watchsyncprograms "github.com/bahramiofficial/watchtower/src/utilities/programs"
)

func SyncProgramWithJson() {
	// Get the directory of the script
	_, scriptPath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Failed to get the script directory")
	}

	// init path json files
	dirPath := filepath.Join(filepath.Dir(scriptPath), "json/")
	watchsyncprograms.SyncProgramToDb(dirPath)
}
