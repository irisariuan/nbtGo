package nbt

import (
	"bytes"
	"goNbt/lib"
	"testing"
)

func TestSerializeTagByte(t *testing.T) {
	tag := &TagByte{
		baseTag: baseTag{tagType: BTagByte, name: "testByte"},
		Value:   42,
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagByte: %v", err)
	}

	// Expected format: [type(1)][name_length(2)][name][value(1)]
	expected := []byte{byte(BTagByte)}
	expected = append(expected, lib.UInt16ToBytes(8, true)...)
	expected = append(expected, []byte("testByte")...)
	expected = append(expected, 42)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagByteSkipHeader(t *testing.T) {
	tag := &TagByte{
		baseTag: baseTag{tagType: BTagByte, name: "testByte"},
		Value:   42,
	}

	data, err := SerializeTag(tag, true)
	if err != nil {
		t.Fatalf("Failed to serialize TagByte with skipHeader: %v", err)
	}

	expected := []byte{42}
	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagShort(t *testing.T) {
	tag := &TagShort{
		baseTag: baseTag{tagType: BTagShort, name: "testShort"},
		Value:   -1234,
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagShort: %v", err)
	}

	expected := []byte{byte(BTagShort)}
	expected = append(expected, lib.UInt16ToBytes(9, true)...)
	expected = append(expected, []byte("testShort")...)
	expected = append(expected, lib.Int16ToBytes(-1234, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagInt(t *testing.T) {
	tag := &TagInt{
		baseTag: baseTag{tagType: BTagInt, name: "testInt"},
		Value:   123456,
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagInt: %v", err)
	}

	expected := []byte{byte(BTagInt)}
	expected = append(expected, lib.UInt16ToBytes(7, true)...)
	expected = append(expected, []byte("testInt")...)
	expected = append(expected, lib.Int32ToBytes(123456, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagLong(t *testing.T) {
	tag := &TagLong{
		baseTag: baseTag{tagType: BTagLong, name: "testLong"},
		Value:   9876543210,
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagLong: %v", err)
	}

	expected := []byte{byte(BTagLong)}
	expected = append(expected, lib.UInt16ToBytes(8, true)...)
	expected = append(expected, []byte("testLong")...)
	expected = append(expected, lib.Int64ToBytes(9876543210, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagFloat(t *testing.T) {
	tag := &TagFloat{
		baseTag: baseTag{tagType: BTagFloat, name: "testFloat"},
		Value:   3.14159,
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagFloat: %v", err)
	}

	expected := []byte{byte(BTagFloat)}
	expected = append(expected, lib.UInt16ToBytes(9, true)...)
	expected = append(expected, []byte("testFloat")...)
	expected = append(expected, lib.Float32ToBytes(3.14159, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagDouble(t *testing.T) {
	tag := &TagDouble{
		baseTag: baseTag{tagType: BTagDouble, name: "testDouble"},
		Value:   2.718281828459045,
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagDouble: %v", err)
	}

	expected := []byte{byte(BTagDouble)}
	expected = append(expected, lib.UInt16ToBytes(10, true)...)
	expected = append(expected, []byte("testDouble")...)
	expected = append(expected, lib.Float64ToBytes(2.718281828459045, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagString(t *testing.T) {
	tag := &TagString{
		baseTag: baseTag{tagType: BTagString, name: "testString"},
		Value:   "Hello, NBT!",
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagString: %v", err)
	}

	expected := []byte{byte(BTagString)}
	expected = append(expected, lib.UInt16ToBytes(10, true)...)
	expected = append(expected, []byte("testString")...)
	expected = append(expected, lib.UInt16ToBytes(11, true)...)
	expected = append(expected, []byte("Hello, NBT!")...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagByteArray(t *testing.T) {
	tag := &TagByteArray{
		baseTag: baseTag{tagType: BTagByteArray, name: "testByteArray"},
		Value:   []byte{1, 2, 3, 4, 5},
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagByteArray: %v", err)
	}

	expected := []byte{byte(BTagByteArray)}
	expected = append(expected, lib.UInt16ToBytes(13, true)...)
	expected = append(expected, []byte("testByteArray")...)
	expected = append(expected, lib.Int32ToBytes(5, true)...)
	expected = append(expected, []byte{1, 2, 3, 4, 5}...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagIntArray(t *testing.T) {
	tag := &TagIntArray{
		baseTag: baseTag{tagType: BTagIntArray, name: "testIntArray"},
		Value:   []int32{100, 200, 300, 400, 500},
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagIntArray: %v", err)
	}

	expected := []byte{byte(BTagIntArray)}
	expected = append(expected, lib.UInt16ToBytes(12, true)...)
	expected = append(expected, []byte("testIntArray")...)
	expected = append(expected, lib.Int32ToBytes(5, true)...)
	expected = append(expected, lib.Int32ToBytes(100, true)...)
	expected = append(expected, lib.Int32ToBytes(200, true)...)
	expected = append(expected, lib.Int32ToBytes(300, true)...)
	expected = append(expected, lib.Int32ToBytes(400, true)...)
	expected = append(expected, lib.Int32ToBytes(500, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagLongArray(t *testing.T) {
	tag := &TagLongArray{
		baseTag: baseTag{tagType: BTagLongArray, name: "testLongArray"},
		Value:   []int64{1000, 2000, 3000, 4000, 5000},
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagLongArray: %v", err)
	}

	expected := []byte{byte(BTagLongArray)}
	expected = append(expected, lib.UInt16ToBytes(13, true)...)
	expected = append(expected, []byte("testLongArray")...)
	expected = append(expected, lib.Int32ToBytes(5, true)...)
	expected = append(expected, lib.Int64ToBytes(1000, true)...)
	expected = append(expected, lib.Int64ToBytes(2000, true)...)
	expected = append(expected, lib.Int64ToBytes(3000, true)...)
	expected = append(expected, lib.Int64ToBytes(4000, true)...)
	expected = append(expected, lib.Int64ToBytes(5000, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagList(t *testing.T) {
	tag := &TagList{
		baseTag:     baseTag{tagType: BTagList, name: "testList"},
		ElementType: BTagInt,
		Value: []NBTTag{
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 10},
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 20},
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 30},
		},
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagList: %v", err)
	}

	expected := []byte{byte(BTagList)}
	expected = append(expected, lib.UInt16ToBytes(8, true)...)
	expected = append(expected, []byte("testList")...)
	expected = append(expected, byte(BTagInt))
	expected = append(expected, lib.Int32ToBytes(3, true)...)
	expected = append(expected, lib.Int32ToBytes(10, true)...)
	expected = append(expected, lib.Int32ToBytes(20, true)...)
	expected = append(expected, lib.Int32ToBytes(30, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagCompound(t *testing.T) {
	tag := &TagCompound{
		baseTag: baseTag{tagType: BTagCompound, name: "testCompound"},
		Value: []NBTTag{
			&TagString{baseTag: baseTag{tagType: BTagString, name: "name"}, Value: "Test"},
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: "age"}, Value: 25},
			&TagEnd{baseTag: baseTag{tagType: BTagEnd, name: ""}},
		},
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagCompound: %v", err)
	}

	expected := []byte{byte(BTagCompound)}
	expected = append(expected, lib.UInt16ToBytes(12, true)...)
	expected = append(expected, []byte("testCompound")...)

	// Child 1: TAG_String "name" = "Test"
	expected = append(expected, byte(BTagString))
	expected = append(expected, lib.UInt16ToBytes(4, true)...)
	expected = append(expected, []byte("name")...)
	expected = append(expected, lib.UInt16ToBytes(4, true)...)
	expected = append(expected, []byte("Test")...)

	// Child 2: TAG_Int "age" = 25
	expected = append(expected, byte(BTagInt))
	expected = append(expected, lib.UInt16ToBytes(3, true)...)
	expected = append(expected, []byte("age")...)
	expected = append(expected, lib.Int32ToBytes(25, true)...)

	// TAG_End
	expected = append(expected, byte(BTagEnd))
	expected = append(expected, lib.UInt16ToBytes(0, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeTagEnd(t *testing.T) {
	tag := &TagEnd{
		baseTag: baseTag{tagType: BTagEnd, name: ""},
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize TagEnd: %v", err)
	}

	expected := []byte{byte(BTagEnd)}
	expected = append(expected, lib.UInt16ToBytes(0, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot:      %v\nExpected: %v", data, expected)
	}
}

func TestSerializeNestedStructure(t *testing.T) {
	tag := &TagCompound{
		baseTag: baseTag{tagType: BTagCompound, name: "root"},
		Value: []NBTTag{
			&TagString{baseTag: baseTag{tagType: BTagString, name: "title"}, Value: "Nested Test"},
			&TagList{
				baseTag:     baseTag{tagType: BTagList, name: "numbers"},
				ElementType: BTagInt,
				Value: []NBTTag{
					&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 1},
					&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 2},
					&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 3},
				},
			},
			&TagCompound{
				baseTag: baseTag{tagType: BTagCompound, name: "nested"},
				Value: []NBTTag{
					&TagByte{baseTag: baseTag{tagType: BTagByte, name: "flag"}, Value: 1},
					&TagEnd{baseTag: baseTag{tagType: BTagEnd, name: ""}},
				},
			},
			&TagEnd{baseTag: baseTag{tagType: BTagEnd, name: ""}},
		},
	}

	data, err := SerializeTag(tag, false)
	if err != nil {
		t.Fatalf("Failed to serialize nested structure: %v", err)
	}

	// Build expected data
	expected := []byte{byte(BTagCompound)}
	expected = append(expected, lib.UInt16ToBytes(4, true)...)
	expected = append(expected, []byte("root")...)

	// Child 1: TAG_String "title" = "Nested Test"
	expected = append(expected, byte(BTagString))
	expected = append(expected, lib.UInt16ToBytes(5, true)...)
	expected = append(expected, []byte("title")...)
	expected = append(expected, lib.UInt16ToBytes(11, true)...)
	expected = append(expected, []byte("Nested Test")...)

	// Child 2: TAG_List "numbers"
	expected = append(expected, byte(BTagList))
	expected = append(expected, lib.UInt16ToBytes(7, true)...)
	expected = append(expected, []byte("numbers")...)
	expected = append(expected, byte(BTagInt))
	expected = append(expected, lib.Int32ToBytes(3, true)...)
	expected = append(expected, lib.Int32ToBytes(1, true)...)
	expected = append(expected, lib.Int32ToBytes(2, true)...)
	expected = append(expected, lib.Int32ToBytes(3, true)...)

	// Child 3: TAG_Compound "nested"
	expected = append(expected, byte(BTagCompound))
	expected = append(expected, lib.UInt16ToBytes(6, true)...)
	expected = append(expected, []byte("nested")...)

	// Nested child: TAG_Byte "flag" = 1
	expected = append(expected, byte(BTagByte))
	expected = append(expected, lib.UInt16ToBytes(4, true)...)
	expected = append(expected, []byte("flag")...)
	expected = append(expected, 1)

	// Nested TAG_End
	expected = append(expected, byte(BTagEnd))
	expected = append(expected, lib.UInt16ToBytes(0, true)...)

	// Root TAG_End
	expected = append(expected, byte(BTagEnd))
	expected = append(expected, lib.UInt16ToBytes(0, true)...)

	if !bytes.Equal(data, expected) {
		t.Errorf("Serialized data mismatch.\nGot length:      %d\nExpected length: %d", len(data), len(expected))
	}
}

func TestSerializeDeserializeRoundTrip(t *testing.T) {
	// Test that serialize -> deserialize produces the same data
	original := &TagCompound{
		baseTag: baseTag{tagType: BTagCompound, name: "test"},
		Value: []NBTTag{
			&TagByte{baseTag: baseTag{tagType: BTagByte, name: "byte"}, Value: 42},
			&TagShort{baseTag: baseTag{tagType: BTagShort, name: "short"}, Value: -1234},
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: "int"}, Value: 123456},
			&TagLong{baseTag: baseTag{tagType: BTagLong, name: "long"}, Value: 9876543210},
			&TagFloat{baseTag: baseTag{tagType: BTagFloat, name: "float"}, Value: 3.14159},
			&TagDouble{baseTag: baseTag{tagType: BTagDouble, name: "double"}, Value: 2.718281828459045},
			&TagString{baseTag: baseTag{tagType: BTagString, name: "string"}, Value: "Hello, NBT!"},
			&TagByteArray{baseTag: baseTag{tagType: BTagByteArray, name: "byteArray"}, Value: []byte{1, 2, 3, 4, 5}},
			&TagIntArray{baseTag: baseTag{tagType: BTagIntArray, name: "intArray"}, Value: []int32{100, 200, 300}},
			&TagLongArray{baseTag: baseTag{tagType: BTagLongArray, name: "longArray"}, Value: []int64{1000, 2000, 3000}},
			&TagList{
				baseTag:     baseTag{tagType: BTagList, name: "list"},
				ElementType: BTagInt,
				Value: []NBTTag{
					&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 10},
					&TagInt{baseTag: baseTag{tagType: BTagInt, name: ""}, Value: 20},
				},
			},
			&TagEnd{baseTag: baseTag{tagType: BTagEnd, name: ""}},
		},
	}

	// Serialize
	serialized, err := SerializeTag(original, false)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// Deserialize
	deserialized, _, err := separateSingleTag(serialized, 0, true)
	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}
	// Note: There may be remaining bytes (e.g., TagEnd name length) which is expected
	// when deserializing without ParseNBT wrapper

	// Verify type
	compound, ok := deserialized.(*TagCompound)
	if !ok {
		t.Fatalf("Expected *TagCompound, got %T", deserialized)
	}

	// Verify name
	if compound.Name() != "test" {
		t.Errorf("Expected name 'test', got '%s'", compound.Name())
	}

	// Verify we have all children (including TagEnd)
	if len(compound.Value) != len(original.Value) {
		t.Errorf("Expected %d children, got %d", len(original.Value), len(compound.Value))
	}

	// Spot check a few values
	if byteTag, ok := compound.Value[0].(*TagByte); ok {
		if byteTag.Value != 42 {
			t.Errorf("Byte value mismatch: expected 42, got %d", byteTag.Value)
		}
	}

	if stringTag, ok := compound.Value[6].(*TagString); ok {
		if stringTag.Value != "Hello, NBT!" {
			t.Errorf("String value mismatch: expected 'Hello, NBT!', got '%s'", stringTag.Value)
		}
	}
}
