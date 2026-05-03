package internal

import (
	"os"
	"strings"
)

func GetPathSize(path string, withHidden bool) (int64, error) {
	var s = int64(0)
	fi, err := os.Lstat(path)
	if err != nil {
		return s, err
	}

	if fi.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return s, err
		}

		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil && info.IsDir() {
				continue
			}

			if !withHidden && strings.HasPrefix(info.Name(), ".") {
				continue
			}

			s += info.Size()
		}
	} else {
		s = fi.Size()
	}

	return s, nil
}
