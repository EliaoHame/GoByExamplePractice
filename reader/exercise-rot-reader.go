package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13Calc(b[i])
	}
	return n, err
}

func rot13Calc(b byte) byte {
	switch {
	case ('A' <= b && b <= 'M') || ('a' <= b && b <= 'm'):
		b = b + 13
	case ('M' < b && b <= 'Z') || ('m' <= b && b <= 'z'):
		b = b - 13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
