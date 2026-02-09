package lib

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"compress/zlib"
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

func ZipToGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	n, err := zw.Write(data)
	if err != nil {
		return nil, err
	}
	if n < len(data) {
		return nil, newByteError("failed to write all data to gzip writer")
	}
	err = zw.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ZipToZlib(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	n, err := zw.Write(data)
	if err != nil {
		return nil, err
	}
	if n < len(data) {
		return nil, newByteError("failed to write all data to zlib writer")
	}
	err = zw.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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
	magicBytes := make([]byte, 2)
	n, err := reader.Read(magicBytes)
	if err != nil {
		return nil, err
	}
	if n < 2 {
		return nil, io.ErrUnexpectedEOF
	}
	if magicBytes[0] != 0x1f || magicBytes[1] != 0x8b {
		// check if it's a zlib format
		if magicBytes[0] == 0x78 && (magicBytes[1] == 0x01 || magicBytes[1] == 0x5e || magicBytes[1] == 0x9c || magicBytes[1] == 0xda) {
			combinedReader := Reset(bufio.NewReader(reader), magicBytes)
			zlibReader, err := zlib.NewReader(combinedReader)
			if err != nil {
				return nil, err
			}
			defer zlibReader.Close()
			return io.ReadAll(zlibReader)
		}

		// Not compressed, return the original data
		combinedReader := Reset(bufio.NewReader(reader), magicBytes)
		return io.ReadAll(combinedReader)
	}
	// It's gzip format
	reader = Reset(bufio.NewReader(reader), magicBytes)
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	return io.ReadAll(gzipReader)
}

// UInt16ToBytes converts a uint16 to a byte slice in big-endian or little-endian format
func UInt16ToBytes(val uint16, bigEndian bool) []byte {
	bytes := make([]byte, 2)
	if bigEndian {
		binary.BigEndian.PutUint16(bytes, val)
	} else {
		binary.LittleEndian.PutUint16(bytes, val)
	}
	return bytes
}

// Int16ToBytes converts an int16 to a byte slice in big-endian or little-endian format
func Int16ToBytes(val int16, bigEndian bool) []byte {
	return UInt16ToBytes(uint16(val), bigEndian)
}

// Int32ToBytes converts an int32 to a byte slice in big-endian or little-endian format
func Int32ToBytes(val int32, bigEndian bool) []byte {
	bytes := make([]byte, 4)
	if bigEndian {
		binary.BigEndian.PutUint32(bytes, uint32(val))
	} else {
		binary.LittleEndian.PutUint32(bytes, uint32(val))
	}
	return bytes
}

// Int64ToBytes converts an int64 to a byte slice in big-endian or little-endian format
func Int64ToBytes(val int64, bigEndian bool) []byte {
	bytes := make([]byte, 8)
	if bigEndian {
		binary.BigEndian.PutUint64(bytes, uint64(val))
	} else {
		binary.LittleEndian.PutUint64(bytes, uint64(val))
	}
	return bytes
}

// Float32ToBytes converts a float32 to a byte slice in big-endian or little-endian format
func Float32ToBytes(val float32, bigEndian bool) []byte {
	bits := math.Float32bits(val)
	bytes := make([]byte, 4)
	if bigEndian {
		binary.BigEndian.PutUint32(bytes, bits)
	} else {
		binary.LittleEndian.PutUint32(bytes, bits)
	}
	return bytes
}

// Float64ToBytes converts a float64 to a byte slice in big-endian or little-endian format
func Float64ToBytes(val float64, bigEndian bool) []byte {
	bits := math.Float64bits(val)
	bytes := make([]byte, 8)
	if bigEndian {
		binary.BigEndian.PutUint64(bytes, bits)
	} else {
		binary.LittleEndian.PutUint64(bytes, bits)
	}
	return bytes
}
