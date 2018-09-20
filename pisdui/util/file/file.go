package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"unicode/utf16"
)

//TODO: make these return errors instead of panicing

/*ReadBytesNInt reads length bytes into a new buffer
and returns the result as a []byte
*/
func ReadBytesNInt(file *os.File, length uint32) []byte {

	byteBuffer := make([]byte, length)
	_, err := file.Read(byteBuffer)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return byteBuffer
}

/*ReadBytesLongLong reads 8 bytes into a new buffer
and returns the result as a uint64*/
func ReadBytesLongLong(file *os.File) uint64 {
	byteBuffer := make([]byte, 8)
	_, err := file.Read(byteBuffer)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return binary.BigEndian.Uint64(byteBuffer)
}

/*ReadBytesLong reads 4 bytes into a new buffer
and returns the result as a uint32*/
func ReadBytesLong(file *os.File) uint32 {

	byteBuffer := make([]byte, 4)
	_, err := file.Read(byteBuffer)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return binary.BigEndian.Uint32(byteBuffer)
}

/*ReadBytesShort reads 2 bytes into a new buffer
and returns the result as an uint16*/
func ReadBytesShort(file *os.File) uint16 {

	byteBuffer := make([]byte, 2)
	_, err := file.Read(byteBuffer)
	if err != nil {
		panic(err)
	}

	return binary.BigEndian.Uint16(byteBuffer)

}

/*ReadSingleByte reads 1 bytes into a new buffer
and returns the result as an uint16*/
func ReadSingleByte(file *os.File) int {

	byteBuffer := make([]byte, 1)
	_, err := file.Read(byteBuffer)
	if err != nil {
		panic(err)
	}

	return int(byteBuffer[0])
}

/*ReadBytesString reads length number of bytes into a new buffer
and returns the result as a string*/
func ReadBytesString(file *os.File, length int) string {
	byteBuffer := make([]byte, length)
	_, err := file.Read(byteBuffer)
	if err != nil {
		panic(err)
	}
	return string(byteBuffer)
}

/*ReadIntoArray16 takes a []byte and creates a new slice containing the values in uint16 form*/
func ReadIntoArray16(file *os.File, length uint32) []uint16 {

	newBufferLength := length / 2
	newBuffer := make([]uint16, newBufferLength)
	var i uint32
	for i = 0; i < newBufferLength; i++ {
		newBuffer = append(newBuffer, ReadBytesShort(file))
	}

	return newBuffer
}

/*ReadIntoArray8 takes a []byte and creates a new slice containing the values in uint8 form*/
func ReadIntoArray8(file *os.File, length uint32) []uint8 {

	newBufferLength := length / 2
	newBuffer := make([]uint8, newBufferLength)
	var i uint32
	for i = 0; i < newBufferLength; i++ {
		newBuffer = append(newBuffer, uint8(ReadSingleByte(file)))
	}

	return newBuffer
}

//ReadRawBytes returns a buffer containing length bytes at the current file pointer index
func ReadRawBytes(file *os.File, length int) []byte {
	buffer := make([]byte, length)
	_, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	return buffer
}

/*ParseUnicodeString parses a unicode string from the
photoshop file into a string*/
func ParseUnicodeString(file *os.File) string {
	length := ReadBytesLong(file)
	if length == 0 {
		return ""
	}

	stringBuffer := make([]uint16, length)

	for i := range stringBuffer {
		stringBuffer[i] = ReadBytesShort(file)
	}

	return string(utf16.Decode(stringBuffer))
}

/*ParsePascalString parses a pascal string directly from file*/
func ParsePascalString(file *os.File) string {
	b := ReadSingleByte(file)
	if b == 0 {
		ReadSingleByte(file)
		return ""
	}

	s := ReadBytesString(file, b)

	if b%2 != 0 {
		ReadSingleByte(file)
	}
	return s
}

/*ReadDouble reads a 64bit int*/
func ReadDouble(file *os.File) (float64, error) {
	buffer := make([]byte, 8)
	_, err := file.Read(buffer)
	if err != nil {
		return 0, err
	}
	reader := bytes.NewReader(buffer)
	var res float64
	binary.Read(reader, binary.BigEndian, &res)
	return res, nil
}

/*ReadFloat reads 4bytes into a float*/
func ReadFloat(file *os.File) float32 {
	buffer := make([]byte, 4)
	_, err := file.Read(buffer)
	if err != nil {
		return 0
	}
	reader := bytes.NewReader(buffer)
	var res float32
	binary.Read(reader, binary.BigEndian, &res)
	return res
}
