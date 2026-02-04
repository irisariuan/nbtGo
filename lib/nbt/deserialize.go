package nbt

import (
	"encoding/hex"
	"fmt"
	"goNbt/lib"
)

type TagParseError interface {
	isFatal() bool
	Error() string
}

type tagParseValueError struct {
	message string
}
type tagParseArrayError struct {
	message string
}

func (e tagParseValueError) Error() string {
	return e.message
}
func (e tagParseValueError) isFatal() bool {
	return true
}
func (e tagParseArrayError) Error() string {
	return e.message
}
func (e tagParseArrayError) isFatal() bool {
	return false
}

func newParseValueError(message string) TagParseError {
	return tagParseValueError{message}
}
func newParseArrayError(message string) TagParseError {
	return tagParseArrayError{message}
}

func ParseNBT(data []byte, isBedrock bool) (NBTTag, TagParseError) {
	tag, remaining, err := separateSingleTag(data, 0, !isBedrock)
	if err != nil {
		return nil, err
	}
	if len(remaining) > 0 {
		fmt.Printf("Warning: %d bytes of extra data after parsing NBT tag (data: %s)\n", len(remaining), hex.EncodeToString(remaining))
		return nil, newParseArrayError("extra data after parsing NBT tag")
	}
	if tag.Type() != BTagCompound && tag.Type() != BTagList {
		return nil, newParseValueError("root tag is not a Compound or List")
	}
	return tag, nil
}

// separateSingleTag parses a single NBT tag (with known length of data) from the given byte slice.
//
// It returns the parsed Tag, any remaining unparsed bytes, and an error if parsing fails.
func separateSingleTag(data []byte, zIndex int, bigEndian bool) (NBTTag, []byte, TagParseError) {
	tagType := tagTypeByte(data[0])
	if tagType == BTagEnd {
		tag := baseTag{tagType, "", zIndex}
		return &TagEnd{baseTag: tag}, data[1:], nil
	}
	nameLength, err := lib.BytesToUInt16(data[1:3], bigEndian)
	if err != nil {
		return nil, nil, newParseValueError("failed to parse name length")
	}
	nameLengthInt := int(nameLength)
	if len(data) < 3+nameLengthInt {
		fmt.Printf("Error: Type: %s, Name Length: (full 0-2 bytes %s in hex) or %d\n", hex.EncodeToString([]byte{byte(tagType)}), hex.EncodeToString(data[0:3]), nameLengthInt)
		return nil, nil, newParseValueError("data too short for tag name")
	}
	name := string(data[3 : 3+nameLengthInt])
	tag := baseTag{tagType, name, zIndex}
	payload := data[3+nameLengthInt:]
	parsed, err := parsePayload(tag, payload, bigEndian)
	if err != nil {
		return nil, nil, err.(TagParseError)
	}
	remaining := data[GetTagFullSize(parsed):]
	return parsed, remaining, nil
}

// parsePayload parses the payload of a tag based on its type.
//
// payload has a unknown length, and may contains other tags data
// which would be ignored while parsing
func parsePayload(tag baseTag, payload []byte, bigEndian bool) (NBTTag, TagParseError) {
	// check length
	minPayloadLength, ok := TagPayloadLength[tag.Type()]
	if !ok {
		return nil, newParseValueError("unknown tag type")
	}
	if minPayloadLength > 0 && len(payload) < minPayloadLength {
		return nil, newParseValueError("payload too short to parse")
	}

	switch tag.Type() {
	case BTagEnd:
		{
			return &TagEnd{baseTag: tag}, nil
		}
	case BTagByte:
		{
			return &TagByte{baseTag: tag, Value: payload[0]}, nil
		}
	case BTagShort:
		{
			value, error := lib.BytesToInt16(payload[0:2], bigEndian)
			if error != nil {
				return nil, newParseValueError("failed to parse short payload")
			}
			return &TagShort{baseTag: tag, Value: value}, nil
		}
	case BTagInt:
		{
			value, error := lib.BytesToInt32(payload[0:4], bigEndian)
			if error != nil {
				return nil, newParseValueError("failed to parse int payload")
			}
			return &TagInt{baseTag: tag, Value: value}, nil
		}
	case BTagLong:
		{
			value, error := lib.BytesToInt64(payload[0:8], bigEndian)
			if error != nil {
				return nil, newParseValueError("failed to parse long payload")
			}
			return &TagLong{baseTag: tag, Value: value}, nil
		}
	case BTagFloat:
		{
			floatValue, err := lib.BytesFloat32(payload[0:4], bigEndian)
			if err != nil {
				return nil, newParseValueError("failed to parse float payload")
			}
			return &TagFloat{baseTag: tag, Value: floatValue}, nil
		}
	case BTagDouble:
		{
			floatValue, err := lib.BytesFloat64(payload[0:8], bigEndian)
			if err != nil {
				return nil, newParseValueError("failed to parse double payload")
			}
			return &TagDouble{baseTag: tag, Value: floatValue}, nil
		}
	case BTagByteArray:
		{
			arrayLengthUint, err := lib.BytesToUInt32(payload[0:4], bigEndian)
			if err != nil {
				return nil, newParseValueError("failed to parse byte array length")
			}
			arrayLength := int(arrayLengthUint)
			if len(payload) < 4+arrayLength {
				return nil, newParseValueError("payload too short for byte array")
			}
			arrayData := payload[4 : 4+arrayLength]
			return &TagByteArray{baseTag: tag, Value: arrayData}, nil
		}
	case BTagString:
		{
			stringLengthUint, err := lib.BytesToUInt16(payload[0:2], bigEndian)
			if err != nil {
				return nil, newParseValueError("failed to parse string length")
			}
			stringLength := int(stringLengthUint)
			if len(payload) < 2+stringLength {
				return nil, newParseValueError("payload too short for string")
			}
			stringData := string(payload[2 : 2+stringLength])
			return &TagString{baseTag: tag, Value: stringData}, nil
		}
	case BTagList:
		{
			listType := tagTypeByte(payload[0])
			listLengthUint, err := lib.BytesToUInt32(payload[1:5], bigEndian)
			if err != nil {
				return nil, newParseValueError("failed to parse list length")
			}
			listLength := int(listLengthUint)
			items := make([]NBTTag, 0) // unknown size, known length
			offset := 5
			for range listLength {
				itemTag := baseTag{listType, "", tag.zIndex + 1}

				item, err := parsePayload(itemTag, payload[offset:], bigEndian)
				if err != nil {
					return nil, newParseArrayError("error parsing list item")
				}
				items = append(items, item)
				offset += item.DataLength() // no type ID and name for list items
			}
			return &TagList{ElementType: listType, baseTag: tag, Value: items}, nil
		}
	case BTagCompound:
		{
			recvTag, remaining, err := separateSingleTag(payload, tag.zIndex+1, bigEndian)
			if err != nil {
				return nil, newParseArrayError("error parsing first compound tag")
			}
			arr := []NBTTag{recvTag}
			for recvTag.Type() != BTagEnd {
				recvTag, remaining, err = separateSingleTag(remaining, tag.zIndex+1, bigEndian)
				if err != nil {
					fmt.Println("Error:", err)
					return nil, newParseArrayError("error in middle while parsing compound tag")
				}
				arr = append(arr, recvTag)
			}
			if recvTag.Type() != BTagEnd && len(remaining) == 0 {
				return nil, newParseArrayError("compound tag ended without TAG_End")
			}
			return &TagCompound{baseTag: tag, Value: arr}, nil
		}
	case BTagIntArray:
		{
			arrSize, err := lib.BytesToInt32(payload[0:4], bigEndian)
			if err != nil {
				return nil, newParseValueError("error parsing array size")
			}
			offset := 4
			arr := make([]int32, arrSize)
			for i := range arrSize {
				// parse per 4 bytes
				if len(payload[offset:]) < 4 {
					return nil, newParseArrayError("error parsing int array tag")
				}
				intValue, err := lib.BytesToInt32(payload[offset:offset+4], bigEndian)
				if err != nil {
					return nil, newParseArrayError("error parsing int array tag value")
				}
				arr[i] = intValue
				offset += 4
			}
			return &TagIntArray{baseTag: tag, Value: arr}, nil
		}
	case BTagLongArray:
		{
			arrSize, err := lib.BytesToInt32(payload[0:4], bigEndian)
			if err != nil {
				return nil, newParseValueError("error parsing array size")
			}
			offset := 4
			arr := make([]int64, arrSize)
			for i := range arrSize {
				// parse per 8 bytes
				if len(payload[offset:]) < 8 {
					return nil, newParseArrayError("error parsing long array tag")
				}
				longValue, err := lib.BytesToInt64(payload[offset:offset+8], bigEndian)
				if err != nil {
					return nil, newParseArrayError("error parsing long array tag value")
				}
				arr[i] = longValue
				offset += 8
			}
			return &TagLongArray{baseTag: tag, Value: arr}, nil
		}
	default:
		{

			return nil, newParseValueError("unknown tag type")
		}
	}
}

func GetTagFullSize(tag NBTTag) int {
	if tag.Type() == BTagEnd {
		return 1
	}
	return 1 + 2 + len(tag.Name()) + tag.DataLength()
}
