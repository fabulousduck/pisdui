package imageresource

import (
	"errors"
	"fmt"
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/printflags"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/info/printflaginfo"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/info/resolutioninfo"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
	"github.com/fabulousduck/pisdui/pisdui/util"
)

type parsedResourceBlock interface {
	GetTypeID() int
}

//Data contains the resource blocks
//used by the photoshop file and the length of the
//section in the photoshop file
type Data struct {
	Length         uint32
	ResourceBlocks []*ResourceBlock
}

//ResourceBlock contains the raw unparsed data from
//a resource block in the photoshop file
type ResourceBlock struct {
	Signature           string
	ID                  uint16
	PascalString        string
	DataSize            uint32
	ParsedResourceBlock parsedResourceBlock
}

//NewData creates a new ImageResources struct
//and returns a pointer to it.
//This exists so the top level pisdui struct can create one
//to prevent import cycles
func NewData() *Data {
	return new(Data)
}

//Parse will read all image resources located in
//the photoshop file and will read them into the ImageResources struct
func (resourceBlockSection *Data) Parse(file *os.File) error {
	resourceBlockSection.Length = util.ReadBytesLong(file)

	currPos, _ := file.Seek(0, 1)
	endPos := int(currPos) + int(resourceBlockSection.Length)
	for int(currPos) < endPos {
		resourceBlockSection.ResourceBlocks = append(resourceBlockSection.ResourceBlocks, resourceBlockSection.parseResourceBlock(file))
		pos, _ := file.Seek(0, 1)
		currPos = pos
		if resourceBlockSection.ResourceBlocks[len(resourceBlockSection.ResourceBlocks)-1].Signature != "8BIM" {
			return errors.New("non 8BIM signature")
		}
	}
	fmt.Println("pos after image resource block parsing : ", currPos)
	return nil
}

func (resourceBlockSection *Data) parseResourceBlock(file *os.File) *ResourceBlock {
	block := new(ResourceBlock)
	block.Signature = util.ReadBytesString(file, 4)

	block.ID = util.ReadBytesShort(file)

	pascalString := util.ParsePascalString(file)

	block.PascalString = pascalString
	block.DataSize = util.ReadBytesLong(file)

	block.ParsedResourceBlock = parseResourceBlock(file, block.ID, block.DataSize)

	if block.DataSize%2 != 0 {
		util.ReadSingleByte(file)
	}
	return block
}

func parseResourceBlock(file *os.File, id uint16, size uint32) parsedResourceBlock {
	var p parsedResourceBlock
	switch id {
	case 1005:
		resolutioninfoObject := resolutioninfo.NewResolutionInfo()
		resolutioninfoObject.Parse(file)
		p = resolutioninfoObject
	case 1011:
		printFlagsObject := printflags.NewPrintFlags()
		printFlagsObject.Parse(file)
		p = printFlagsObject
	case 1074:
		fallthrough
	case 1075:
		fallthrough
	case 1076:
		fallthrough
	case 1080:
		fallthrough
	case 1082:
		fallthrough
	case 1083:
		fallthrough
	case 1088:
		descriptorVersion := util.ReadBytesLong(file)
		descriptorObject := descriptor.NewDescriptor()
		descriptorObject.Parse(file)
		descriptorObject.Version = descriptorVersion
		p = descriptorObject
		break

	case 10000:
		printFlagInfoObject := printflaginfo.NewPrintFlagInfo()
		printFlagInfoObject.Parse(file)
		p = printFlagInfoObject
		break
	default:
		util.ReadBytesNInt(file, size)
		break
	}
	return p
}
