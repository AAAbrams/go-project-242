package code

import "code/internal"

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := internal.ResolvePathSize(path, recursive, all)
	if err != nil {
		return "", err
	}
	return internal.OutputFmt(size, human), nil
}
