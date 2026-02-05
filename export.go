//go:build cshared

package main

/*
 #include <stdio.h>
 #include <errno.h>
 #include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"encoding/json"
	"goNbt/lib"
	"goNbt/lib/nbt"
	"unsafe"
)

// ParseNBT parses NBT binary data and returns JSON string
//
//export ParseNBT
func ParseNBT(data *C.char, length C.int, isBedrock C.int) *C.char {
	goData := C.GoBytes(unsafe.Pointer(data), length)

	// Unzip if needed
	unzippedData, err := lib.UnzipReader(bytes.NewReader(goData))
	if err != nil {
		// If unzip fails, try to parse as-is
		unzippedData = goData
	}

	tag, err := nbt.ParseNBT(unzippedData, isBedrock != 0)
	if err != nil {
		return C.CString("ERROR: " + err.Error())
	}

	jsonBytes, err := json.MarshalIndent(tag, "", "  ")
	if err != nil {
		return C.CString("ERROR: " + err.Error())
	}

	return C.CString(string(jsonBytes))
}

// SerializeNBT serializes JSON string to NBT binary data
//
//export SerializeNBT
func SerializeNBT(jsonData *C.char, compress *C.char, outLength *C.int) *C.char {
	goJSON := C.GoString(jsonData)
	compressType := C.GoString(compress)

	var tag nbt.TagCompound
	err := json.Unmarshal([]byte(goJSON), &tag)
	if err != nil {
		*outLength = 0
		return C.CString("ERROR: " + err.Error())
	}

	serializedBytes, err := nbt.SerializeTag(&tag, false)
	if err != nil {
		*outLength = 0
		return C.CString("ERROR: " + err.Error())
	}

	// Apply compression if requested
	switch compressType {
	case "gzip":
		serializedBytes, err = lib.ZipToGzip(serializedBytes)
		if err != nil {
			*outLength = 0
			return C.CString("ERROR: " + err.Error())
		}
	case "zlib":
		serializedBytes, err = lib.ZipToZlib(serializedBytes)
		if err != nil {
			*outLength = 0
			return C.CString("ERROR: " + err.Error())
		}
	}

	*outLength = C.int(len(serializedBytes))

	// Allocate C memory for binary data
	cBytes := C.CBytes(serializedBytes)
	return (*C.char)(cBytes)
}

// FreeMemory frees memory allocated by the library
//
//export FreeMemory
func FreeMemory(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

func main() {}
