package ioworks

import (
	"io"
	"os"
)

func CopyFileDefer(from string, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	dst, err := os.Create(to)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)

}
