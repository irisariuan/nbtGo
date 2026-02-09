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
	"compress/gzip"
	"compress/zlib"
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

func IsNbtData(data []byte) bool {
	if len(data) < 3 {
		return false
	}
	if data[0] == 0x1F && data[1] == 0x8B {
		// Gzip header
		gzipReader, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return false
		}
		defer gzipReader.Close()
		firstByte := make([]byte, 1)
		n, err := gzipReader.Read(firstByte)
		if err != nil || n != 1 {
			return false
		}
		return firstByte[0] == byte(nbt.BTagCompound) || firstByte[0] == byte(nbt.BTagList)
	}
	if data[0] == 0x78 && (data[1] == 0x01 || data[1] == 0x5e || data[1] == 0x9c || data[1] == 0xda) {
		zlibReader, err := zlib.NewReader(bytes.NewReader(data))
		if err != nil {
			return false
		}
		defer zlibReader.Close()
		firstByte := make([]byte, 1)
		n, err := zlibReader.Read(firstByte)
		if err != nil || n != 1 {
			return false
		}
		return firstByte[0] == byte(nbt.BTagCompound) || firstByte[0] == byte(nbt.BTagList)
	}
	return data[0] == byte(nbt.BTagCompound) || data[0] == byte(nbt.BTagList)
}

// FreeMemory frees memory allocated by the library
//
//export FreeMemory
func FreeMemory(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

func main() {}
