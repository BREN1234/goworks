package ioworks

import (
	"io"
	"os"
)

func CopyFile(from string, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	dst, err := os.Create(to)
	if err != nil {
		return 0, err
	}
	written, err := io.Copy(dst, src)
	dst.Close()
	src.Close()
	return written, err
}
