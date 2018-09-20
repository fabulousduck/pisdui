package header

import "github.com/fabulousduck/pisdui/pisdui/util/shape"

type Header struct {
	Size            uint32
	Cmmid           uint32
	Version         uint32
	DeviceClass     string
	ColorSpace      string
	Pcs             string
	DateTime        *DateTime
	Magic           uint32
	Platform        string
	Flags           uint32
	Manufacturer    string
	Mdodel          uint32
	Attributes      uint64
	RenderingIntent uint32
	Illuminant      *shape.Vec3_16
	Creator         string
	ProfileID       Union
	Reserved        byte
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
	ID8  byte
	ID16 uint16
	ID32 uint32
}
