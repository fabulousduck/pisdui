package pixelaspectratio

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type PixelAspectRatio struct {
	Version     uint32
	Coordinates float64
}

func (pixelAspectRatio *PixelAspectRatio) GetTypeID() int {
	return 1064
}

func NewPixelAspectRatio() *PixelAspectRatio {
	return new(PixelAspectRatio)
}

func (pixelAspectRatio *PixelAspectRatio) Parse(file *os.File) {
	pixelAspectRatio.Version = util.ReadBytesLong(file)
	Coordinates, _ := util.ReadDouble(file) //TODO handle this error
	pixelAspectRatio.Coordinates = Coordinates
}
