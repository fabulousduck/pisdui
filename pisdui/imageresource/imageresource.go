package imageresource

import (
	"errors"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/pisdhooy/fmtbytes"
	"github.com/pisdhooy/icc"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/backgroundcolor"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/id"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/measurementscale"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/pixelaspectratio"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/sec"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/slice"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/xmp"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/printflags"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/printscale"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/info/printflaginfo"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/info/resolutioninfo"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/info/version"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
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
	resourceBlockSection.Length = fmtbytes.ReadBytesLong(file)

	currPos, _ := file.Seek(0, 1)
	endPos := currPos + int64(resourceBlockSection.Length)
	for currPos < endPos {
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
	block.Signature = fmtbytes.ReadBytesString(file, 4)

	block.ID = fmtbytes.ReadBytesShort(file)

	pascalString := fmtbytes.ParsePascalString(file)

	block.PascalString = pascalString
	block.DataSize = fmtbytes.ReadBytesLong(file)

	block.ParsedResourceBlock = parseResourceBlockData(file, block.ID, block.DataSize)

	if block.DataSize%2 != 0 {
		fmtbytes.ReadSingleByte(file)
	}
	return block
}

func parseResourceBlockData(file *os.File, resourceId uint16, size uint32) parsedResourceBlock {
	var p parsedResourceBlock
	spew.Dump(resourceId)
	switch resourceId {
	case 1005:
		resolutioninfoObject := resolutioninfo.NewResolutionInfo()
		resolutioninfoObject.Parse(file)
		p = resolutioninfoObject
	case 1010:
		backgroundColorObject := backgroundcolor.NewBackgroundColor()
		backgroundColorObject.Parse(file)
		p = backgroundColorObject
	case 1011:
		printFlagsObject := printflags.NewPrintFlags()
		printFlagsObject.Parse(file)
		p = printFlagsObject
	case 1039:
		ICCProfileObject := icc.NewICCProfile()
		ICCProfileObject.Parse(file)
		p = ICCProfileObject
	case 1044:
		IDObject := id.NewID()
		IDObject.Parse(file)
		p = IDObject
	case 1050:
		sliceObject := slice.NewSlice()
		sliceObject.Parse(file)
		p = sliceObject
	case 1057:
		versionObject := version.NewVersion()
		versionObject.Parse(file)
		p = versionObject
	case 1060:
		XMPObject := xmp.NewXMP()
		XMPObject.Parse(file, size)
		p = XMPObject
		spew.Dump(file.Seek(0, 1))
	case 1061:
		DigestObject := sec.NewSec()
		DigestObject.Parse(file)
		p = DigestObject
	case 1062:
		printScaleObject := printscale.NewPrintScale()
		printScaleObject.Parse(file)
		p = printScaleObject
	case 1064:
		pixelAspectRatioObject := pixelaspectratio.NewPixelAspectRatio()
		pixelAspectRatioObject.Parse(file)
		p = pixelAspectRatioObject
	case 1074:
		measurementScaleObject := measurementscale.NewMeasurementScale()
		measurementScaleObject.Parse(file)
		p = measurementScaleObject
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
		fmt.Println("PRINT VERSION PARSE")
		descriptorVersion := fmtbytes.ReadBytesLong(file)
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
		fmt.Printf("sig : %d dataSize %d\n", resourceId, size)
		fmtbytes.ReadBytesNInt(file, size)
		break
	}

	return p
}
