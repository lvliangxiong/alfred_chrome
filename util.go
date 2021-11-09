package main

import (
	"io"
	"os"
)

func copyFile(sourceFile string, destFile string) error {
	buf := make([]byte, 1024)

	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}

	defer func(source *os.File) {
		err := source.Close()
		if err != nil {

		}
	}(source)

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}

	defer func(dest *os.File) {
		err := dest.Close()
		if err != nil {

		}
	}(dest)

	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		_, err = dest.Write(buf[:n])
		if err != nil {
			return err
		}
	}

	return nil
}
