package system

import (
	"fmt"
	"integriWatch/logger"
	"integriWatch/utils"
	"os"
)

var (
	appDataDir  *string
	initialised bool
)

func Init() error {
	fmt.Println(initialised)

	// Ensure re-initialisation not occurring
	if initialised {
		return fmt.Errorf("init called twice")
	}
	initialised = true

	// Set app data directory and ensure exists
	setAppDataDir()

	// Initialise logger
	logger.Init(*appDataDir)

	// Load config

	return nil
}

func setAppDataDir() {
	dir, _ := os.UserCacheDir()
	appDataDirTMP := dir + string(os.PathSeparator) + "IntegriShield"
	appDataDir = &appDataDirTMP
	utils.EnsureDir(*appDataDir)
}
