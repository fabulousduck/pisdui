package util

import (
	"encoding/binary"
	"fmt"
	"os"
	"unicode/utf16"
)

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
