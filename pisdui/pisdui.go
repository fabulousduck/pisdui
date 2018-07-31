package pisdui

import (
	"io/ioutil"
)

type ColorModeData struct {
}

type ImageResources struct {
}

type LayerMaskInfo struct {
}

type ImageData struct {
}

type File struct {
	Header         FileHeader
	ColorModeData  ColorModeData
	ImageResources ImageResources
	LayerMaskInfo  LayerMaskInfo
	ImageData      ImageData
}

type Pisdui struct {
	File         File
	FileContents []byte
}

func NewInterpreter() *Pisdui {
	return new(Pisdui)
}

func (interpreter *Pisdui) LoadFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	interpreter.FileContents = data
}

func (interpreter *Pisdui) Parse() {
	interpreter.ParseHeader()
	interpreter.parseColorModeData()
	interpreter.parseImageResources()
	interpreter.parseLayersAndMasks()
	interpreter.parseImageData()
}

func (interpreter *Pisdui) parseColorModeData() {

}

func (interpreter *Pisdui) parseImageResources() {

}

func (interpreter *Pisdui) parseLayersAndMasks() {

}

func (interpreter *Pisdui) parseImageData() {

}
