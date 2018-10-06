package slice

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/slice/header"
	"github.com/fabulousduck/pisdui/pisdui/util/shape"
	"github.com/pisdhooy/fsutil"
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
	headerVersion := fsutil.ReadBytesLong(file)
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
	block.ID = fsutil.ReadBytesLong(file)
	block.GroupID = fsutil.ReadBytesLong(file)
	block.Origin = fsutil.ReadBytesLong(file)
	if block.Origin == 1 {
		block.AssocLayerID = fsutil.ReadBytesLong(file)
	}
	block.Name = fsutil.ParseUnicodeString(file)
	block.Type = fsutil.ReadBytesLong(file)

	dimensionsObject.ParseSliceFormat(file)

	block.Dimensions = dimensionsObject

	block.Url = fsutil.ParseUnicodeString(file)
	block.Target = fsutil.ParseUnicodeString(file)
	block.Message = fsutil.ParseUnicodeString(file)
	block.AltTag = fsutil.ParseUnicodeString(file)
	block.CellTextIsHTML = fsutil.ReadSingleByte(file) == 1
	block.CellText = fsutil.ParseUnicodeString(file)
	block.HorizontalAlignment = fsutil.ReadBytesLong(file)
	block.VerticalAlignment = fsutil.ReadBytesLong(file)
	block.AlphaColor = fsutil.ReadSingleByte(file)
	block.Red = fsutil.ReadSingleByte(file)
	block.Green = fsutil.ReadSingleByte(file)
	block.Blue = fsutil.ReadSingleByte(file)
	block.DescriptorVersion = fsutil.ReadBytesLong(file)

	descriptorObject.Parse(file)

	block.Descriptor = descriptorObject
}
