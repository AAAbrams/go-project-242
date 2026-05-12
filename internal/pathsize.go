package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := resolvePathSize(path, recursive, all)
	if err != nil {
		return "", err
	}
	return outputFmt(size, path, human), nil
}

func resolvePathSize(path string, recursive, all bool) (int64, error) {
	var size = int64(0)
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return size, err
	}

	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return size, err
		}

		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "warning: cannot get info for %s: %v\n", entry.Name(), err)
				continue
			}

			if !all && strings.HasPrefix(info.Name(), ".") {
				continue
			}

			if info.IsDir() {
				if recursive {
					internalPath := filepath.Join(path, info.Name())
					internalSize, internalErr := resolvePathSize(internalPath, recursive, all)
					if internalErr != nil {
						fmt.Fprintf(os.Stderr, "warning: cannot calculate size for %s: %v\n", internalPath, internalErr)
						continue
					}
					size += internalSize
				}

			} else {
				size += info.Size()
			}
		}
	} else {
		size = fileInfo.Size()
	}

	return size, nil
}

const (
	kb = 1024
	mb = kb * kb
	gb = mb * kb
	tb = gb * kb
	pb = tb * kb
	eb = pb * kb
)

func outputFmt(size int64, path string, human bool) string {
	sl := "B"
	if !human {
		return fmt.Sprintf("%d"+sl+"\t%s", size, path)
	}

	fs := float64(size)
	switch {
	case fs >= kb && fs < mb:
		sl = "KB"
		fs /= kb
	case fs >= mb && fs < gb:
		sl = "MB"
		fs /= mb
	case fs >= gb && fs < tb:
		sl = "GB"
		fs /= gb
	case fs >= tb && fs < pb:
		sl = "TB"
		fs /= tb
	case fs >= pb && fs < eb:
		sl = "PB"
		fs /= pb
	case fs >= eb:
		sl = "EB"
		fs /= eb
	}

	format := "%.1f" + sl + "\t%s"
	return fmt.Sprintf(format, fs, path)
}
