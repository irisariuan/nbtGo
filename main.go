// nbtGo is a command-line tool for converting gzip-compressed NBT data to JSON format.
//
// Usage:
//
//	cat level.dat | nbtGo
//	nbtGo < player.dat
//
// The tool reads gzip-compressed NBT data from stdin and outputs formatted JSON to stdout.
// This is useful for inspecting Minecraft data files like level.dat, player data, and more.
//
// For Java Edition (big-endian) files - most common Minecraft data files
// For Bedrock Edition files, you would need to modify the code to use isBedrock=true
package main

import (
	"encoding/json"
	"fmt"
	"goNbt/lib"
	"goNbt/lib/nbt"
	"os"
)

func main() {
	// Read and decompress gzip data from stdin
	allBytes, err := lib.UnzipReader(os.Stdin)

	if err != nil {
		panic(err)
	}

	// Parse the NBT data (false = Java Edition / big-endian)
	tag, err := nbt.ParseNBT(allBytes, false)
	if err != nil {
		panic(err)
	}

	// Convert to formatted JSON and print
	jsonTag, err := json.MarshalIndent(tag, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonTag))
}
