package shape

import (
	"os"
	"github.com/pisdhooy/fmtbytes"
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
	rectangle.Top = fmtbytes.ReadBytesLong(file)
	rectangle.Left = fmtbytes.ReadBytesLong(file)
	rectangle.Bottom = fmtbytes.ReadBytesLong(file)
	rectangle.Right = fmtbytes.ReadBytesLong(file)
}

func (rectangle *Rectangle) ParseSliceFormat(file *os.File) {
	rectangle.Left = fmtbytes.ReadBytesLong(file)
	rectangle.Top = fmtbytes.ReadBytesLong(file)
	rectangle.Right = fmtbytes.ReadBytesLong(file)
	rectangle.Bottom = fmtbytes.ReadBytesLong(file)
}
