//go:build windows

package browser

import (
	"os"
	"path/filepath"
)

// GetAllUserProfiles returns all user profile directories on Windows
// Scans C:\Users\ to find all user folders (handles profiles like "User.COMPUTERNAME")
func GetAllUserProfiles() ([]string, error) {
	systemDrive := os.Getenv("SystemDrive")
	if systemDrive == "" {
		systemDrive = "C:"
	}
	usersDir := systemDrive + "\\Users"
	
	entries, err := os.ReadDir(usersDir)
	if err != nil {
		return nil, err
	}

	var profiles []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		// Skip system folders
		name := entry.Name()
		if name == "Public" || name == "Default" || name == "All Users" || name == "Default User" {
			continue
		}

		profilePath := filepath.Join(usersDir, name)
		
		// Verify it's a real user profile (has AppData folder)
		appDataPath := filepath.Join(profilePath, "AppData")
		if _, err := os.Stat(appDataPath); err == nil {
			profiles = append(profiles, profilePath)
		}
	}

	return profiles, nil
}
