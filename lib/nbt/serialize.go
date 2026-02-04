package nbt

import (
	"encoding/binary"
	"math"
)

type SerializeError struct {
	message string
}

func (e SerializeError) Error() string {
	return e.message
}

func createSerializeError(message string) SerializeError {
	return SerializeError{message}
}

func SerializeTag(tag NBTTag, skipHeader bool) ([]byte, error) {
	switch tag.Type() {
	case BTagByte:
		{
			if skipHeader {
				return []byte{tag.(*TagByte).Value}, nil
			}
			return createPayload(tag, []byte{tag.(*TagByte).Value}), nil
		}
	case BTagShort:
		{
			payload := make([]byte, 2)
			binary.BigEndian.PutUint16(payload, uint16(tag.(*TagShort).Value))
			if skipHeader {
				return payload, nil
			}
			return createPayload(tag, payload), nil
		}
	case BTagInt:
		{
			payload := make([]byte, 4)
			binary.BigEndian.PutUint32(payload, uint32(tag.(*TagInt).Value))
			if skipHeader {
				return payload, nil
			}
			return createPayload(tag, payload), nil
		}
	case BTagLong:
		{
			payload := make([]byte, 8)
			binary.BigEndian.PutUint64(payload, uint64(tag.(*TagLong).Value))
			if skipHeader {
				return payload, nil
			}
			return createPayload(tag, payload), nil
		}
	case BTagFloat:
		{
			payload := make([]byte, 4)
			binary.BigEndian.PutUint32(payload, math.Float32bits(tag.(*TagFloat).Value))
			if skipHeader {
				return payload, nil
			}
			return createPayload(tag, payload), nil
		}
	case BTagDouble:
		{
			payload := make([]byte, 8)
			binary.BigEndian.PutUint64(payload, math.Float64bits(tag.(*TagDouble).Value))
			if skipHeader {
				return payload, nil
			}
			return createPayload(tag, payload), nil
		}
	case BTagByteArray:
		{
			byteArrayTag := tag.(*TagByteArray)
			arrayLength := len(byteArrayTag.Value)
			payload := make([]byte, 4+arrayLength)
			binary.BigEndian.PutUint32(payload, uint32(arrayLength))
			copy(payload[4:], byteArrayTag.Value)
			if skipHeader {
				return payload, nil
			}
			return createPayload(tag, payload), nil
		}
	case BTagString:
		{
			stringTag := tag.(*TagString)
			stringBytes := []byte(stringTag.Value)
			stringLength := len(stringBytes)
			payload := make([]byte, 2+stringLength)
			binary.BigEndian.PutUint16(payload, uint16(stringLength))
			copy(payload[2:], stringBytes)
			if skipHeader {
				return payload, nil
			}
			return createPayload(tag, payload), nil
		}
	case BTagEnd:
		{
			return createPayload(tag, []byte{}), nil
		}
	case BTagList:
		{
			listTag := tag.(*TagList)
			payload := []byte{byte(listTag.ElementType)}
			listLength := len(listTag.Value)
			lengthBytes := make([]byte, 4)
			binary.BigEndian.PutUint32(lengthBytes, uint32(listLength))
			payload = append(payload, lengthBytes...)
			for _, element := range listTag.Value {
				elementBytes, err := SerializeTag(element, true)
				if err != nil {
					return nil, err
				}
				payload = append(payload, elementBytes...)
			}
			return createPayload(tag, payload), nil
		}
	case BTagCompound:
		{
			compoundTag := tag.(*TagCompound)
			payload := []byte{}
			for _, childTag := range compoundTag.Value {
				childBytes, err := SerializeTag(childTag, false)
				if err != nil {
					return nil, err
				}
				payload = append(payload, childBytes...)
			}
			return createPayload(tag, payload), nil
		}
	}

	// For unsupported tag types
	return nil, createSerializeError("serialization for this tag type not implemented")
}

func createPayload(tag NBTTag, payload []byte) []byte {
	header := []byte{byte(tag.Type())}
	nameBytes := []byte(tag.Name())
	nameLength := len(nameBytes)
	binary.BigEndian.AppendUint16(header, uint16(nameLength))
	header = append(header, nameBytes...)
	header = append(header, payload...)
	return header
}
