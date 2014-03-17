package main

import (
		"io"
		"os"
		"strings"
       )

type rot13Reader struct {
	r io.Reader
}

func rot13Char(ch byte) byte {
	switch {
	case ch >= 'a' && ch <= 'm':
		ch += 13
	case ch >= 'A' && ch <= 'M':
		ch += 13
	case ch >= 'n' && ch <= 'z':
		ch -= 13
	case ch >= 'N' && ch <= 'Z':
		ch -= 13
	}
	return ch
}

func (rot13 rot13Reader) Read(p []byte) (int, error) {
	n, err := rot13.r.Read(p)
	for ii, ch := range p {
		p[ii] = rot13Char(ch)
	}
	return n, err
}

func main() {
	s := strings.NewReader(
		   "Lbh penpxrq gur pbqr!")
	   r := rot13Reader{s}
   io.Copy(os.Stdout, &r)
}
