package nbt

import (
	"encoding/json"
	"testing"
)

func TestTagByteJSON(t *testing.T) {
	original := &TagByte{
		baseTag: baseTag{tagType: BTagByte, name: "testByte"},
		Value:   42,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagByte: %v", err)
	}

	var unmarshaled TagByte
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagByte: %v", err)
	}

	if unmarshaled.Value != original.Value {
		t.Errorf("Expected value %d, got %d", original.Value, unmarshaled.Value)
	}
	if unmarshaled.name != original.name {
		t.Errorf("Expected name %s, got %s", original.name, unmarshaled.name)
	}
}

func TestTagShortJSON(t *testing.T) {
	original := &TagShort{
		baseTag: baseTag{tagType: BTagShort, name: "testShort"},
		Value:   -1234,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagShort: %v", err)
	}

	var unmarshaled TagShort
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagShort: %v", err)
	}

	if unmarshaled.Value != original.Value {
		t.Errorf("Expected value %d, got %d", original.Value, unmarshaled.Value)
	}
	if unmarshaled.name != original.name {
		t.Errorf("Expected name %s, got %s", original.name, unmarshaled.name)
	}
}

func TestTagIntJSON(t *testing.T) {
	original := &TagInt{
		baseTag: baseTag{tagType: BTagInt, name: "testInt"},
		Value:   123456,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagInt: %v", err)
	}

	var unmarshaled TagInt
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagInt: %v", err)
	}

	if unmarshaled.Value != original.Value {
		t.Errorf("Expected value %d, got %d", original.Value, unmarshaled.Value)
	}
	if unmarshaled.name != original.name {
		t.Errorf("Expected name %s, got %s", original.name, unmarshaled.name)
	}
}

func TestTagLongJSON(t *testing.T) {
	original := &TagLong{
		baseTag: baseTag{tagType: BTagLong, name: "testLong"},
		Value:   9876543210,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagLong: %v", err)
	}

	var unmarshaled TagLong
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagLong: %v", err)
	}

	if unmarshaled.Value != original.Value {
		t.Errorf("Expected value %d, got %d", original.Value, unmarshaled.Value)
	}
	if unmarshaled.name != original.name {
		t.Errorf("Expected name %s, got %s", original.name, unmarshaled.name)
	}
}

func TestTagFloatJSON(t *testing.T) {
	original := &TagFloat{
		baseTag: baseTag{tagType: BTagFloat, name: "testFloat"},
		Value:   3.14159,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagFloat: %v", err)
	}

	var unmarshaled TagFloat
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagFloat: %v", err)
	}

	if unmarshaled.Value != original.Value {
		t.Errorf("Expected value %f, got %f", original.Value, unmarshaled.Value)
	}
	if unmarshaled.name != original.name {
		t.Errorf("Expected name %s, got %s", original.name, unmarshaled.name)
	}
}

func TestTagDoubleJSON(t *testing.T) {
	original := &TagDouble{
		baseTag: baseTag{tagType: BTagDouble, name: "testDouble"},
		Value:   2.718281828459045,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagDouble: %v", err)
	}

	var unmarshaled TagDouble
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagDouble: %v", err)
	}

	if unmarshaled.Value != original.Value {
		t.Errorf("Expected value %f, got %f", original.Value, unmarshaled.Value)
	}
	if unmarshaled.name != original.name {
		t.Errorf("Expected name %s, got %s", original.name, unmarshaled.name)
	}
}

func TestTagStringJSON(t *testing.T) {
	original := &TagString{
		baseTag: baseTag{tagType: BTagString, name: "testString"},
		Value:   "Hello, NBT!",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagString: %v", err)
	}

	var unmarshaled TagString
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagString: %v", err)
	}

	if unmarshaled.Value != original.Value {
		t.Errorf("Expected value %s, got %s", original.Value, unmarshaled.Value)
	}
	if unmarshaled.name != original.name {
		t.Errorf("Expected name %s, got %s", original.name, unmarshaled.name)
	}
}

func TestTagByteArrayJSON(t *testing.T) {
	original := &TagByteArray{
		baseTag: baseTag{tagType: BTagByteArray, name: "testByteArray"},
		Value:   []byte{1, 2, 3, 4, 5},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagByteArray: %v", err)
	}

	var unmarshaled TagByteArray
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagByteArray: %v", err)
	}

	if len(unmarshaled.Value) != len(original.Value) {
		t.Errorf("Expected length %d, got %d", len(original.Value), len(unmarshaled.Value))
	}
	for i := range original.Value {
		if unmarshaled.Value[i] != original.Value[i] {
			t.Errorf("At index %d: expected %d, got %d", i, original.Value[i], unmarshaled.Value[i])
		}
	}
}

func TestTagIntArrayJSON(t *testing.T) {
	original := &TagIntArray{
		baseTag: baseTag{tagType: BTagIntArray, name: "testIntArray"},
		Value:   []int32{100, 200, 300, 400, 500},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagIntArray: %v", err)
	}

	var unmarshaled TagIntArray
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagIntArray: %v", err)
	}

	if len(unmarshaled.Value) != len(original.Value) {
		t.Errorf("Expected length %d, got %d", len(original.Value), len(unmarshaled.Value))
	}
	for i := range original.Value {
		if unmarshaled.Value[i] != original.Value[i] {
			t.Errorf("At index %d: expected %d, got %d", i, original.Value[i], unmarshaled.Value[i])
		}
	}
}

func TestTagLongArrayJSON(t *testing.T) {
	original := &TagLongArray{
		baseTag: baseTag{tagType: BTagLongArray, name: "testLongArray"},
		Value:   []int64{1000, 2000, 3000, 4000, 5000},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagLongArray: %v", err)
	}

	var unmarshaled TagLongArray
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagLongArray: %v", err)
	}

	if len(unmarshaled.Value) != len(original.Value) {
		t.Errorf("Expected length %d, got %d", len(original.Value), len(unmarshaled.Value))
	}
	for i := range original.Value {
		if unmarshaled.Value[i] != original.Value[i] {
			t.Errorf("At index %d: expected %d, got %d", i, original.Value[i], unmarshaled.Value[i])
		}
	}
}

func TestTagListJSON(t *testing.T) {
	original := &TagList{
		baseTag:     baseTag{tagType: BTagList, name: "testList"},
		ElementType: BTagInt,
		Value: []NBTTag{
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: "item1"}, Value: 10},
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: "item2"}, Value: 20},
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: "item3"}, Value: 30},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagList: %v", err)
	}

	var unmarshaled TagList
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagList: %v", err)
	}

	if unmarshaled.ElementType != original.ElementType {
		t.Errorf("Expected element type %d, got %d", original.ElementType, unmarshaled.ElementType)
	}
	if len(unmarshaled.Value) != len(original.Value) {
		t.Errorf("Expected length %d, got %d", len(original.Value), len(unmarshaled.Value))
	}

	for i := range original.Value {
		origInt := original.Value[i].(*TagInt)
		unmInt := unmarshaled.Value[i].(*TagInt)
		if unmInt.Value != origInt.Value {
			t.Errorf("At index %d: expected %d, got %d", i, origInt.Value, unmInt.Value)
		}
	}
}

func TestTagCompoundJSON(t *testing.T) {
	original := &TagCompound{
		baseTag: baseTag{tagType: BTagCompound, name: "testCompound"},
		Value: []NBTTag{
			&TagString{baseTag: baseTag{tagType: BTagString, name: "name"}, Value: "Test"},
			&TagInt{baseTag: baseTag{tagType: BTagInt, name: "age"}, Value: 25},
			&TagDouble{baseTag: baseTag{tagType: BTagDouble, name: "score"}, Value: 99.5},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagCompound: %v", err)
	}

	var unmarshaled TagCompound
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagCompound: %v", err)
	}

	if len(unmarshaled.Value) != len(original.Value) {
		t.Errorf("Expected length %d, got %d", len(original.Value), len(unmarshaled.Value))
	}

	// Check first element (string)
	origStr := original.Value[0].(*TagString)
	unmStr := unmarshaled.Value[0].(*TagString)
	if unmStr.Value != origStr.Value {
		t.Errorf("String value: expected %s, got %s", origStr.Value, unmStr.Value)
	}

	// Check second element (int)
	origInt := original.Value[1].(*TagInt)
	unmInt := unmarshaled.Value[1].(*TagInt)
	if unmInt.Value != origInt.Value {
		t.Errorf("Int value: expected %d, got %d", origInt.Value, unmInt.Value)
	}

	// Check third element (double)
	origDbl := original.Value[2].(*TagDouble)
	unmDbl := unmarshaled.Value[2].(*TagDouble)
	if unmDbl.Value != origDbl.Value {
		t.Errorf("Double value: expected %f, got %f", origDbl.Value, unmDbl.Value)
	}
}

func TestTagEndJSON(t *testing.T) {
	original := &TagEnd{
		baseTag: baseTag{tagType: BTagEnd, name: ""},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal TagEnd: %v", err)
	}

	var unmarshaled TagEnd
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal TagEnd: %v", err)
	}

	if unmarshaled.tagType != BTagEnd {
		t.Errorf("Expected type BTagEnd, got %d", unmarshaled.tagType)
	}
}

func TestNestedStructureJSON(t *testing.T) {
	// Create a complex nested structure
	original := &TagCompound{
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
					&TagDouble{baseTag: baseTag{tagType: BTagDouble, name: "pi"}, Value: 3.14159},
				},
			},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal nested structure: %v", err)
	}

	var unmarshaled TagCompound
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal nested structure: %v", err)
	}

	// Verify root level
	if len(unmarshaled.Value) != len(original.Value) {
		t.Errorf("Expected %d root elements, got %d", len(original.Value), len(unmarshaled.Value))
	}

	// Verify the nested list
	origList := original.Value[1].(*TagList)
	unmList := unmarshaled.Value[1].(*TagList)
	if len(unmList.Value) != len(origList.Value) {
		t.Errorf("Expected list length %d, got %d", len(origList.Value), len(unmList.Value))
	}

	// Verify the nested compound
	origNested := original.Value[2].(*TagCompound)
	unmNested := unmarshaled.Value[2].(*TagCompound)
	if len(unmNested.Value) != len(origNested.Value) {
		t.Errorf("Expected nested compound length %d, got %d", len(origNested.Value), len(unmNested.Value))
	}
}
