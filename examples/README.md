# nbtGo Examples

This directory contains example programs demonstrating how to use the nbtGo library.

## Running the Examples

All examples can be run using `go run`:

```bash
cd examples
go run <example-file>.go [arguments]
```

## Available Examples

### 1. parse_to_json.go

Parses a gzip-compressed NBT file and outputs it as formatted JSON.

**Usage:**
```bash
go run examples/parse_to_json.go level.dat
```

**What it demonstrates:**
- Opening and reading NBT files
- Decompressing gzip data
- Parsing NBT binary format
- Converting NBT to JSON

### 2. create_nbt.go

Creates a custom NBT structure programmatically and serializes it to binary format.

**Usage:**
```bash
go run examples/create_nbt.go
```

**What it demonstrates:**
- Creating NBT tag structures in code
- Using different tag types (compound, list, primitives)
- Serializing NBT to binary format
- Writing NBT data to files
- Converting NBT to JSON for verification

### 3. byte_conversion.go

Demonstrates the byte conversion utilities for working with different data types and endianness.

**Usage:**
```bash
go run examples/byte_conversion.go
```

**What it demonstrates:**
- Converting integers to/from bytes
- Converting floats to/from bytes
- Big-endian vs little-endian conversion
- Working with different integer sizes (int16, int32, int64)

## Testing with Minecraft Files

If you have Minecraft installed, you can test these examples with real Minecraft data files:

### Java Edition
```bash
# On Linux
go run examples/parse_to_json.go ~/.minecraft/saves/YourWorld/level.dat

# On Windows
go run examples/parse_to_json.go %APPDATA%\.minecraft\saves\YourWorld\level.dat

# On macOS
go run examples/parse_to_json.go ~/Library/Application\ Support/minecraft/saves/YourWorld/level.dat
```

### Player Data
```bash
# Parse player data file
go run examples/parse_to_json.go ~/.minecraft/saves/YourWorld/playerdata/<uuid>.dat
```

## Creating Your Own Examples

When creating your own programs using nbtGo:

1. Import the necessary packages:
   ```go
   import (
       "goNbt/lib"
       "goNbt/lib/nbt"
   )
   ```

2. For compressed files, use `lib.UnzipReader()`
3. For parsing, use `nbt.ParseNBT(data, isBedrock)`
4. For serializing, use `nbt.SerializeTag(tag, skipHeader)`

## Common Use Cases

### Reading Minecraft World Data
```go
file, _ := os.Open("level.dat")
data, _ := lib.UnzipReader(file)
tag, _ := nbt.ParseNBT(data, false) // false = Java Edition
```

### Creating Custom NBT Data
```go
tag := &nbt.TagCompound{
    Value: []nbt.NBTTag{
        &nbt.TagString{Value: "example"},
        &nbt.TagEnd{},
    },
}
bytes, _ := nbt.SerializeTag(tag, false)
```

### Converting to JSON for Debugging
```go
jsonData, _ := json.MarshalIndent(tag, "", "  ")
fmt.Println(string(jsonData))
```

## Need Help?

- Check the main README.md for API documentation
- Look at the test files in `lib/nbt/*_test.go` for more examples
- Open an issue on GitHub if you have questions

## Contributing Examples

If you create a useful example program, consider contributing it! See CONTRIBUTING.md for guidelines.
