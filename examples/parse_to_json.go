// Example: Parse an NBT file and convert it to JSON
//
// This example demonstrates how to read a gzip-compressed NBT file,
// parse it, and convert it to JSON format for easy viewing.
//
// Usage:
//   go run examples/parse_to_json.go level.dat

package main

import (
	"encoding/json"
	"fmt"
	"goNbt/lib"
	"goNbt/lib/nbt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run parse_to_json.go <nbt-file>")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Open the NBT file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decompress the gzip data
	data, err := lib.UnzipReader(file)
	if err != nil {
		fmt.Printf("Error decompressing file: %v\n", err)
		os.Exit(1)
	}

	// Parse the NBT data (false = Java Edition / big-endian)
	tag, err := nbt.ParseNBT(data, false)
	if err != nil {
		fmt.Printf("Error parsing NBT: %v\n", err)
		os.Exit(1)
	}

	// Convert to formatted JSON
	jsonData, err := json.MarshalIndent(tag, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
