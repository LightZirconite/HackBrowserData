//go:build !windows

package window

// Hide is a no-op on non-Windows platforms
func Hide() error {
	return nil
}
