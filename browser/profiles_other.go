//go:build !windows

package browser

import "os"

// GetAllUserProfiles returns only the current user profile on non-Windows systems
func GetAllUserProfiles() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return []string{home}, nil
}
