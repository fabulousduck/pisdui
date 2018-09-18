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

func NewSlice() *Slice {
	return new(Slice)
}

func (slice *Slice) Parse(file *os.File) {
	header := new(header)
}
