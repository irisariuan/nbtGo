# nbtGo

A Go library for parsing and serializing NBT (Named Binary Tag) format, the data structure used by Minecraft for storing game data.

## Features

- ✅ Full NBT specification support
- ✅ Parse NBT binary data into Go structures
- ✅ Serialize Go structures back to NBT format
- ✅ JSON marshaling/unmarshaling for easy data inspection
- ✅ Support for both Java Edition (big-endian) and Bedrock Edition formats
- ✅ Gzip compression support
- ✅ All NBT tag types supported

## Installation

```bash
go get github.com/irisariuan/nbtGo
```

## Supported NBT Tag Types

The library supports all standard NBT tag types:

- `TAG_End` - Marks the end of a compound tag
- `TAG_Byte` - Single signed byte
- `TAG_Short` - Signed 16-bit integer
- `TAG_Int` - Signed 32-bit integer
- `TAG_Long` - Signed 64-bit integer
- `TAG_Float` - 32-bit floating point
- `TAG_Double` - 64-bit floating point
- `TAG_Byte_Array` - Array of signed bytes
- `TAG_String` - UTF-8 string
- `TAG_List` - List of unnamed tags (homogeneous)
- `TAG_Compound` - Collection of named tags
- `TAG_Int_Array` - Array of signed 32-bit integers
- `TAG_Long_Array` - Array of signed 64-bit integers

## Usage

### Command Line Tool

The main application reads gzip-compressed NBT data from stdin and outputs it as formatted JSON:

```bash
# Parse a Minecraft level.dat file
cat level.dat | go run main.go

# Or build and use the binary
go build -o nbtGo
cat level.dat | ./nbtGo
```

### As a Library

#### Parsing NBT Data

```go
package main

import (
    "fmt"
    "goNbt/lib/nbt"
    "os"
)

func main() {
    // Read NBT data (raw bytes, not gzipped)
    data, err := os.ReadFile("data.nbt")
    if err != nil {
        panic(err)
    }

    // Parse NBT data (false = Java Edition / big-endian)
    tag, err := nbt.ParseNBT(data, false)
    if err != nil {
        panic(err)
    }

    // Work with the parsed data
    nbt.PrintTag(tag)
}
```

#### Serializing NBT Data

```go
package main

import (
    "goNbt/lib/nbt"
)

func main() {
    // Create an NBT structure
    tag := &nbt.TagCompound{
        baseTag: nbt.baseTag{
            tagType: nbt.BTagCompound,
            name:    "root",
        },
        Value: []nbt.NBTTag{
            &nbt.TagString{
                baseTag: nbt.baseTag{
                    tagType: nbt.BTagString,
                    name:    "hello",
                },
                Value: "world",
            },
            &nbt.TagInt{
                baseTag: nbt.baseTag{
                    tagType: nbt.BTagInt,
                    name:    "count",
                },
                Value: 42,
            },
            &nbt.TagEnd{},
        },
    }

    // Serialize to bytes
    bytes, err := nbt.SerializeTag(tag, false)
    if err != nil {
        panic(err)
    }

    // bytes now contains the NBT binary data
}
```

#### JSON Conversion

```go
package main

import (
    "encoding/json"
    "fmt"
    "goNbt/lib/nbt"
)

func main() {
    // Assuming you have a parsed NBT tag
    tag := &nbt.TagString{
        baseTag: nbt.baseTag{
            tagType: nbt.BTagString,
            name:    "example",
        },
        Value: "Hello, NBT!",
    }

    // Convert to JSON
    jsonData, err := json.MarshalIndent(tag, "", "  ")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(jsonData))
    // Output:
    // {
    //   "type": "string",
    //   "name": "example",
    //   "value": "Hello, NBT!"
    // }
}
```

#### Working with Gzip Compression

```go
package main

import (
    "goNbt/lib"
    "goNbt/lib/nbt"
    "os"
)

func main() {
    // Read and decompress gzip-compressed NBT file
    file, err := os.Open("level.dat")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Unzip the data
    data, err := lib.UnzipReader(file)
    if err != nil {
        panic(err)
    }

    // Parse the NBT data
    tag, err := nbt.ParseNBT(data, false)
    if err != nil {
        panic(err)
    }

    // Work with the parsed tag
    nbt.PrintTag(tag)
}
```

#### Byte Conversion Utilities

The library provides utilities for converting between bytes and various data types:

```go
package main

import (
    "fmt"
    "goNbt/lib"
)

func main() {
    // Convert int32 to bytes (big-endian)
    bytes := lib.Int32ToBytes(12345, true)
    fmt.Printf("Bytes: %v\n", bytes)

    // Convert bytes back to int32
    value, err := lib.BytesToInt32(bytes, true)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Value: %d\n", value)

    // Similar functions exist for:
    // - Int16ToBytes / BytesToInt16
    // - Int64ToBytes / BytesToInt64
    // - Float32ToBytes / BytesFloat32
    // - Float64ToBytes / BytesFloat64
    // - UInt16ToBytes / BytesToUInt16
    // - BytesToUInt32 / BytesToUInt64
}
```

## API Reference

### Package `nbt`

#### Types

##### `NBTTag` Interface

All tag types implement the `NBTTag` interface:

```go
type NBTTag interface {
    Type() tagTypeByte
    Name() string
    DataLength() int
    ZIndex() int
}
```

##### Tag Structures

- `TagByte` - Contains a `Value byte` field
- `TagShort` - Contains a `Value int16` field
- `TagInt` - Contains a `Value int32` field
- `TagLong` - Contains a `Value int64` field
- `TagFloat` - Contains a `Value float32` field
- `TagDouble` - Contains a `Value float64` field
- `TagString` - Contains a `Value string` field
- `TagByteArray` - Contains a `Value []byte` field
- `TagIntArray` - Contains a `Value []int32` field
- `TagLongArray` - Contains a `Value []int64` field
- `TagList` - Contains `ElementType tagTypeByte` and `Value []NBTTag` fields
- `TagCompound` - Contains a `Value []NBTTag` field
- `TagEnd` - Marker tag with no value

#### Functions

##### `ParseNBT(data []byte, isBedrock bool) (NBTTag, TagParseError)`

Parses NBT binary data into a tag structure.

- `data`: Raw NBT bytes
- `isBedrock`: Set to `true` for Bedrock Edition (little-endian), `false` for Java Edition (big-endian)
- Returns: Parsed root tag (must be `TAG_Compound` or `TAG_List`) and error

##### `SerializeTag(tag NBTTag, skipHeader bool) ([]byte, error)`

Serializes an NBT tag to binary format.

- `tag`: The tag to serialize
- `skipHeader`: If `true`, skips the tag type and name header (used for list elements)
- Returns: Serialized bytes and error

##### `PrintTag(tag NBTTag)`

Prints tag information to stdout for debugging purposes.

### Package `lib`

#### Functions

##### Decompression

- `UnzipReader(reader io.Reader) ([]byte, error)` - Reads and decompresses gzip data

##### Byte Conversion (Big-Endian and Little-Endian Support)

All conversion functions take a `bigEndian bool` parameter:

**Integer Conversions:**
- `BytesToInt16(bytes []byte, bigEndian bool) (int16, error)`
- `BytesToInt32(bytes []byte, bigEndian bool) (int32, error)`
- `BytesToInt64(bytes []byte, bigEndian bool) (int64, error)`
- `BytesToUInt16(bytes []byte, bigEndian bool) (uint16, error)`
- `BytesToUInt32(bytes []byte, bigEndian bool) (uint32, error)`
- `BytesToUInt64(bytes []byte, bigEndian bool) (uint64, error)`

**Integer to Bytes:**
- `Int16ToBytes(val int16, bigEndian bool) []byte`
- `Int32ToBytes(val int32, bigEndian bool) []byte`
- `Int64ToBytes(val int64, bigEndian bool) []byte`
- `UInt16ToBytes(val uint16, bigEndian bool) []byte`

**Float Conversions:**
- `BytesFloat32(bytes []byte, bigEndian bool) (float32, error)`
- `BytesFloat64(bytes []byte, bigEndian bool) (float64, error)`
- `Float32ToBytes(val float32, bigEndian bool) []byte`
- `Float64ToBytes(val float64, bigEndian bool) []byte`

## Building

```bash
# Build the command-line tool
go build -o nbtGo

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Testing

The library includes comprehensive tests for all tag types:

```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./lib/nbt

# Verbose output
go test -v ./...
```

## NBT Format Reference

NBT is a structured binary format consisting of:

1. **Tag Type** (1 byte) - Identifies the type of tag
2. **Name Length** (2 bytes, unsigned, big-endian) - Length of the tag name
3. **Name** (UTF-8 string) - The name of the tag
4. **Payload** (variable) - The actual data, format depends on tag type

For more details on the NBT format, see the [Minecraft Wiki NBT Documentation](https://minecraft.wiki/w/NBT_format).

## License

This project is open source. Please check the repository for license information.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## Examples

### Example: Reading Minecraft Level Data

```go
package main

import (
    "encoding/json"
    "fmt"
    "goNbt/lib"
    "goNbt/lib/nbt"
    "os"
)

func main() {
    // Open a Minecraft level.dat file
    file, err := os.Open("level.dat")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Decompress the gzip data
    data, err := lib.UnzipReader(file)
    if err != nil {
        panic(err)
    }

    // Parse the NBT data
    tag, err := nbt.ParseNBT(data, false)
    if err != nil {
        panic(err)
    }

    // Convert to JSON for easy viewing
    jsonData, err := json.MarshalIndent(tag, "", "  ")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(jsonData))
}
```

### Example: Creating Custom NBT Structures

```go
package main

import (
    "goNbt/lib/nbt"
    "os"
)

func main() {
    // Create a player inventory structure
    playerData := &nbt.TagCompound{
        Value: []nbt.NBTTag{
            &nbt.TagString{
                baseTag: nbt.baseTag{tagType: nbt.BTagString, name: "playerName"},
                Value:   "Steve",
            },
            &nbt.TagInt{
                baseTag: nbt.baseTag{tagType: nbt.BTagInt, name: "health"},
                Value:   20,
            },
            &nbt.TagList{
                baseTag:     nbt.baseTag{tagType: nbt.BTagList, name: "inventory"},
                ElementType: nbt.BTagCompound,
                Value:       []nbt.NBTTag{
                    // Add inventory items here
                },
            },
            &nbt.TagEnd{},
        },
    }

    // Serialize the structure
    bytes, err := nbt.SerializeTag(playerData, false)
    if err != nil {
        panic(err)
    }

    // Write to file
    os.WriteFile("player.nbt", bytes, 0644)
}
```

## Related Projects

- [Minecraft NBT Format Specification](https://minecraft.wiki/w/NBT_format)
- [NBT Editor Tools](https://github.com/topics/nbt-editor)

## Authors

- irisariuan

## Acknowledgments

Special thanks to the Minecraft community for documenting the NBT format specification.
