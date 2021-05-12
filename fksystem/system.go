package fksystem

import (
	"os"
	"path/filepath"
	"runtime"
)

// DirSeparator returns directory separator symbol
func DirSeparator() string {

	dirSeparator := "/"
	if runtime.GOOS == "windows" {
		dirSeparator = "\\"
	}

	return dirSeparator

}

// CurrentDir returns current working directory
func CurrentDir() string {
	dir, error := os.Getwd()
	if error != nil {
		panic(error.Error())
	}
	return dir
}

// CountCPUs return number of available CPU cores
func CountCPUs() int {
	return runtime.NumCPU()
}

// EnsureDir ensures that the directory exists
func EnsureDir(path string) error {
	parent := filepath.Dir(path)
	if parent == path {
		return nil
	} else {
		err := EnsureDir(parent)
		if err != nil {
			return err
		}
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(path, os.ModePerm)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
