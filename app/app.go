package app

import (
	"os"

	stdlog "log"
	"path/filepath"
)

//const appName = "HowellApp"

var (
	// DefaultNodeHome defines default home directories for howelld
	DefaultNodeHome string

)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		stdlog.Println("Failed to get home dir %2", err)
	}
	
	DefaultNodeHome = filepath.Join(userHomeDir, ".howell")
}
