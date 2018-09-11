package shape

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type Rectangle struct {
	Top    uint32
	Left   uint32
	Bottom uint32
	Right  uint32
}

func NewRectangle() *Rectangle {
	return new(Rectangle)
}

func (rectangle *Rectangle) Parse(file *os.File) {
	rectangle.Top = util.ReadBytesLong(file)
}
