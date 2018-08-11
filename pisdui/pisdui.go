package pisdui

import (
	"fmt"
	"os"

	"github.com/fabulousduck/pisdui/pisdui/colormode"
	"github.com/fabulousduck/pisdui/pisdui/header"
	"github.com/fabulousduck/pisdui/pisdui/imagedata"
	"github.com/fabulousduck/pisdui/pisdui/imageresource"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask"
)

/*PSD contains all parsed data from the photoshop file*/
type PSD struct {
	Fp             *os.File
	Header         *header.Data
	ColorModeData  *colormode.Data
	ImageResources *imageresource.Data
	LayerMaskInfo  *layerandmask.Data
	ImageData      *imagedata.Data
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

	psd.Fp = file
}

/*Parse takes the loaded file and parses it into
usable structs separated into the different main
parts of the file*/
func (psd *PSD) Parse() {
	header := header.NewData()
	header.Parse(psd.Fp)
	psd.Header = header

	colorModeData := colormode.NewData()
	colorModeData.Parse(psd.Fp, psd.Header.ColorMode)
	psd.ColorModeData = colorModeData

	imageResourceData := imageresource.NewData()
	imageResourceData.Parse(psd.Fp)
	psd.ImageResources = imageResourceData

}
