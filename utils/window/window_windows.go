//go:build windows

package window

import (
	"syscall"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetConsoleWindow = kernel32.NewProc("GetConsoleWindow")
	procShowWindow       = user32.NewProc("ShowWindow")
)

const (
	SW_HIDE = 0
)

// Hide hides the console window (Windows only)
func Hide() error {
	hwnd, _, _ := procGetConsoleWindow.Call()
	if hwnd == 0 {
		return nil // No console window
	}
	_, _, err := procShowWindow.Call(hwnd, SW_HIDE)
	if err != nil && err.Error() != "The operation completed successfully." {
		return err
	}
	return nil
}
