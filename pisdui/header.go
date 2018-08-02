package pisdui

import (
	"encoding/binary"
	"fmt"
)

type headerIndex struct {
	start  int
	length int
}

type FileHeader struct {
	bytes     []byte
	signature string
	version   uint16
	reserved  []byte
}

func (interpreter *Pisdui) ParseHeader() {
	headerByteSize := 24
	interpreter.File.Header.bytes = interpreter.FileContents[:headerByteSize]
	interpreter.File.Header.readSignature()
	interpreter.File.Header.readVersion()
	interpreter.File.Header.readReserved()
	fmt.Println(interpreter.File.Header.signature)
}

func (fh *FileHeader) readSignature() {
	signatureStart := 0
	signatureEnd := 4
	fh.signature = string(fh.bytes[signatureStart:signatureEnd])
	if fh.signature != "8BPS" {
		panic("Invalid header signature. got " + fh.signature + " Expected 8BPS")
	}
}

func (fh *FileHeader) readVersion() {
	versionStart := 4
	versionEnd := 6
	fh.version = binary.BigEndian.Uint16(fh.bytes[versionStart:versionEnd])
	if fh.version != 1 {
		panic("Invalid file version.")
	}
}

func (fh *FileHeader) readReserved() {
	reservedStart := 7
	reservedEnd := 14
	fh.reserved = fh.bytes[reservedStart:reservedEnd]
	for i := 0; i < len(fh.reserved); i++ {
		if binary.BigEndian.Uint16(fh.reserved) != 0 {
			panic("reserved space not 0")
		}
	}
}
