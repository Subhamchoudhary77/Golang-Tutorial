// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
		w.size += int64(binary.Size(v))
	}
}

func main() {
	fmt.Println("test")
}
