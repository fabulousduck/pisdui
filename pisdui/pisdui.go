package pisdui

import (
	"fmt"
	"os"

	"github.com/fabulousduck/pisdui/pisdui/header"
	"github.com/fabulousduck/pisdui/pisdui/imageresource"
)

type LayerMaskInfo struct {
}

type ImageData struct {
}

type PSD struct {
	Fp             *os.File
	Header         header.FileHeader
	ColorModeData  ColorModeData
	ImageResources imageresource.ImageResourceData
	LayerMaskInfo  LayerMaskInfo
	ImageData      ImageData
}

/*NewPSD creates a new PSD struct
to read the file pointer into and
the data read from the photoshop file*/
func NewPSD() *PSD {
	return new(PSD)
}

/*LoadFile loads opens the file and
places the file pointer <*os.File>
into the PSD object*/
func (psd *PSD) LoadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	interpreter.FileContents = file
}

/*Parse takes the loaded file and parses it into
usable structs separated into the different main
parts of the file*/
func (psd *PSD) Parse() {
	header := header.NewFileHeader()
	PSD.Header = header.ParseHeader()

	psd.ParseColorModeData()
	psd.ParseImageResources()
	psd.parseLayersAndMasks()
	psd.parseImageData()
}
