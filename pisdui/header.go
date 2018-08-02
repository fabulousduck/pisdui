package pisdui

import (
	"encoding/binary"
)

type FileHeader struct {
	bytes     []byte
	signature string
	version   uint16
	reserved  []byte
	channels  uint16
	height    uint16
	width     uint16
	depth     uint16
	colorMode uint16
}

func (interpreter *Pisdui) ParseHeader() {
	headerByteSize := 24
	interpreter.File.Header.bytes = interpreter.FileContents[:headerByteSize]
	interpreter.File.Header.readSignature()
	interpreter.File.Header.readVersion()
	interpreter.File.Header.readReserved()
	interpreter.File.Header.readChannels()
	interpreter.File.Header.readDimensions()
	interpreter.File.Header.readDepth()
	interpreter.File.Header.readColorMode()
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

func (fh *FileHeader) readChannels() {
	channelsStart := 15
	channelsEnd := 20
	fh.channels = binary.BigEndian.Uint16(fh.bytes[channelsStart:channelsEnd])
}

func (fh *FileHeader) readDimensions() {
	heightStart := 21
	heightEnd := 26
	widthStart := 27
	widthEnd := 32

	fh.height = binary.BigEndian.Uint16(fh.bytes[heightStart:heightEnd])
	fh.width = binary.BigEndian.Uint16(fh.bytes[widthStart:widthEnd])
}

func (fh *FileHeader) readDepth() {
	depthStart := 33
	depthEnd := 36
	fh.depth = binary.BigEndian.Uint16(fh.bytes[depthStart:depthEnd])
}

func (fh *FileHeader) readColorMode() {
	colorModeStart := 37
	colorModeEnd := 40
	fh.depth = binary.BigEndian.Uint16(fh.bytes[colorModeStart:colorModeEnd])
}
