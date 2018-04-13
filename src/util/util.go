package util

import (
	"os"
	"path/filepath"
)

func GetCurrentDir() (string, error) {
	return filepath.Abs("./")
}

func IsFile(name string) (bool, error) {
	info, error := os.Stat(name)
	if error != nil {
		return false, error
	}
	return info.Mode().IsRegular(), nil
}

type Files []os.FileInfo

func (files Files) Len() int {
	return len(files)
}

func (files Files) Less(i, j int) bool {
	return files[i].Name() < files[j].Name()
}

func (files Files) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

func CreateDir(dest string) error {
	_, err := os.Stat(dest)
	if os.IsNotExist(err) {
		err = os.Mkdir(dest, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	//	fmt.Printf("[debug] directory %s\n", dest)

	return nil
}
