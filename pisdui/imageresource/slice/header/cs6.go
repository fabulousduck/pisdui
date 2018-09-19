package header

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
	"github.com/fabulousduck/pisdui/pisdui/util/shape"
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
	header.GroupName = util.ParseUnicodeString(file)
	header.NumSlices = util.ReadBytesLong(file)
}
