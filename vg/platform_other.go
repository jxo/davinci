// +build !arm !arm64
// +build !js

package vg

import (
	"encoding/binary"
	"log"
	"math"
	"unsafe"
)

type Float float32

var shaderHeader = `
#define DAVINCI_GL2 1
#define UNIFORMARRAY_SIZE 11
`

func prepareTextureBuffer(data []byte, w, h, bpp int) []byte {
	return data
}

func castFloat32ToByte(vertexes []float32) []byte {
	// Convert []float32 list to []byte without copy
	var b []byte
	if len(vertexes) > 65536 {
		b = make([]byte, len(vertexes)*4)
		for i, v := range vertexes {
			binary.LittleEndian.PutUint32(b[4*i:], math.Float32bits(v))
		}
	} else {
		b = (*(*[1 << 20]byte)(unsafe.Pointer(&vertexes[0])))[:len(vertexes)*4]
	}
	return b
}

func dumpLog(fmt string, values ...interface{}) {
	log.Printf(fmt, values...)
}