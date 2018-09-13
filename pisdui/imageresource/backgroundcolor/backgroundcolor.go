package backgroundcolor

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type BackgroundColor struct {
	ColorSpaceId string
	ColorData    []uint8
}

func (backgroundColor *BackgroundColor) GetTypeID() int {
	return 1010
}

func NewBackgroundColor() *BackgroundColor {
	return new(BackgroundColor)
}

func (backgroundColor *BackgroundColor) Parse(file *os.File) {
	backgroundColor.parseColorSpaceID(file)
	backgroundColor.ColorData = util.ReadIntoArray8(file, 4)
}

func (backgroundcolor *BackgroundColor) parseColorSpaceID(file *os.File) {
	colorSpaceID := util.ReadBytesShort(file)
	switch colorSpaceID {
	case 0:
		backgroundcolor.ColorSpaceId = "RGB"
		break
	case 1:
		backgroundcolor.ColorSpaceId = "HSB"
		break
	case 2:
		backgroundcolor.ColorSpaceId = "CMYK"
		break
	case 7:
		backgroundcolor.ColorSpaceId = "Lab"
		break
	case 8:
		backgroundcolor.ColorSpaceId = "Grayscale"
		break
	}
}
