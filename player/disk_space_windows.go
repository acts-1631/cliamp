//go:build windows

package player

import (
	"path/filepath"

	"golang.org/x/sys/windows"
)

func availableDiskBytes(path string) (uint64, error) {
	dir, err := windows.UTF16PtrFromString(filepath.Dir(path))
	if err != nil {
		return 0, err
	}
	var available uint64
	err = windows.GetDiskFreeSpaceEx(dir, &available, nil, nil)
	return available, err
}
