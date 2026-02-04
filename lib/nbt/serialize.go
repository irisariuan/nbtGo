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
	switch t := tag.(type) {
	case *TagByte:
		if skipHeader {
			return []byte{t.Value}, nil
		}
		return createPayload(t, []byte{t.Value}), nil
	case *TagShort:
		payload := make([]byte, 2)
		binary.BigEndian.PutUint16(payload, uint16(t.Value))
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagInt:
		payload := make([]byte, 4)
		binary.BigEndian.PutUint32(payload, uint32(t.Value))
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagLong:
		payload := make([]byte, 8)
		binary.BigEndian.PutUint64(payload, uint64(t.Value))
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagFloat:
		payload := make([]byte, 4)
		binary.BigEndian.PutUint32(payload, math.Float32bits(t.Value))
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagDouble:
		payload := make([]byte, 8)
		binary.BigEndian.PutUint64(payload, math.Float64bits(t.Value))
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagByteArray:
		arrayLength := len(t.Value)
		payload := make([]byte, 4+arrayLength)
		binary.BigEndian.PutUint32(payload, uint32(arrayLength))
		copy(payload[4:], t.Value)
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagString:
		stringBytes := []byte(t.Value)
		stringLength := len(stringBytes)
		payload := make([]byte, 2+stringLength)
		binary.BigEndian.PutUint16(payload, uint16(stringLength))
		copy(payload[2:], stringBytes)
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagIntArray:
		arrayLength := len(t.Value)
		payload := make([]byte, 4+arrayLength*4)
		binary.BigEndian.PutUint32(payload, uint32(arrayLength))
		offset := 4
		for _, val := range t.Value {
			binary.BigEndian.PutUint32(payload[offset:], uint32(val))
			offset += 4
		}
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagLongArray:
		arrayLength := len(t.Value)
		payload := make([]byte, 4+arrayLength*8)
		binary.BigEndian.PutUint32(payload, uint32(arrayLength))
		offset := 4
		for _, val := range t.Value {
			binary.BigEndian.PutUint64(payload[offset:], uint64(val))
			offset += 8
		}
		if skipHeader {
			return payload, nil
		}
		return createPayload(t, payload), nil
	case *TagEnd:
		return createPayload(t, nil), nil
	case *TagList:
		payload := []byte{byte(t.ElementType)}
		listLength := len(t.Value)
		lengthBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(lengthBytes, uint32(listLength))
		payload = append(payload, lengthBytes...)
		for _, element := range t.Value {
			elementBytes, err := SerializeTag(element, true)
			if err != nil {
				return nil, err
			}
			payload = append(payload, elementBytes...)
		}
		return createPayload(t, payload), nil
	case *TagCompound:
		payload := []byte{}
		for _, childTag := range t.Value {
			childBytes, err := SerializeTag(childTag, false)
			if err != nil {
				return nil, err
			}
			payload = append(payload, childBytes...)
		}
		return createPayload(t, payload), nil
	default:
		// For unsupported tag types
		return nil, createSerializeError("serialization for this tag type not implemented")
	}
}

func createPayload(tag NBTTag, payload []byte) []byte {
	header := []byte{byte(tag.Type())}
	if tag.Type() == BTagEnd {
		// TAG_End has no name or payload
		return header
	}
	nameBytes := []byte(tag.Name())
	nameLength := len(nameBytes)
	header = binary.BigEndian.AppendUint16(header, uint16(nameLength))
	if nameLength > 0 {
		header = append(header, nameBytes...)
	}
	if len(payload) > 0 {
		header = append(header, payload...)
	}
	return header
}
