// Package lib provides utility functions for byte manipulation and data conversion.
//
// This package contains helper functions used by the nbt package for handling
// binary data conversions and compression. The utilities support both big-endian
// and little-endian byte ordering, making them suitable for working with both
// Java Edition and Bedrock Edition Minecraft data formats.
//
// # Byte Conversion
//
// The package provides bidirectional conversion between Go types and byte slices:
//
// Integer conversions (signed):
//   - Int16ToBytes / BytesToInt16
//   - Int32ToBytes / BytesToInt32
//   - Int64ToBytes / BytesToInt64
//
// Integer conversions (unsigned):
//   - UInt16ToBytes / BytesToUInt16
//   - BytesToUInt32
//   - BytesToUInt64
//
// Floating-point conversions:
//   - Float32ToBytes / BytesFloat32
//   - Float64ToBytes / BytesFloat64
//
// All conversion functions accept a bigEndian parameter to specify byte order:
//   - true: Big-endian (Java Edition, network byte order)
//   - false: Little-endian (Bedrock Edition, most modern CPUs)
//
// # Example Usage
//
// Converting integers to bytes:
//
//	// Convert int32 to big-endian bytes
//	bytes := lib.Int32ToBytes(12345, true)
//
//	// Convert back to int32
//	value, err := lib.BytesToInt32(bytes, true)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Converting floating-point numbers:
//
//	// Convert float64 to little-endian bytes
//	bytes := lib.Float64ToBytes(3.14159, false)
//
//	// Convert back to float64
//	value, err := lib.BytesFloat64(bytes, false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Compression Support
//
// The package provides gzip decompression utilities:
//
//	// Decompress gzip data from a reader
//	file, _ := os.Open("level.dat")
//	data, err := lib.UnzipReader(file)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// This is particularly useful for Minecraft data files which are typically
// stored in gzip-compressed format.
//
// # Error Handling
//
// Conversion functions return an error if the provided byte slice is too short
// for the requested conversion:
//
//	bytes := []byte{0x01} // Only 1 byte
//	_, err := lib.BytesToInt32(bytes, true) // Needs 4 bytes
//	if err != nil {
//	    fmt.Println(err) // "not enough bytes to convert to int32"
//	}
//
// # Internal Utilities
//
// The package also includes internal utilities like bytesReader and Reset for
// advanced stream manipulation, primarily used by the nbt package parsing logic.
//
// # Endianness
//
// Big-endian (most significant byte first):
//   - Used by Java Edition Minecraft
//   - Network byte order (standard for network protocols)
//   - Example: 0x12345678 stored as [0x12, 0x34, 0x56, 0x78]
//
// Little-endian (least significant byte first):
//   - Used by Bedrock Edition Minecraft
//   - Native byte order on x86/x64 processors
//   - Example: 0x12345678 stored as [0x78, 0x56, 0x34, 0x12]
//
// # Thread Safety
//
// All functions in this package are stateless and thread-safe. They can be called
// concurrently from multiple goroutines without synchronization.
package lib
