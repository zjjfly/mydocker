package util

import "os"

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func IsDir(f string) bool {
	fi, err := os.Stat(f)
	return err == nil && fi.IsDir()
}

func IsFile(f string) bool {
	fi, err := os.Stat(f)
	return err == nil && !fi.IsDir()
}
