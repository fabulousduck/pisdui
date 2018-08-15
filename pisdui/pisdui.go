package pisdui

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/colormode"
	"github.com/fabulousduck/pisdui/pisdui/header"
	"github.com/fabulousduck/pisdui/pisdui/imagedata"
	"github.com/fabulousduck/pisdui/pisdui/imageresource"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask"
)

// PSD contains all parsed data from the photoshop file
type PSD struct {
	Fp             *os.File
	Header         *header.Data
	ColorModeData  *colormode.Data
	ImageResources *imageresource.Data
	LayerMaskInfo  *layerandmask.Data
	ImageData      *imagedata.Data
}

// NewPSD creates a new PSD struct
// to read the file pointer into and
// the data read from the photoshop file
func NewPSD(path string) (*PSD, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &PSD{Fp: file}, nil
}

// Parse takes the loaded file and parses it into
// usable structs separated into the different main
// parts of the file
func (psd *PSD) Parse() *PSD {
	h := header.NewData()
	h.Parse(psd.Fp)
	psd.Header = h

	colorModeData := colormode.NewData()
	colorModeData.Parse(psd.Fp, psd.Header.ColorMode)
	psd.ColorModeData = colorModeData

	imageResourceData := imageresource.NewData()
	imageResourceData.Parse(psd.Fp)
	psd.ImageResources = imageResourceData

	return psd
}
