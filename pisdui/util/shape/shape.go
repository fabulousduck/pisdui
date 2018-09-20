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

type Vec3_16 struct {
	X uint16
	Y uint16
	Z uint16
}

func NewRectangle() *Rectangle {
	return new(Rectangle)
}

func NewVec3_16() *Vec3_16 {
	return new(Vec3_16)
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

func (vec3_16 *Vec3_16) Parse(file *os.File) {
	vec3_16.X = util.ReadBytesShort(file)
	vec3_16.Y = util.ReadBytesShort(file)
	vec3_16.Z = util.ReadBytesShort(file)
}
