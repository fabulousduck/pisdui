package imageresource

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/pisdhooy/fmtbytes"
	"github.com/pisdhooy/icc"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/backgroundcolor"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/exif"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/global/altitude"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/global/angle"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/id"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/measurementscale"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/pixelaspectratio"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/sec"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/slice"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/thumbnail"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/unicode"
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
	Length           uint32
	ResolutionInfo   *resolutioninfo.Resolutioninfo
	BackgroundColor  *backgroundcolor.BackgroundColor
	PrintFlags       *printflags.PrintFlags
	Thumbnail        *thumbnail.Thumbnail
	Angle            *angle.Angle
	ICCProfile       *icc.ICCProfile
	ID               *id.ID
	UnicodeString    *unicode.UnicodeString
	Altitude         *altitude.Altitude
	Slice            *slice.Slice
	Version          *version.Version
	Exif             *exif.Exif
	XMP              *xmp.XMP
	Digest           *sec.Sec
	Printscale       *printscale.PrintScale
	Pixelaspectratio *pixelaspectratio.PixelAspectRatio
	MeasurementSacle *measurementscale.MeasurementScale
	Descriptor       *descriptor.Descriptor
	PrintFlagInfo    *printflaginfo.PrintFlagInfo
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
		resourceBlockSection.parseResourceBlockData(file)
		pos, _ := file.Seek(0, 1)
		currPos = pos
	}
	fmt.Println("pos after image resource block parsing : ", currPos)
	return nil
}

func (imageResourceData *Data) parseResourceBlockData(file *os.File) {
	signature := fmtbytes.ReadBytesString(file, 4)
	if signature != "8BIM" {
		panic("non 8BIM resource block found")
	}
	resourceID := fmtbytes.ReadBytesShort(file)
	pascalString := fmtbytes.ParsePascalString(file)
	if pascalString != "" {
		//Todo : find a psd that uses this an parse it correctly
	}
	size := fmtbytes.ReadBytesLong(file)

	switch resourceID {
	case 1005:
		imageResourceData.ResolutionInfo = resolutioninfo.NewResolutionInfo()
		imageResourceData.ResolutionInfo.Parse(file)
	case 1010:
		imageResourceData.BackgroundColor = backgroundcolor.NewBackgroundColor()
		imageResourceData.BackgroundColor.Parse(file)
	case 1011:
		imageResourceData.PrintFlags = printflags.NewPrintFlags()
		imageResourceData.PrintFlags.Parse(file)
	case 1036:
		imageResourceData.Thumbnail = thumbnail.NewThumbnail()
		imageResourceData.Thumbnail.Parse(file)
	case 1037:
		imageResourceData.Angle = angle.NewAngle()
		imageResourceData.Angle.Parse(file)
	case 1039:
		imageResourceData.ICCProfile = icc.NewICCProfile()
		imageResourceData.ICCProfile.Parse(file)
	case 1044:
		imageResourceData.ID = id.NewID()
		imageResourceData.ID.Parse(file)
	case 1045:
		imageResourceData.UnicodeString = unicode.NewUnicodeString()
		imageResourceData.UnicodeString.Parse(file)
	case 1049:
		imageResourceData.Altitude = altitude.NewAltitude()
		imageResourceData.Altitude.Parse(file)
	case 1050:
		imageResourceData.Slice = slice.NewSlice()
		imageResourceData.Slice.Parse(file)
	case 1057:
		imageResourceData.Version = version.NewVersion()
		imageResourceData.Version.Parse(file)
	case 1058:
		imageResourceData.Exif = exif.NewExif()
		imageResourceData.Exif.Parse(file, size)
	case 1060:
		imageResourceData.XMP = xmp.NewXMP()
		imageResourceData.XMP.Parse(file, size)
	case 1061:
		imageResourceData.Digest = sec.NewSec()
		imageResourceData.Digest.Parse(file)
	case 1062:
		imageResourceData.Printscale = printscale.NewPrintScale()
		imageResourceData.Printscale.Parse(file)
	case 1064:
		imageResourceData.Pixelaspectratio = pixelaspectratio.NewPixelAspectRatio()
		imageResourceData.Pixelaspectratio.Parse(file)
	case 1074:
		imageResourceData.MeasurementSacle = measurementscale.NewMeasurementScale()
		imageResourceData.MeasurementSacle.Parse(file)
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
		imageResourceData.Descriptor = descriptor.NewDescriptor()
		imageResourceData.Descriptor.Version = fmtbytes.ReadBytesLong(file)
		imageResourceData.Descriptor.Parse(file)
	case 10000:
		imageResourceData.PrintFlagInfo = printflaginfo.NewPrintFlagInfo()
		imageResourceData.PrintFlagInfo.Parse(file)
	default:
		fmt.Println("RESOURCE ID")
		spew.Dump(resourceID)
		fmtbytes.ReadBytesNInt(file, size)
	}
	if size%2 != 0 {
		fmtbytes.ReadSingleByte(file)
	}
}
