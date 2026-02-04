// Example: Byte conversion utilities
//
// This example demonstrates the byte conversion utilities provided by the lib package.
//
// Usage:
//   go run examples/byte_conversion.go

package main

import (
	"fmt"
	"goNbt/lib"
)

func main() {
	fmt.Println("=== Byte Conversion Examples ===\n")

	// Integer conversions
	fmt.Println("--- Integer Conversions ---")
	
	// Convert int32 to bytes (big-endian)
	value32 := int32(12345)
	bytes32 := lib.Int32ToBytes(value32, true)
	fmt.Printf("Int32 %d to bytes (big-endian): %v\n", value32, bytes32)
	
	// Convert back
	restored32, _ := lib.BytesToInt32(bytes32, true)
	fmt.Printf("Bytes back to int32: %d\n\n", restored32)

	// Convert int64 to bytes (little-endian)
	value64 := int64(9876543210)
	bytes64 := lib.Int64ToBytes(value64, false)
	fmt.Printf("Int64 %d to bytes (little-endian): %v\n", value64, bytes64)
	
	restored64, _ := lib.BytesToInt64(bytes64, false)
	fmt.Printf("Bytes back to int64: %d\n\n", restored64)

	// Float conversions
	fmt.Println("--- Float Conversions ---")
	
	// Convert float32 to bytes
	float32Val := float32(3.14159)
	bytesFloat32 := lib.Float32ToBytes(float32Val, true)
	fmt.Printf("Float32 %f to bytes (big-endian): %v\n", float32Val, bytesFloat32)
	
	restoredFloat32, _ := lib.BytesFloat32(bytesFloat32, true)
	fmt.Printf("Bytes back to float32: %f\n\n", restoredFloat32)

	// Convert float64 to bytes
	float64Val := float64(2.718281828)
	bytesFloat64 := lib.Float64ToBytes(float64Val, true)
	fmt.Printf("Float64 %f to bytes (big-endian): %v\n", float64Val, bytesFloat64)
	
	restoredFloat64, _ := lib.BytesFloat64(bytesFloat64, true)
	fmt.Printf("Bytes back to float64: %f\n\n", restoredFloat64)

	// Endianness comparison
	fmt.Println("--- Endianness Comparison ---")
	testValue := int32(0x12345678)
	
	bigEndianBytes := lib.Int32ToBytes(testValue, true)
	littleEndianBytes := lib.Int32ToBytes(testValue, false)
	
	fmt.Printf("Value: 0x%X\n", testValue)
	fmt.Printf("Big-endian bytes:    %v\n", bigEndianBytes)
	fmt.Printf("Little-endian bytes: %v\n", littleEndianBytes)
}
