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
	rectangle.Left = util.ReadBytesLong(file)
	rectangle.Bottom = util.ReadBytesLong(file)
	rectangle.Right = util.ReadBytesLong(file)
}

func (rectangle *Rectangle) ParseSliceFormat(file *os.File) {
	rectangle.Left = util.ReadBytesLong(file)
	rectangle.Top = util.ReadBytesLong(file)
	rectangle.Right = util.ReadBytesLong(file)
	rectangle.Bottom = util.ReadBytesLong(file)
}
