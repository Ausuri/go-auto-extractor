// File: filewatch/filewatch_tools.go
// Package filewatch provides common useful functions that can be used in interface implementations of filewatch.

package filewatch

import (
	"io/fs"
	"os"
	"path/filepath"
)

func getSubDirectories(path string) (directoryPathList []string, err error) {

	var subdirectories []string

	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != "." && path != "/" {
			subdirectories = append(subdirectories, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return subdirectories, nil
}

func getEventType(path string) (itemType EventType, err error) {

	info, err := os.Stat(path)
	if err != nil {
		return UnknownEventType, err
	}

	if info.IsDir() {
		itemType = CreateDirectory
	} else {
		itemType = CreateFile
	}

	return itemType, nil

}
