package pisdui

import (
	"fmt"
	"os"
)

type ColorModeData struct {
}

type ImageResources struct {
}

type LayerMaskInfo struct {
}

type ImageData struct {
}

type PSD struct {
	Header         FileHeader
	ColorModeData  ColorModeData
	ImageResources ImageResources
	LayerMaskInfo  LayerMaskInfo
	ImageData      ImageData
}

type Pisdui struct {
	PSD          PSD
	FileContents *os.File
}

func NewInterpreter() *Pisdui {
	return new(Pisdui)
}

func (interpreter *Pisdui) LoadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	interpreter.FileContents = file
}

func (interpreter *Pisdui) Parse() {
	interpreter.ParseHeader()
	interpreter.parseColorModeData()
	interpreter.parseImageResources()
	interpreter.parseLayersAndMasks()
	interpreter.parseImageData()
}
