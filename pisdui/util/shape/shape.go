package shape

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	rectangle.Top = fsutil.ReadBytesLong(file)
	rectangle.Left = fsutil.ReadBytesLong(file)
	rectangle.Bottom = fsutil.ReadBytesLong(file)
	rectangle.Right = fsutil.ReadBytesLong(file)
}

func (rectangle *Rectangle) ParseSliceFormat(file *os.File) {
	rectangle.Left = fsutil.ReadBytesLong(file)
	rectangle.Top = fsutil.ReadBytesLong(file)
	rectangle.Right = fsutil.ReadBytesLong(file)
	rectangle.Bottom = fsutil.ReadBytesLong(file)
}
