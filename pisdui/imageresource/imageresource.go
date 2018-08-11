package imageresource

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

/*ImageResourceData contains the resource blocks
used by the photoshop file and the length of the
section in the photoshop file*/
type ImageResourceData struct {
	Length         uint32
	ResourceBlocks []ResourceBlock
}

/*ResourceBlock contains the raw unparsed data from
a resource block in the photoshop file*/
type ResourceBlock struct {
	byteSize     uint32
	Signature    string
	ID           uint16
	PascalString string
	DataSize     uint32
	DataBlock    []byte
}

/*NewImageResources creates a new ImageResources struct
and returns a pointer to it.
This exists so the top level pisdui struct can create one
to prevent import cycles*/
func NewImageResources() *ImageResourceData {
	return new(ImageResourceData)
}

/*ParseImageResources will read all image resources located in
the photoshop file and will read them into the ImageResources struct*/
func (ir *ImageResourceData) ParseImageResources(file *os.File) {
	ir.Length = util.ReadBytesLong(file)
	var i uint32
	for i = 0; i < ir.Length; {
		block := ir.parseResourceBlock(file)
		ir.ResourceBlocks = append(ir.ResourceBlocks, *block)
		i += block.byteSize
	}
}

func (ir *ImageResourceData) parseResourceBlock(file *os.File) *ResourceBlock {
	block := new(ResourceBlock)
	block.Signature = util.ReadBytesString(file, 4)
	block.ID = util.ReadBytesShort(file)
	pascalString, stringLength := ir.parsePascalString(file)
	block.PascalString = pascalString
	block.DataSize = util.ReadBytesLong(file)
	block.DataBlock = util.ReadBytesNInt(file, block.DataSize)

	if block.DataSize%2 != 0 {
		util.ReadSingleByte(file)
	}

	block.byteSize = uint32(4 + 2 + stringLength + 4 + int(block.DataSize))
	return block
}

func (ir *ImageResourceData) parsePascalString(file *os.File) (string, int) {
	b := util.ReadSingleByte(file)
	if b == 0 {
		util.ReadSingleByte(file)
		return "", 1
	}

	s := util.ReadBytesString(file, b)

	if b%2 != 0 {
		util.ReadSingleByte(file)
	}
	return s, len(s)
}
