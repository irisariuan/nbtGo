// Package nbt provides parsing, serialization, and manipulation of NBT (Named Binary Tag) data structures.
//
// NBT is a binary data format used by Minecraft for storing structured game data such as world information,
// player inventories, block entities, and more. This package supports the complete NBT specification including
// all standard tag types.
//
// # Overview
//
// NBT data consists of hierarchical tags, each with a type, name, and payload. Tags can be primitive types
// (byte, short, int, long, float, double, string) or container types (compound, list, arrays).
//
// # Supported Tag Types
//
// The package supports all standard NBT tag types:
//   - TAG_End (0): Marks the end of a compound tag
//   - TAG_Byte (1): Single signed byte
//   - TAG_Short (2): Signed 16-bit integer (big-endian)
//   - TAG_Int (3): Signed 32-bit integer (big-endian)
//   - TAG_Long (4): Signed 64-bit integer (big-endian)
//   - TAG_Float (5): 32-bit IEEE 754 floating point
//   - TAG_Double (6): 64-bit IEEE 754 floating point
//   - TAG_Byte_Array (7): Array of signed bytes
//   - TAG_String (8): UTF-8 encoded string
//   - TAG_List (9): List of unnamed tags (all same type)
//   - TAG_Compound (10): Collection of named tags
//   - TAG_Int_Array (11): Array of signed 32-bit integers
//   - TAG_Long_Array (12): Array of signed 64-bit integers
//
// # Parsing NBT Data
//
// To parse NBT binary data, use the ParseNBT function:
//
//	data, _ := os.ReadFile("level.dat")
//	tag, err := nbt.ParseNBT(data, false) // false for Java Edition (big-endian)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// The second parameter indicates whether the data is from Bedrock Edition (little-endian).
// Set it to false for Java Edition data (big-endian), which is the most common format.
//
// # Serializing NBT Data
//
// To serialize Go structures back to NBT format:
//
//	tag := &TagString{
//	    baseTag: baseTag{tagType: BTagString, name: "hello"},
//	    Value:   "world",
//	}
//	bytes, err := SerializeTag(tag, false)
//
// The second parameter (skipHeader) should be false for top-level tags and true only
// when serializing elements inside a list.
//
// # JSON Conversion
//
// All tag types implement json.Marshaler and json.Unmarshaler for easy conversion:
//
//	jsonData, err := json.Marshal(tag)
//	// or
//	var tag TagString
//	err := json.Unmarshal(jsonData, &tag)
//
// # Working with Compound Tags
//
// Compound tags are the most common container type, similar to a map or dictionary:
//
//	compound := &TagCompound{
//	    baseTag: baseTag{tagType: BTagCompound, name: "root"},
//	    Value: []NBTTag{
//	        &TagString{...},
//	        &TagInt{...},
//	        &TagEnd{}, // Compound tags must end with TAG_End
//	    },
//	}
//
// # Working with List Tags
//
// List tags contain multiple unnamed tags of the same type:
//
//	list := &TagList{
//	    baseTag:     baseTag{tagType: BTagList, name: "items"},
//	    ElementType: BTagInt,
//	    Value: []NBTTag{
//	        &TagInt{Value: 1},
//	        &TagInt{Value: 2},
//	        &TagInt{Value: 3},
//	    },
//	}
//
// # Error Handling
//
// Parsing functions return TagParseError which includes information about whether
// the error is fatal:
//
//	tag, err := ParseNBT(data, false)
//	if err != nil {
//	    if parseErr, ok := err.(TagParseError); ok {
//	        if parseErr.isFatal() {
//	            // Handle fatal error
//	        }
//	    }
//	}
//
// # Endianness
//
// The package supports both big-endian (Java Edition) and little-endian (Bedrock Edition) formats.
// The endianness is specified during parsing and serialization. Java Edition uses big-endian
// by default, which is the standard NBT format.
//
// # Example: Reading a Minecraft Level File
//
//	import (
//	    "goNbt/lib"
//	    "goNbt/lib/nbt"
//	    "os"
//	)
//
//	func readLevel() {
//	    file, _ := os.Open("level.dat")
//	    defer file.Close()
//
//	    // Most Minecraft files are gzip-compressed
//	    data, _ := lib.UnzipReader(file)
//
//	    // Parse the NBT data (false = Java Edition)
//	    tag, _ := nbt.ParseNBT(data, false)
//
//	    // Work with the parsed tag
//	    nbt.PrintTag(tag)
//	}
//
// # Binary Format Details
//
// Each NBT tag in binary format consists of:
//  1. Tag type byte (1 byte)
//  2. Name length (2 bytes, unsigned, big-endian)
//  3. Name string (UTF-8 encoded)
//  4. Payload (varies by tag type)
//
// TAG_End is special: it only consists of the type byte (0x00) with no name or payload.
//
// # Thread Safety
//
// The parsing and serialization functions are stateless and thread-safe. However,
// the tag structures themselves are not synchronized and should not be modified
// concurrently without external synchronization.
package nbt
