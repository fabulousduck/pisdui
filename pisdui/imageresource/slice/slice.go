package slice

import (
	"os"

	"github.com/davecgh/go-spew/spew"

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
		//TODO: debug why this runs into an EOF
		//probably something todo with offset being incorrect
		//when reading beginsS
		spew.Dump(file.Seek(0, 1))
		blockObject := NewBlock()
		blockObject.Parse(file)
		spew.Dump(blockObject)
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
	// descriptorObject := descriptor.NewDescriptor()

	block.ID = util.ReadBytesLong(file)
	block.GroupID = util.ReadBytesLong(file)
	block.Origin = util.ReadBytesLong(file)
	block.AssocLayerID = util.ReadBytesLong(file)
	block.Name = util.ParseUnicodeString(file)
	block.Type = util.ReadBytesLong(file)

	dimensionsObject.Parse(file)

	block.Dimensions = dimensionsObject
	// block.Url = util.ParseUnicodeString(file)
	// block.Target = util.ParseUnicodeString(file)
	// block.Message = util.ParseUnicodeString(file)
	// block.AltTag = util.ParseUnicodeString(file)
	// block.CellTextIsHTML = util.ReadSingleByte(file) == 1
	// block.CellText = util.ParseUnicodeString(file)
	// block.HorizontalAlignment = util.ReadBytesLong(file)
	// block.VerticalAlignment = util.ReadBytesLong(file)
	// block.AlphaColor = util.ReadSingleByte(file)
	// block.Red = util.ReadSingleByte(file)
	// block.Green = util.ReadSingleByte(file)
	// block.Blue = util.ReadSingleByte(file)
	// block.DescriptorVersion = util.ReadBytesLong(file)

	// descriptorObject.Parse(file)

	// block.Descriptor = descriptorObject
}
