package browser

import (
	"os"
	"path/filepath"
)

// home dir path - fallback to current user
var homeDir, _ = os.UserHomeDir()

// GetBrowserProfilePath returns the browser profile path
// If the default homeDir path doesn't exist, it scans all user profiles
func GetBrowserProfilePath(relativePath string) string {
	// Try current user first
	fullPath := filepath.Join(homeDir, relativePath)
	if fileExists(fullPath) {
		return fullPath
	}

	// If not found, scan all user profiles (Windows only)
	profiles, err := GetAllUserProfiles()
	if err != nil {
		return fullPath // Return original if scan fails
	}

	for _, profile := range profiles {
		testPath := filepath.Join(profile, relativePath)
		if fileExists(testPath) {
			return testPath
		}
	}

	return fullPath // Return original as fallback
}

// fileExists checks if a file or directory exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

const (
	chromeName     = "Chrome"
	chromeBetaName = "Chrome Beta"
	chromiumName   = "Chromium"
	edgeName       = "Microsoft Edge"
	braveName      = "Brave"
	operaName      = "Opera"
	operaGXName    = "OperaGX"
	vivaldiName    = "Vivaldi"
	coccocName     = "CocCoc"
	yandexName     = "Yandex"
	firefoxName    = "Firefox"
	speed360Name   = "360speed"
	qqBrowserName  = "QQ"
	dcBrowserName  = "DC"
	sogouName      = "Sogou"
	arcName        = "Arc"
)
