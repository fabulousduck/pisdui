package imageresource

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

/*Data contains the resource blocks
used by the photoshop file and the length of the
section in the photoshop file*/
type Data struct {
	Length         uint32
	ResourceBlocks []*ResourceBlock
}

/*ResourceBlock contains the raw unparsed data from
a resource block in the photoshop file*/
type ResourceBlock struct {
	Signature           string
	ID                  uint16
	PascalString        string
	DataSize            uint32
	DataBlock           []byte
	ParsedResourceBlock parsedResourceBlock
}

/*NewData creates a new ImageResources struct
and returns a pointer to it.
This exists so the top level pisdui struct can create one
to prevent import cycles*/
func NewData() *Data {
	return new(Data)
}

/*Parse will read all image resources located in
the photoshop file and will read them into the ImageResources struct*/
func (resourceBlockSection *Data) Parse(file *os.File) {
	resourceBlockSection.Length = util.ReadBytesLong(file)

	currPos, _ := file.Seek(0, 1)
	endPos := int(currPos) + int(resourceBlockSection.Length)

	for p, _ := file.Seek(0, 1); int(p) < endPos; {
		resourceBlockSection.ResourceBlocks = append(
			resourceBlockSection.ResourceBlocks,
			resourceBlockSection.parseResourceBlock(file))
	}

}

func (resourceBlockSection *Data) parseResourceBlock(file *os.File) *ResourceBlock {
	readByteCount := 0
	block := new(ResourceBlock)
	block.Signature = util.ReadBytesString(file, 4)

	block.ID = util.ReadBytesShort(file)

	pascalString := util.ParsePascalString(file)

	block.PascalString = pascalString
	block.DataSize = util.ReadBytesLong(file)
	readByteCount += 4

	block.ParsedResourceBlock = parseResourceBlock(file, block.ID)

	if block.DataSize%2 != 0 {
		util.ReadSingleByte(file)
	}
	return block
}

func parseResourceBlock(file *os.File, id uint16) parsedResourceBlock {
	var p parsedResourceBlock
	switch id {
	case 1088:
		p = parseDescriptor(file)
		break
	}
	return p
}
