package nbt

import (
	"encoding/json"
	"fmt"
)

// JSONTag is a helper structure for JSON serialization
type JSONTag struct {
	Type  string          `json:"type"`
	Name  string          `json:"name,omitempty"`
	Value json.RawMessage `json:"value"`
}

// MarshalJSON implements json.Marshaler for TagByte
func (t *TagByte) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "byte",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagByte
func (t *TagByte) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value byte   `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "byte" {
		return fmt.Errorf("expected type 'byte', got '%s'", temp.Type)
	}
	t.tagType = BTagByte
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagShort
func (t *TagShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "short",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagShort
func (t *TagShort) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value int16  `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "short" {
		return fmt.Errorf("expected type 'short', got '%s'", temp.Type)
	}
	t.tagType = BTagShort
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagInt
func (t *TagInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "int",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagInt
func (t *TagInt) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value int32  `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "int" {
		return fmt.Errorf("expected type 'int', got '%s'", temp.Type)
	}
	t.tagType = BTagInt
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagLong
func (t *TagLong) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "long",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagLong
func (t *TagLong) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value int64  `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "long" {
		return fmt.Errorf("expected type 'long', got '%s'", temp.Type)
	}
	t.tagType = BTagLong
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagFloat
func (t *TagFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "float",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagFloat
func (t *TagFloat) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string  `json:"type"`
		Name  string  `json:"name"`
		Value float32 `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "float" {
		return fmt.Errorf("expected type 'float', got '%s'", temp.Type)
	}
	t.tagType = BTagFloat
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagDouble
func (t *TagDouble) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "double",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagDouble
func (t *TagDouble) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string  `json:"type"`
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "double" {
		return fmt.Errorf("expected type 'double', got '%s'", temp.Type)
	}
	t.tagType = BTagDouble
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagString
func (t *TagString) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "string",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagString
func (t *TagString) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "string" {
		return fmt.Errorf("expected type 'string', got '%s'", temp.Type)
	}
	t.tagType = BTagString
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagByteArray
func (t *TagByteArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "byteArray",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagByteArray
func (t *TagByteArray) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value []byte `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "byteArray" {
		return fmt.Errorf("expected type 'byteArray', got '%s'", temp.Type)
	}
	t.tagType = BTagByteArray
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagIntArray
func (t *TagIntArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "intArray",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagIntArray
func (t *TagIntArray) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string  `json:"type"`
		Name  string  `json:"name"`
		Value []int32 `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "intArray" {
		return fmt.Errorf("expected type 'intArray', got '%s'", temp.Type)
	}
	t.tagType = BTagIntArray
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagLongArray
func (t *TagLongArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "longArray",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagLongArray
func (t *TagLongArray) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string  `json:"type"`
		Name  string  `json:"name"`
		Value []int64 `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "longArray" {
		return fmt.Errorf("expected type 'longArray', got '%s'", temp.Type)
	}
	t.tagType = BTagLongArray
	t.name = temp.Name
	t.Value = temp.Value
	return nil
}

// MarshalJSON implements json.Marshaler for TagList
func (t *TagList) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":        "list",
		"name":        t.name,
		"elementType": tagTypeToString(t.ElementType),
		"value":       t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagList
func (t *TagList) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type        string            `json:"type"`
		Name        string            `json:"name"`
		ElementType string            `json:"elementType"`
		Value       []json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "list" {
		return fmt.Errorf("expected type 'list', got '%s'", temp.Type)
	}

	t.tagType = BTagList
	t.name = temp.Name
	t.ElementType = stringToTagType(temp.ElementType)

	t.Value = make([]NBTTag, len(temp.Value))
	for i, rawTag := range temp.Value {
		tag, err := unmarshalNBTTag(rawTag)
		if err != nil {
			return fmt.Errorf("error unmarshaling list element %d: %w", i, err)
		}
		t.Value[i] = tag
	}

	return nil
}

// MarshalJSON implements json.Marshaler for TagCompound
func (t *TagCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":  "compound",
		"name":  t.name,
		"value": t.Value,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagCompound
func (t *TagCompound) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string            `json:"type"`
		Name  string            `json:"name"`
		Value []json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "compound" {
		return fmt.Errorf("expected type 'compound', got '%s'", temp.Type)
	}

	t.tagType = BTagCompound
	t.name = temp.Name

	t.Value = make([]NBTTag, len(temp.Value))
	for i, rawTag := range temp.Value {
		tag, err := unmarshalNBTTag(rawTag)
		if err != nil {
			return fmt.Errorf("error unmarshaling compound element %d: %w", i, err)
		}
		t.Value[i] = tag
	}

	return nil
}

// MarshalJSON implements json.Marshaler for TagEnd
func (t *TagEnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type": "end",
		"name": t.name,
	})
}

// UnmarshalJSON implements json.Unmarshaler for TagEnd
func (t *TagEnd) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "end" {
		return fmt.Errorf("expected type 'end', got '%s'", temp.Type)
	}
	t.tagType = BTagEnd
	t.name = temp.Name
	return nil
}

// Helper function to convert tagTypeByte to string
func tagTypeToString(tagType tagTypeByte) string {
	switch tagType {
	case BTagEnd:
		return "end"
	case BTagByte:
		return "byte"
	case BTagShort:
		return "short"
	case BTagInt:
		return "int"
	case BTagLong:
		return "long"
	case BTagFloat:
		return "float"
	case BTagDouble:
		return "double"
	case BTagByteArray:
		return "byteArray"
	case BTagString:
		return "string"
	case BTagList:
		return "list"
	case BTagCompound:
		return "compound"
	case BTagIntArray:
		return "intArray"
	case BTagLongArray:
		return "longArray"
	default:
		return "unknown"
	}
}

// Helper function to convert string to tagTypeByte
func stringToTagType(typeStr string) tagTypeByte {
	switch typeStr {
	case "end":
		return BTagEnd
	case "byte":
		return BTagByte
	case "short":
		return BTagShort
	case "int":
		return BTagInt
	case "long":
		return BTagLong
	case "float":
		return BTagFloat
	case "double":
		return BTagDouble
	case "byteArray":
		return BTagByteArray
	case "string":
		return BTagString
	case "list":
		return BTagList
	case "compound":
		return BTagCompound
	case "intArray":
		return BTagIntArray
	case "longArray":
		return BTagLongArray
	default:
		return BTagEnd
	}
}

// unmarshalNBTTag unmarshals a JSON raw message into the appropriate NBTTag type
func unmarshalNBTTag(data []byte) (NBTTag, error) {
	var typeCheck struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(data, &typeCheck); err != nil {
		return nil, err
	}

	switch typeCheck.Type {
	case "byte":
		var tag TagByte
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "short":
		var tag TagShort
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "int":
		var tag TagInt
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "long":
		var tag TagLong
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "float":
		var tag TagFloat
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "double":
		var tag TagDouble
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "string":
		var tag TagString
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "byteArray":
		var tag TagByteArray
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "intArray":
		var tag TagIntArray
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "longArray":
		var tag TagLongArray
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "list":
		var tag TagList
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "compound":
		var tag TagCompound
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	case "end":
		var tag TagEnd
		if err := json.Unmarshal(data, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	default:
		return nil, fmt.Errorf("unknown tag type: %s", typeCheck.Type)
	}
}
