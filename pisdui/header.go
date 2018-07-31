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
	version   uint64
}

func (interpreter *Pisdui) ParseHeader() {
	headerByteSize := 24
	interpreter.File.Header.bytes = interpreter.FileContents[:headerByteSize]
	interpreter.File.Header.readSignature()
	interpreter.File.Header.readVersion()
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
	m := fh.bytes[versionStart:versionEnd]
	fh.version = binary.BigEndian.Uint64(m)
}
