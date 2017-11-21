package main

import "fmt"

// generate binary from bytes
func bytes2Binary(bs []byte) (b string) {
	for _, c := range bs {
		b = fmt.Sprintf("%s%.8b", b, c)
	}
	return
}
