package icc

import (
	"os"

	"github.com/davecgh/go-spew/spew"

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
	Magic           uint32
	Platform        string
	Flags           uint32
	Manufacturer    string
	Model           uint32
	Attributes      uint64
	RenderingIntent uint32
	Illuminant      *shape.Vec3_16
	Creator         string
	ProfileID       *Union
	Reserved        int
}

type DateTime struct {
	Year    uint16
	Month   uint16
	Day     uint16
	Hours   uint16
	Minutes uint16
	Seconds uint16
}

type Union struct {
	ID8  int
	ID16 uint16
	ID32 uint32
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

func NewUnion() *Union {
	return new(Union)
}

func (union *Union) Parse(file *os.File) {
	union.ID8 = util.ReadSingleByte(file)
	union.ID16 = util.ReadBytesShort(file)
	union.ID32 = util.ReadBytesLong(file)
}

func (header *Header) Parse(file *os.File) {
	// dateTimeObject := NewDateTime()
	// illuminantObject := shape.NewVec3_16()
	// profileIDObject := NewUnion()

	header.Size = util.ReadBytesLong(file)
	header.Cmmid = util.ReadBytesString(file, 4)
	spew.Dump(header)
	header.Version = util.ReadBytesLong(file)
	header.DeviceClass = util.ReadBytesString(file, 4)
	header.ColorSpace = util.ReadBytesString(file, 4)
	header.Pcs = util.ReadBytesString(file, 4)

	// dateTimeObject.Parse(file)

	// header.DateTime = dateTimeObject
	// header.Magic = util.ReadBytesLong(file)
	// header.Platform = util.ParseUnicodeString(file)
	// header.Flags = util.ReadBytesLong(file)
	// header.Manufacturer = util.ParseUnicodeString(file)
	// header.Model = util.ReadBytesLong(file)
	// header.Attributes = util.ReadBytesLongLong(file)
	// header.RenderingIntent = util.ReadBytesLong(file)

	// illuminantObject.Parse(file)

	// header.Illuminant = illuminantObject
	// header.Creator = util.ParseUnicodeString(file)

	// profileIDObject.Parse(file)

	// header.ProfileID = profileIDObject
	// header.Reserved = util.ReadSingleByte(file)
}
