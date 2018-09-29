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

type Vec3_32 struct {
	X uint32
	Y uint32
	Z uint32
}

func NewRectangle() *Rectangle {
	return new(Rectangle)
}

func NewVec3_32() *Vec3_32 {
	return new(Vec3_32)
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

func (vec3_32 *Vec3_32) Parse(file *os.File) {
	vec3_32.X = util.ReadBytesLong(file)
	vec3_32.Y = util.ReadBytesLong(file)
	vec3_32.Z = util.ReadBytesLong(file)
}
