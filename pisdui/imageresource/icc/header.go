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
	Magic           string
	Platform        string
	Flags           uint32
	Manufacturer    string
	Model           uint32
	Attributes      uint64
	RenderingIntent uint32
	Illuminant      *shape.Vec3_16
	Creator         uint32
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
	ID8  []int    //128bytes
	ID16 []uint16 //128bytes
	ID32 []uint32 //128bytes
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
	for i := 0; i < 16; i++ {
		union.ID8 = append(union.ID8, util.ReadSingleByte(file))
	}
	for i := 0; i < 8; i++ {
		union.ID16 = append(union.ID16, util.ReadBytesShort(file))
	}
	for i := 0; i < 4; i++ {
		union.ID32 = append(union.ID32, util.ReadBytesLong(file))
	}
}

func (header *Header) Parse(file *os.File) {
	dateTimeObject := NewDateTime()
	illuminantObject := shape.NewVec3_16()
	profileIDObject := NewUnion()

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
	spew.Dump(file.Seek(0, 1))

	header.Manufacturer = util.ReadBytesString(file, 4)
	header.Model = util.ReadBytesLong(file)
	header.Attributes = util.ReadBytesLongLong(file)

	header.RenderingIntent = util.ReadBytesLong(file)

	illuminantObject.Parse(file)

	header.Illuminant = illuminantObject
	header.Creator = util.ReadBytesLong(file)
	profileIDObject.Parse(file)

	header.ProfileID = profileIDObject
	header.Reserved = util.ReadSingleByte(file)
	spew.Dump(header)

}
