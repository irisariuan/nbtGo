package nbt

import (
	"goNbt/lib"
	"testing"
)

func TestDeserializeTagByte(t *testing.T) {
	// Create NBT data: TAG_Byte named "testByte" with value 42
	data := []byte{
		byte(BTagByte), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(8, true)...) // Name length
	data = append(data, []byte("testByte")...)
	data = append(data, 42) // Value

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagByte: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagByte, ok := tag.(*TagByte)
	if !ok {
		t.Fatalf("Expected *TagByte, got %T", tag)
	}
	if tagByte.Name() != "testByte" {
		t.Errorf("Expected name 'testByte', got '%s'", tagByte.Name())
	}
	if tagByte.Value != 42 {
		t.Errorf("Expected value 42, got %d", tagByte.Value)
	}
}

func TestDeserializeTagShort(t *testing.T) {
	// Create NBT data: TAG_Short named "testShort" with value -1234
	data := []byte{
		byte(BTagShort), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(9, true)...) // Name length
	data = append(data, []byte("testShort")...)
	data = append(data, lib.Int16ToBytes(-1234, true)...) // Value

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagShort: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagShort, ok := tag.(*TagShort)
	if !ok {
		t.Fatalf("Expected *TagShort, got %T", tag)
	}
	if tagShort.Name() != "testShort" {
		t.Errorf("Expected name 'testShort', got '%s'", tagShort.Name())
	}
	if tagShort.Value != -1234 {
		t.Errorf("Expected value -1234, got %d", tagShort.Value)
	}
}

func TestDeserializeTagInt(t *testing.T) {
	// Create NBT data: TAG_Int named "testInt" with value 123456
	data := []byte{
		byte(BTagInt), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(7, true)...) // Name length
	data = append(data, []byte("testInt")...)
	data = append(data, lib.Int32ToBytes(123456, true)...) // Value

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagInt: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagInt, ok := tag.(*TagInt)
	if !ok {
		t.Fatalf("Expected *TagInt, got %T", tag)
	}
	if tagInt.Name() != "testInt" {
		t.Errorf("Expected name 'testInt', got '%s'", tagInt.Name())
	}
	if tagInt.Value != 123456 {
		t.Errorf("Expected value 123456, got %d", tagInt.Value)
	}
}

func TestDeserializeTagLong(t *testing.T) {
	// Create NBT data: TAG_Long named "testLong" with value 9876543210
	data := []byte{
		byte(BTagLong), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(8, true)...) // Name length
	data = append(data, []byte("testLong")...)
	data = append(data, lib.Int64ToBytes(9876543210, true)...) // Value

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagLong: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagLong, ok := tag.(*TagLong)
	if !ok {
		t.Fatalf("Expected *TagLong, got %T", tag)
	}
	if tagLong.Name() != "testLong" {
		t.Errorf("Expected name 'testLong', got '%s'", tagLong.Name())
	}
	if tagLong.Value != 9876543210 {
		t.Errorf("Expected value 9876543210, got %d", tagLong.Value)
	}
}

func TestDeserializeTagFloat(t *testing.T) {
	// Create NBT data: TAG_Float named "testFloat" with value 3.14159
	data := []byte{
		byte(BTagFloat), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(9, true)...) // Name length
	data = append(data, []byte("testFloat")...)
	data = append(data, lib.Float32ToBytes(3.14159, true)...) // Value

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagFloat: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagFloat, ok := tag.(*TagFloat)
	if !ok {
		t.Fatalf("Expected *TagFloat, got %T", tag)
	}
	if tagFloat.Name() != "testFloat" {
		t.Errorf("Expected name 'testFloat', got '%s'", tagFloat.Name())
	}
	// Float comparison with small epsilon
	diff := tagFloat.Value - 3.14159
	if diff < -0.0001 || diff > 0.0001 {
		t.Errorf("Expected value ~3.14159, got %f", tagFloat.Value)
	}
}

func TestDeserializeTagDouble(t *testing.T) {
	// Create NBT data: TAG_Double named "testDouble" with value 2.718281828459045
	data := []byte{
		byte(BTagDouble), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(10, true)...) // Name length
	data = append(data, []byte("testDouble")...)
	data = append(data, lib.Float64ToBytes(2.718281828459045, true)...) // Value

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagDouble: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagDouble, ok := tag.(*TagDouble)
	if !ok {
		t.Fatalf("Expected *TagDouble, got %T", tag)
	}
	if tagDouble.Name() != "testDouble" {
		t.Errorf("Expected name 'testDouble', got '%s'", tagDouble.Name())
	}
	if tagDouble.Value != 2.718281828459045 {
		t.Errorf("Expected value 2.718281828459045, got %f", tagDouble.Value)
	}
}

func TestDeserializeTagString(t *testing.T) {
	// Create NBT data: TAG_String named "testString" with value "Hello, NBT!"
	data := []byte{
		byte(BTagString), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(10, true)...) // Name length
	data = append(data, []byte("testString")...)
	data = append(data, lib.UInt16ToBytes(11, true)...) // String length
	data = append(data, []byte("Hello, NBT!")...)

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagString: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagString, ok := tag.(*TagString)
	if !ok {
		t.Fatalf("Expected *TagString, got %T", tag)
	}
	if tagString.Name() != "testString" {
		t.Errorf("Expected name 'testString', got '%s'", tagString.Name())
	}
	if tagString.Value != "Hello, NBT!" {
		t.Errorf("Expected value 'Hello, NBT!', got '%s'", tagString.Value)
	}
}

func TestDeserializeTagByteArray(t *testing.T) {
	// Create NBT data: TAG_Byte_Array named "testByteArray" with value [1, 2, 3, 4, 5]
	data := []byte{
		byte(BTagByteArray), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(13, true)...) // Name length
	data = append(data, []byte("testByteArray")...)
	data = append(data, lib.Int32ToBytes(5, true)...) // Array length
	data = append(data, []byte{1, 2, 3, 4, 5}...)     // Array data

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagByteArray: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagByteArray, ok := tag.(*TagByteArray)
	if !ok {
		t.Fatalf("Expected *TagByteArray, got %T", tag)
	}
	if tagByteArray.Name() != "testByteArray" {
		t.Errorf("Expected name 'testByteArray', got '%s'", tagByteArray.Name())
	}
	expected := []byte{1, 2, 3, 4, 5}
	if len(tagByteArray.Value) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(tagByteArray.Value))
	}
	for i := range expected {
		if tagByteArray.Value[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], tagByteArray.Value[i])
		}
	}
}

func TestDeserializeTagIntArray(t *testing.T) {
	// Create NBT data: TAG_Int_Array named "testIntArray" with value [100, 200, 300, 400, 500]
	data := []byte{
		byte(BTagIntArray), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(12, true)...) // Name length
	data = append(data, []byte("testIntArray")...)
	data = append(data, lib.Int32ToBytes(5, true)...) // Array length
	data = append(data, lib.Int32ToBytes(100, true)...)
	data = append(data, lib.Int32ToBytes(200, true)...)
	data = append(data, lib.Int32ToBytes(300, true)...)
	data = append(data, lib.Int32ToBytes(400, true)...)
	data = append(data, lib.Int32ToBytes(500, true)...)

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagIntArray: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagIntArray, ok := tag.(*TagIntArray)
	if !ok {
		t.Fatalf("Expected *TagIntArray, got %T", tag)
	}
	if tagIntArray.Name() != "testIntArray" {
		t.Errorf("Expected name 'testIntArray', got '%s'", tagIntArray.Name())
	}
	expected := []int32{100, 200, 300, 400, 500}
	if len(tagIntArray.Value) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(tagIntArray.Value))
	}
	for i := range expected {
		if tagIntArray.Value[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], tagIntArray.Value[i])
		}
	}
}

func TestDeserializeTagLongArray(t *testing.T) {
	// Create NBT data: TAG_Long_Array named "testLongArray" with value [1000, 2000, 3000, 4000, 5000]
	data := []byte{
		byte(BTagLongArray), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(13, true)...) // Name length
	data = append(data, []byte("testLongArray")...)
	data = append(data, lib.Int32ToBytes(5, true)...) // Array length
	data = append(data, lib.Int64ToBytes(1000, true)...)
	data = append(data, lib.Int64ToBytes(2000, true)...)
	data = append(data, lib.Int64ToBytes(3000, true)...)
	data = append(data, lib.Int64ToBytes(4000, true)...)
	data = append(data, lib.Int64ToBytes(5000, true)...)

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagLongArray: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagLongArray, ok := tag.(*TagLongArray)
	if !ok {
		t.Fatalf("Expected *TagLongArray, got %T", tag)
	}
	if tagLongArray.Name() != "testLongArray" {
		t.Errorf("Expected name 'testLongArray', got '%s'", tagLongArray.Name())
	}
	expected := []int64{1000, 2000, 3000, 4000, 5000}
	if len(tagLongArray.Value) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(tagLongArray.Value))
	}
	for i := range expected {
		if tagLongArray.Value[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], tagLongArray.Value[i])
		}
	}
}

func TestDeserializeTagList(t *testing.T) {
	// Create NBT data: TAG_List named "testList" containing 3 TAG_Int values: 10, 20, 30
	data := []byte{
		byte(BTagList), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(8, true)...) // Name length
	data = append(data, []byte("testList")...)
	data = append(data, byte(BTagInt))                // Element type
	data = append(data, lib.Int32ToBytes(3, true)...) // List length
	data = append(data, lib.Int32ToBytes(10, true)...)
	data = append(data, lib.Int32ToBytes(20, true)...)
	data = append(data, lib.Int32ToBytes(30, true)...)

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagList: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagList, ok := tag.(*TagList)
	if !ok {
		t.Fatalf("Expected *TagList, got %T", tag)
	}
	if tagList.Name() != "testList" {
		t.Errorf("Expected name 'testList', got '%s'", tagList.Name())
	}
	if tagList.ElementType != BTagInt {
		t.Errorf("Expected element type BTagInt, got %d", tagList.ElementType)
	}
	if len(tagList.Value) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(tagList.Value))
	}

	expected := []int32{10, 20, 30}
	for i, expectedVal := range expected {
		tagInt, ok := tagList.Value[i].(*TagInt)
		if !ok {
			t.Fatalf("Element %d: expected *TagInt, got %T", i, tagList.Value[i])
		}
		if tagInt.Value != expectedVal {
			t.Errorf("Element %d: expected value %d, got %d", i, expectedVal, tagInt.Value)
		}
	}
}

func TestDeserializeTagCompound(t *testing.T) {
	// Create NBT data: TAG_Compound named "testCompound" containing:
	// - TAG_String "name" = "Test"
	// - TAG_Int "age" = 25
	// - TAG_End
	data := []byte{
		byte(BTagCompound), // Tag type
	}
	data = append(data, lib.UInt16ToBytes(12, true)...) // Name length
	data = append(data, []byte("testCompound")...)

	// First child: TAG_String "name" = "Test"
	data = append(data, byte(BTagString))
	data = append(data, lib.UInt16ToBytes(4, true)...) // Name length
	data = append(data, []byte("name")...)
	data = append(data, lib.UInt16ToBytes(4, true)...) // String length
	data = append(data, []byte("Test")...)

	// Second child: TAG_Int "age" = 25
	data = append(data, byte(BTagInt))
	data = append(data, lib.UInt16ToBytes(3, true)...) // Name length
	data = append(data, []byte("age")...)
	data = append(data, lib.Int32ToBytes(25, true)...)

	// TAG_End
	data = append(data, byte(BTagEnd))

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagCompound: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagCompound, ok := tag.(*TagCompound)
	if !ok {
		t.Fatalf("Expected *TagCompound, got %T", tag)
	}
	if tagCompound.Name() != "testCompound" {
		t.Errorf("Expected name 'testCompound', got '%s'", tagCompound.Name())
	}
	if len(tagCompound.Value) != 3 {
		t.Fatalf("Expected 3 child tags (including TagEnd), got %d", len(tagCompound.Value))
	}

	// Check first child
	tagString, ok := tagCompound.Value[0].(*TagString)
	if !ok {
		t.Fatalf("First child: expected *TagString, got %T", tagCompound.Value[0])
	}
	if tagString.Name() != "name" {
		t.Errorf("First child: expected name 'name', got '%s'", tagString.Name())
	}
	if tagString.Value != "Test" {
		t.Errorf("First child: expected value 'Test', got '%s'", tagString.Value)
	}

	// Check second child
	tagInt, ok := tagCompound.Value[1].(*TagInt)
	if !ok {
		t.Fatalf("Second child: expected *TagInt, got %T", tagCompound.Value[1])
	}
	if tagInt.Name() != "age" {
		t.Errorf("Second child: expected name 'age', got '%s'", tagInt.Name())
	}
	if tagInt.Value != 25 {
		t.Errorf("Second child: expected value 25, got %d", tagInt.Value)
	}

	// Check third child (TagEnd)
	_, ok = tagCompound.Value[2].(*TagEnd)
	if !ok {
		t.Fatalf("Third child: expected *TagEnd, got %T", tagCompound.Value[2])
	}
}

func TestDeserializeTagEnd(t *testing.T) {
	// TAG_End has no name and no payload
	data := []byte{byte(BTagEnd)}

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse TagEnd: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	tagEnd, ok := tag.(*TagEnd)
	if !ok {
		t.Fatalf("Expected *TagEnd, got %T", tag)
	}
	if tagEnd.Type() != BTagEnd {
		t.Errorf("Expected type BTagEnd, got %d", tagEnd.Type())
	}
}

func TestDeserializeNestedStructure(t *testing.T) {
	// Create a nested structure:
	// TAG_Compound "root"
	//   - TAG_String "title" = "Nested Test"
	//   - TAG_List "numbers" (TAG_Int) = [1, 2, 3]
	//   - TAG_Compound "nested"
	//     - TAG_Byte "flag" = 1
	//     - TAG_End
	//   - TAG_End

	data := []byte{
		byte(BTagCompound), // Root tag type
	}
	data = append(data, lib.UInt16ToBytes(4, true)...) // Name length
	data = append(data, []byte("root")...)

	// Child 1: TAG_String "title" = "Nested Test"
	data = append(data, byte(BTagString))
	data = append(data, lib.UInt16ToBytes(5, true)...) // Name length
	data = append(data, []byte("title")...)
	data = append(data, lib.UInt16ToBytes(11, true)...) // String length
	data = append(data, []byte("Nested Test")...)

	// Child 2: TAG_List "numbers" containing [1, 2, 3]
	data = append(data, byte(BTagList))
	data = append(data, lib.UInt16ToBytes(7, true)...) // Name length
	data = append(data, []byte("numbers")...)
	data = append(data, byte(BTagInt))                // Element type
	data = append(data, lib.Int32ToBytes(3, true)...) // List length
	data = append(data, lib.Int32ToBytes(1, true)...)
	data = append(data, lib.Int32ToBytes(2, true)...)
	data = append(data, lib.Int32ToBytes(3, true)...)

	// Child 3: TAG_Compound "nested"
	data = append(data, byte(BTagCompound))
	data = append(data, lib.UInt16ToBytes(6, true)...) // Name length
	data = append(data, []byte("nested")...)

	// Nested child: TAG_Byte "flag" = 1
	data = append(data, byte(BTagByte))
	data = append(data, lib.UInt16ToBytes(4, true)...) // Name length
	data = append(data, []byte("flag")...)
	data = append(data, 1) // Value

	// End nested compound
	data = append(data, byte(BTagEnd))

	// End root compound
	data = append(data, byte(BTagEnd))

	tag, remaining, err := separateSingleTag(data, 0, true)
	if err != nil {
		t.Fatalf("Failed to parse nested structure: %v", err)
	}
	if len(remaining) != 0 {
		t.Errorf("Expected no remaining bytes, got %d", len(remaining))
	}

	root, ok := tag.(*TagCompound)
	if !ok {
		t.Fatalf("Expected *TagCompound, got %T", tag)
	}
	if root.Name() != "root" {
		t.Errorf("Expected name 'root', got '%s'", root.Name())
	}
	if len(root.Value) != 4 {
		t.Fatalf("Expected 4 children (including TagEnd), got %d", len(root.Value))
	}

	// Verify the nested list
	tagList, ok := root.Value[1].(*TagList)
	if !ok {
		t.Fatalf("Second child: expected *TagList, got %T", root.Value[1])
	}
	if len(tagList.Value) != 3 {
		t.Errorf("List: expected 3 elements, got %d", len(tagList.Value))
	}

	// Verify the nested compound
	nestedCompound, ok := root.Value[2].(*TagCompound)
	if !ok {
		t.Fatalf("Third child: expected *TagCompound, got %T", root.Value[2])
	}
	if nestedCompound.Name() != "nested" {
		t.Errorf("Nested compound: expected name 'nested', got '%s'", nestedCompound.Name())
	}
	if len(nestedCompound.Value) != 2 {
		t.Fatalf("Nested compound: expected 2 children (including TagEnd), got %d", len(nestedCompound.Value))
	}

	// Verify the byte in nested compound
	flagByte, ok := nestedCompound.Value[0].(*TagByte)
	if !ok {
		t.Fatalf("Nested byte: expected *TagByte, got %T", nestedCompound.Value[0])
	}
	if flagByte.Value != 1 {
		t.Errorf("Nested byte: expected value 1, got %d", flagByte.Value)
	}

	// Verify TagEnd in nested compound
	_, ok = nestedCompound.Value[1].(*TagEnd)
	if !ok {
		t.Fatalf("Nested compound second child: expected *TagEnd, got %T", nestedCompound.Value[1])
	}

	// Verify the fourth child is TagEnd
	_, ok = root.Value[3].(*TagEnd)
	if !ok {
		t.Fatalf("Fourth child: expected *TagEnd, got %T", root.Value[3])
	}
}

func TestParseNBTRoot(t *testing.T) {
	// Test the ParseNBT function with a complete NBT structure
	// Create a simple TAG_Compound root
	data := []byte{
		byte(BTagCompound), // Root tag type
	}
	data = append(data, lib.UInt16ToBytes(0, true)...) // Empty name (common for root)

	// Add a single TAG_Int child
	data = append(data, byte(BTagInt))
	data = append(data, lib.UInt16ToBytes(5, true)...) // Name length
	data = append(data, []byte("value")...)
	data = append(data, lib.Int32ToBytes(42, true)...)

	// End compound
	data = append(data, byte(BTagEnd))

	tag, err := ParseNBT(data, false)
	if err != nil {
		t.Fatalf("Failed to parse NBT: %v", err)
	}

	compound, ok := tag.(*TagCompound)
	if !ok {
		t.Fatalf("Expected *TagCompound, got %T", tag)
	}
	if len(compound.Value) != 2 {
		t.Fatalf("Expected 2 children (including TagEnd), got %d", len(compound.Value))
	}

	childInt, ok := compound.Value[0].(*TagInt)
	if !ok {
		t.Fatalf("Child: expected *TagInt, got %T", compound.Value[0])
	}
	if childInt.Value != 42 {
		t.Errorf("Child: expected value 42, got %d", childInt.Value)
	}

	// Verify second child is TagEnd
	_, ok = compound.Value[1].(*TagEnd)
	if !ok {
		t.Fatalf("Second child: expected *TagEnd, got %T", compound.Value[1])
	}
}
