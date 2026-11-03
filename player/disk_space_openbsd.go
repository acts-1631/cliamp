//go:build openbsd

package player

import "golang.org/x/sys/unix"

func availableDiskBytes(path string) (uint64, error) {
	var stat unix.Statfs_t
	err := unix.Statfs(path, &stat)
	if err != nil || stat.F_bavail <= 0 {
		return 0, err
	}
	return uint64(stat.F_bavail) * uint64(stat.F_bsize), nil
}
