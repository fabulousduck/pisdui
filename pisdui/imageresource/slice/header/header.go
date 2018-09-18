package header

import (
	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
	"github.com/fabulousduck/pisdui/pisdui/util/shape"
)

type HeaderCS6 struct {
	Version           uint32
	DescriptorVersion uint32
	Descriptor        *descriptor.Descriptor
}

//HeaderCS7 is for both CS7 and CS8.
//it is denoted to 7 since that changes from 6, where 7 and 8 are the same
type HeaderCS7 struct {
	Version           uint32
	BoundingRectangle *shape.Rectangle
	Name              string
	NumSlices         uint32
}

func (headerCS6 HeaderCS6) GetHeaderVersion() uint16 {
	return 6
}

func (headerCS7 HeaderCS7) GetHeaderVersion() uint16 {
	return 7
}
