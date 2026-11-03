//go:build linux || darwin || freebsd || dragonfly

package player

import "golang.org/x/sys/unix"

func availableDiskBytes(path string) (uint64, error) {
	var stat unix.Statfs_t
	err := unix.Statfs(path, &stat)
	if err != nil || stat.Bavail <= 0 {
		return 0, err
	}
	return uint64(stat.Bavail) * uint64(stat.Bsize), nil
}
