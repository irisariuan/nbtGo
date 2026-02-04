package nbt

type tagTypeByte byte

const (
	// no payload
	BTagEnd tagTypeByte = 0

	// Contains any number of tags, delimited by TAG_End.
	// Each tag consisting of 1 byte / 8 bits tag ID,
	// followed by 2 bytes / 16 bits, unsigned, big-endian for size, then an UTF-8 formatted string containing the tag name.
	// Lastly, the payload data.
	BTagCompound tagTypeByte = 10

	// 1 byte / 8 bits, signed
	BTagByte tagTypeByte = 1

	// 2 bytes / 16 bits, signed, big endian
	BTagShort tagTypeByte = 2

	// 4 bytes / 32 bits, signed, big endian
	BTagInt tagTypeByte = 3

	// 8 bytes / 64 bits, signed, big endian
	BTagLong tagTypeByte = 4

	// 4 bytes / 32 bits, signed, big endian, IEEE 754-2008, binary32
	BTagFloat tagTypeByte = 5

	// 8 bytes / 64 bits, signed, big endian, IEEE 754-2008, binary64
	BTagDouble tagTypeByte = 6

	// 4 bytes / 32 bits, signed, big endian for size, then the bytes of length size
	BTagByteArray tagTypeByte = 7

	// 2 bytes / 16 bits, unsigned, big endian for size,
	// then the bytes of length size as UTF-8 formatted character data.
	// not null-terminated
	BTagString tagTypeByte = 8

	// Lists are homogeneous arrays (for same type) of nameless tags
	// 1 byte / 8 bits for the tag ID of the list's contents,
	// then 4 bytes / 32 bits, big-endian for size.
	// Followed by length size number of items of tag id
	BTagList tagTypeByte = 9

	// 4 bytes / 32 bits, signed, big-endian for size, then size number of TAG_Int payloads.
	BTagIntArray tagTypeByte = 11

	// 4 bytes / 32 bits, signed, big-endian for size, then size number of TAG_Long payloads.
	BTagLongArray tagTypeByte = 12
)

var TagPayloadLength map[tagTypeByte]int = map[tagTypeByte]int{
	BTagCompound:  -2, // Variable length
	BTagEnd:       0,
	BTagByte:      1,
	BTagShort:     2,
	BTagInt:       4,
	BTagLong:      8,
	BTagFloat:     4,
	BTagDouble:    8,
	BTagByteArray: -1, // Variable length
	BTagString:    -1, // Variable length
	BTagList:      -1, // Variable length
	BTagIntArray:  -1, // Variable length
	BTagLongArray: -1, // Variable length
}

// A tag is an individual part of the data tree. The first byte in a tag is the tag type (ID),
//
// followed by a two byte big-endian unsigned integer (ushort) for the length of the name,
//
// then the name as a string in UTF-8 format
//
// (Note TAG_End is not named and does not contain the extra 2 bytes; the name is assumed to be empty).
//
// Finally, depending on the type of the tag, the bytes that follow are part of that tag's payload.
//
// NBTTag is the interface all tag types implement
type NBTTag interface {
	Type() tagTypeByte
	Name() string
	DataLength() int
	ZIndex() int
}

// Base fields common to all tags
type baseTag struct {
	tagType tagTypeByte
	name    string
	zIndex  int
}

func (t *baseTag) Type() tagTypeByte { return t.tagType }
func (t *baseTag) Name() string      { return t.name }
func (t *baseTag) ZIndex() int       { return t.zIndex }

// TagByte represents a single signed byte
type TagByte struct {
	baseTag
	Value byte
}

func (t *TagByte) DataLength() int { return 1 }

// TagShort represents a signed 16-bit integer
type TagShort struct {
	baseTag
	Value int16
}

func (t *TagShort) DataLength() int { return 2 }

// TagInt represents a signed 32-bit integer
type TagInt struct {
	baseTag
	Value int32
}

func (t *TagInt) DataLength() int { return 4 }

// TagLong represents a signed 64-bit integer
type TagLong struct {
	baseTag
	Value int64
}

func (t *TagLong) DataLength() int { return 8 }

// TagFloat represents a 32-bit float
type TagFloat struct {
	baseTag
	Value float32
}

func (t *TagFloat) DataLength() int { return 4 }

// TagDouble represents a 64-bit float
type TagDouble struct {
	baseTag
	Value float64
}

func (t *TagDouble) DataLength() int { return 8 }

// TagString represents a UTF-8 string
type TagString struct {
	baseTag
	Value string
}

func (t *TagString) DataLength() int { return 2 + len(t.Value) } // 2 bytes for length + string bytes

// TagByteArray represents an array of bytes
type TagByteArray struct {
	baseTag
	Value []byte
}

func (t *TagByteArray) DataLength() int { return 4 + len(t.Value) } // 4 bytes for length + byte array

// TagIntArray represents an array of int32
type TagIntArray struct {
	baseTag
	Value []int32
}

func (t *TagIntArray) DataLength() int { return 4 + len(t.Value)*4 } // 4 bytes for length + int32 array

// TagLongArray represents an array of int64
type TagLongArray struct {
	baseTag
	Value []int64
}

func (t *TagLongArray) DataLength() int { return 4 + len(t.Value)*8 } // 4 bytes for length + int64 array

// TagList represents a list of tags (all same type)
type TagList struct {
	baseTag
	ElementType tagTypeByte
	Value       []NBTTag
}

func (t *TagList) DataLength() int {
	totalLength := 1 + 4 // 1 byte for list element type + 4 bytes for length
	for _, item := range t.Value {
		totalLength += item.DataLength()
	}
	return totalLength
}

// TagCompound represents a collection of named tags
//
// Each tag is stored sequentially, ending with a TagEnd
type TagCompound struct {
	baseTag
	Value []NBTTag
}

func (t *TagCompound) DataLength() int {
	totalLength := 0
	for _, item := range t.Value {
		totalLength += GetTagFullSize(item)
	}
	return totalLength
}

// TagEnd marks the end of a compound
type TagEnd struct {
	baseTag
}

func (t *TagEnd) DataLength() int { return 0 }
