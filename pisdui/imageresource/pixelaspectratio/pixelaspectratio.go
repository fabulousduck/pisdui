package pixelaspectratio

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	pixelAspectRatio.Version = fsutil.ReadBytesLong(file)
	Coordinates, _ := fsutil.ReadDouble(file) //TODO handle this error
	pixelAspectRatio.Coordinates = Coordinates
}
