package cudacomp

// #cgo LDFLAGS: -L./ -llz4decom
//#include "lz4_cpu.h"
/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// lz4解压缩
func Cuda_Lz4_compress(val []byte) (value []byte) {

	var size C.size_t
	cByteStream := C.go_data_comp((*C.char)(unsafe.Pointer(&val[0])), C.int(len(val)), &size)
	// Convert the C array to a Go slice
	buf := (*[1 << 30]byte)(unsafe.Pointer(cByteStream))[:size:size]

	return buf
}

func Cuda_Lz4_decompress(data []byte) (value []byte) {
	//---------------------------解压
	var size C.size_t
	cByteStream := C.go_data_re((*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)), &size)
	// Convert the C array to a Go slice
	out := (*[1 << 30]byte)(unsafe.Pointer(cByteStream))[:size:size]

	return out
}
