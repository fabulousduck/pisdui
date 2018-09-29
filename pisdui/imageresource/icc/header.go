package icc

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
	"github.com/fabulousduck/pisdui/pisdui/util/shape"
)

type Header struct {
	Size            uint32
	Cmmid           string
	Version         uint32
	DeviceClass     string
	ColorSpace      string
	Pcs             string
	DateTime        *DateTime
	Magic           string
	Platform        string
	Flags           uint32
	Manufacturer    string
	Model           uint32
	Attributes      uint64
	RenderingIntent uint32
	Illuminant      *shape.Vec3_32
	Creator         string
	ProfileId       []byte
	Reserved        []byte
}

type DateTime struct {
	Year    uint16
	Month   uint16
	Day     uint16
	Hours   uint16
	Minutes uint16
	Seconds uint16
}

func NewHeader() *Header {
	return new(Header)
}

func NewDateTime() *DateTime {
	return new(DateTime)
}

func (dateTime *DateTime) Parse(file *os.File) {
	dateTime.Year = util.ReadBytesShort(file)
	dateTime.Month = util.ReadBytesShort(file)
	dateTime.Day = util.ReadBytesShort(file)
	dateTime.Hours = util.ReadBytesShort(file)
	dateTime.Minutes = util.ReadBytesShort(file)
	dateTime.Seconds = util.ReadBytesShort(file)
}

func (header *Header) Parse(file *os.File) {
	dateTimeObject := NewDateTime()
	illuminantObject := shape.NewVec3_32()

	header.Size = util.ReadBytesLong(file)
	header.Cmmid = util.ReadBytesString(file, 4)
	header.Version = util.ReadBytesLong(file)
	header.DeviceClass = util.ReadBytesString(file, 4)
	header.ColorSpace = util.ReadBytesString(file, 4)
	header.Pcs = util.ReadBytesString(file, 4)

	dateTimeObject.Parse(file)

	header.DateTime = dateTimeObject

	header.Magic = util.ReadBytesString(file, 4)
	header.Platform = util.ReadBytesString(file, 4)
	header.Flags = util.ReadBytesLong(file)

	header.Manufacturer = util.ReadBytesString(file, 4)

	header.Model = util.ReadBytesLong(file)
	header.Attributes = util.ReadBytesLongLong(file)

	header.RenderingIntent = util.ReadBytesLong(file)

	illuminantObject.Parse(file)

	header.Illuminant = illuminantObject
	header.Creator = util.ReadBytesString(file, 4)

	header.ProfileId = util.ReadBytesNInt(file, 16)
	header.Reserved = util.ReadBytesNInt(file, 28)
}

func (header *Header) GetFullname(field string) string {
	nameMap := map[string]string{
		"ADBE": "Adobe Systems Inc.",
		"ACMS": "Agfa Graphics N.V. ",
		"appl": "Apple Computer",
		"CCMS": "Canon",
		"UCCM": "Canon",
		"UCMS": "Canon",
		"EFI":  "EFI",
		"FF ":  "Fuji Film Electronic Imaging",
		"EXAC": "ExactCODE GmbH",
		"HCMM": "Global Graphics Software Inc",
		"argl": "Graeme Gill",
		"LgoS": "GretagMacbeth",
		"HDM ": "Heidelberger Druckmaschinen AG",
		"lcms": "Hewlett Packard ",
		"RMIX": "ICC",
		"KCMS": "Kodak",
		"MCML": "Konica Minolta",
		"WCS":  "Microsoft",
		"SIGN": "Mutoh",
		"ONYX": "Onyx Graphics",
		"RGMS": "Rolf Gierling Multitools",
		"SICC": "SampleICC",
		"TCMM": "Toshiba TEC Corporation",
		"32BT": "the imaging factory",
		"vivo": "Vivo Mobile Communication",
		"WTG ": "Ware To Go",
		"zc00": "Zoran Corporation",
	}

	if val, ok := nameMap[field]; ok {
		return val
	}
	return field
}
