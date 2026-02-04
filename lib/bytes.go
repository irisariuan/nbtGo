package lib

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"io"
	"math"
)

type byteError struct {
	message string
}

func (e byteError) Error() string {
	return e.message
}
func newByteError(message string) byteError {
	return byteError{message: message}
}

func bytesReader(b []byte) io.Reader {
	return &byteReader{b: b}
}

type byteReader struct {
	b []byte
	i int
}

func (r *byteReader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.b) {
		return 0, io.EOF
	}
	n = copy(p, r.b[r.i:])
	r.i += n
	return n, nil
}

func Reset(reader *bufio.Reader, bytes []byte) io.Reader {
	return io.MultiReader(bytesReader(bytes), reader)
}

func BytesFloat32(bytes []byte, bigEndian bool) (float32, error) {
	if len(bytes) < 4 {
		return 0, newByteError("not enough bytes to convert to float32")
	}
	if bigEndian {
		bits := binary.BigEndian.Uint32(bytes)
		float := math.Float32frombits(bits)
		return float, nil
	}
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float, nil
}

func BytesFloat64(bytes []byte, bigEndian bool) (float64, error) {
	if len(bytes) < 8 {
		return 0, newByteError("not enough bytes to convert to float64")
	}
	if bigEndian {
		bits := binary.BigEndian.Uint64(bytes)
		float := math.Float64frombits(bits)
		return float, nil
	}
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float, nil
}

func BytesToInt16(bytes []byte, bigEndian bool) (int16, error) {
	if len(bytes) < 2 {
		return 0, newByteError("not enough bytes to convert to int16")
	}
	if bigEndian {
		value := int16(binary.BigEndian.Uint16(bytes))
		return value, nil
	}
	value := int16(binary.LittleEndian.Uint16(bytes))
	return value, nil
}

func BytesToInt32(bytes []byte, bigEndian bool) (int32, error) {
	if len(bytes) < 4 {
		return 0, newByteError("not enough bytes to convert to int32")
	}
	if bigEndian {
		value := int32(binary.BigEndian.Uint32(bytes))
		return value, nil
	}
	value := int32(binary.LittleEndian.Uint32(bytes))
	return value, nil
}

func BytesToInt64(bytes []byte, bigEndian bool) (int64, error) {
	if len(bytes) < 8 {
		return 0, newByteError("not enough bytes to convert to int64")
	}
	if bigEndian {
		value := int64(binary.BigEndian.Uint64(bytes))
		return value, nil
	}
	value := int64(binary.LittleEndian.Uint64(bytes))
	return value, nil
}

func BytesToUInt16(bytes []byte, bigEndian bool) (uint16, error) {
	if len(bytes) < 2 {
		return 0, newByteError("not enough bytes to convert to uint16")
	}
	if bigEndian {
		return binary.BigEndian.Uint16(bytes), nil
	}
	return binary.LittleEndian.Uint16(bytes), nil
}

func BytesToUInt32(bytes []byte, bigEndian bool) (uint32, error) {
	if len(bytes) < 4 {
		return 0, newByteError("not enough bytes to convert to uint32")
	}
	if bigEndian {
		return binary.BigEndian.Uint32(bytes), nil
	}
	return binary.LittleEndian.Uint32(bytes), nil
}

func BytesToUInt64(bytes []byte, bigEndian bool) (uint64, error) {
	if len(bytes) < 8 {
		return 0, newByteError("not enough bytes to convert to uint64")
	}
	if bigEndian {
		return binary.BigEndian.Uint64(bytes), nil
	}
	return binary.LittleEndian.Uint64(bytes), nil
}

func UnzipReader(reader io.Reader) ([]byte, error) {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	return io.ReadAll(gzipReader)
}
