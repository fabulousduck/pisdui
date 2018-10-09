package slice

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/slice/header"
	"github.com/fabulousduck/pisdui/pisdui/util/shape"
	"github.com/pisdhooy/fmtbytes"
)

type Slice struct {
	Header *header.HeaderInterface
	Blocks []*Block
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
	headerVersion := fmtbytes.ReadBytesLong(file)
	slice.Header = header.ParseHeader(file, headerVersion)
	switch headerVersion {
	case 6:
		headerObject := header.CastBackCS6(*slice.Header)
		for i := 0; i < int(headerObject.NumSlices); i++ {
			block := NewBlock()
			block.Parse(file)
			slice.Blocks = append(slice.Blocks, block)
		}
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
	dimensionsObject := shape.NewRectangle()
	descriptorObject := descriptor.NewDescriptor()
	block.ID = fmtbytes.ReadBytesLong(file)
	block.GroupID = fmtbytes.ReadBytesLong(file)
	block.Origin = fmtbytes.ReadBytesLong(file)
	if block.Origin == 1 {
		block.AssocLayerID = fmtbytes.ReadBytesLong(file)
	}
	block.Name = fmtbytes.ParseUnicodeString(file)
	block.Type = fmtbytes.ReadBytesLong(file)

	dimensionsObject.ParseSliceFormat(file)

	block.Dimensions = dimensionsObject

	block.Url = fmtbytes.ParseUnicodeString(file)
	block.Target = fmtbytes.ParseUnicodeString(file)
	block.Message = fmtbytes.ParseUnicodeString(file)
	block.AltTag = fmtbytes.ParseUnicodeString(file)
	block.CellTextIsHTML = fmtbytes.ReadSingleByte(file) == 1
	block.CellText = fmtbytes.ParseUnicodeString(file)
	block.HorizontalAlignment = fmtbytes.ReadBytesLong(file)
	block.VerticalAlignment = fmtbytes.ReadBytesLong(file)
	block.AlphaColor = fmtbytes.ReadSingleByte(file)
	block.Red = fmtbytes.ReadSingleByte(file)
	block.Green = fmtbytes.ReadSingleByte(file)
	block.Blue = fmtbytes.ReadSingleByte(file)
	block.DescriptorVersion = fmtbytes.ReadBytesLong(file)

	descriptorObject.Parse(file)

	block.Descriptor = descriptorObject
}
