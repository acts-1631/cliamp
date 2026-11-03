//go:build !windows && !linux && !darwin && !freebsd && !openbsd && !dragonfly

package player

// Unsupported platforms still enforce the process-wide absolute limit.
func availableDiskBytes(string) (uint64, error) {
	return ^uint64(0), nil
}
