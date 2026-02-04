// Example: Create a custom NBT structure and serialize it
//
// This example shows how to create NBT structures programmatically
// and serialize them to binary format.
//
// Usage:
//   go run examples/create_nbt.go

package main

import (
	"encoding/json"
	"fmt"
	"goNbt/lib/nbt"
	"os"
)

func main() {
	// Create a sample player data structure
	playerData := &nbt.TagCompound{
		Value: []nbt.NBTTag{
			&nbt.TagString{
				Value: "Steve",
			},
			&nbt.TagInt{
				Value: 20,
			},
			&nbt.TagFloat{
				Value: 100.0,
			},
			&nbt.TagList{
				ElementType: nbt.BTagInt,
				Value: []nbt.NBTTag{
					&nbt.TagInt{Value: 1},
					&nbt.TagInt{Value: 2},
					&nbt.TagInt{Value: 64},
				},
			},
			&nbt.TagCompound{
				Value: []nbt.NBTTag{
					&nbt.TagDouble{Value: 100.5},
					&nbt.TagDouble{Value: 64.0},
					&nbt.TagDouble{Value: 100.5},
					&nbt.TagEnd{},
				},
			},
			&nbt.TagEnd{},
		},
	}

	// Set the base tag information
	// Note: In a real use case, you would properly initialize baseTag fields
	// This is simplified for demonstration

	// Serialize to NBT format
	bytes, err := nbt.SerializeTag(playerData, false)
	if err != nil {
		fmt.Printf("Error serializing NBT: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Serialized NBT data: %d bytes\n", len(bytes))

	// Also show as JSON for verification
	jsonData, err := json.MarshalIndent(playerData, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nJSON representation:")
	fmt.Println(string(jsonData))

	// Write to file
	err = os.WriteFile("example_output.nbt", bytes, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nNBT data written to example_output.nbt")
}
