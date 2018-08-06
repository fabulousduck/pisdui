package pisdui

import (
	"encoding/binary"
	"fmt"
	"os"
)

/*ReadBytesNInt reads length bytes into a new buffer
and returns the result as a []byte
*/
func ReadBytesNInt(file *os.File, length uint64) []byte {
	byteBuffer := make([]byte, length)
	_, err := file.Read(byteBuffer)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return byteBuffer
}

/*ReadBytesLong reads 4 bytes into a new buffer
and returns the result as an uint32*/
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

/*ReadBytesString reads length number of bytes into a new buffer
and returns the result as a string*/
func ReadBytesString(file *os.File, length int) string {
	byteBuffer := make([]byte, length)
	_, err := file.Read(byteBuffer)
	fmt.Println(byteBuffer)
	if err != nil {
		panic(err)
	}
	return string(byteBuffer)
}
