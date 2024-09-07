package utils

import "os"

func EnsureDir(dir string) {
	_ = os.MkdirAll(dir, os.ModeDir)
}
