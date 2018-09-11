package slice

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
	"github.com/fabulousduck/pisdui/pisdui/util/shape"
)

type HeaderInterface interface {
	GetHeaderVersion() uint16
}

type Slice struct {
	Header *HeaderInterface
	Block  []*Block
}

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

type Block struct {
	ID                  uint32
	GroupID             uint32
	Origin              uint32
	AssocLayerID        uint32
	Name                string
	Type                uint32
	Dimensions          *shape.Rectangle
	Url                 string
	Target              string
	Message             string
	AltTag              string
	CellTextIsHTML      bool
	CellText            string
	HorizontalAlignment uint32
	VerticalAlignment   uint32
	AlphaColor          int
	Red                 int
	Green               int
	Blue                int
	DescriptorVersion   uint32
	Descriptor          *descriptor.Descriptor
}

func (headerCS6 HeaderCS6) GetHeaderVersion() uint16 {
	return 6
}

func (headerCS7 HeaderCS7) GetHeaderVersion() uint16 {
	return 7
}

func NewSlice() *Slice {
	return new(Slice)
}

func (slice *Slice) Parse(file *os.File) {

}
