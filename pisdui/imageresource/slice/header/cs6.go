package header

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util/shape"
	"github.com/pisdhooy/fmtbytes"
)

type HeaderCS6 struct {
	Version           uint32
	BoundingRectangle *shape.Rectangle
	GroupName         string
	NumSlices         uint32
}

func NewCS6Header() *HeaderCS6 {
	return new(HeaderCS6)
}

func (header *HeaderCS6) Parse(file *os.File) {
	rectangleObject := shape.NewRectangle()
	rectangleObject.Parse(file)
	header.BoundingRectangle = rectangleObject
	header.GroupName = fmtbytes.ParseUnicodeString(file)
	header.NumSlices = fmtbytes.ReadBytesLong(file)
}
