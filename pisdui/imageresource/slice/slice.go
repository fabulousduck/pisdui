package slice

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/slice/header"
	util "github.com/fabulousduck/pisdui/pisdui/util/file"
	"github.com/fabulousduck/pisdui/pisdui/util/shape"
)

type Slice struct {
	Header *header.HeaderInterface
	Block  *[]Block
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

func (slice *Slice) GetTypeID() int {
	return 1050
}

func NewSlice() *Slice {
	return new(Slice)
}

func (slice *Slice) Parse(file *os.File) {
	sliceObject := new(Slice)
	headerVersion := util.ReadBytesLong(file)
	header := header.ParseHeader(file, headerVersion)
	sliceObject.Header = header

	switch headerVersion {
	case 6:

		break
	case 7:
		fallthrough
	case 8:
		//TODO: find a psd file that uses this and implement this
		break
	}
}

func NewBlock() *Block {
	return new(Block)
}

func (block *Block) Parse(file *os.File) {

}
