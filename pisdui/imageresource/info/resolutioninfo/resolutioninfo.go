package resolutioninfo

import (
	"encoding/binary"
	"os"
)

type Resolutioninfo struct {
	HorizontalResolution     uint32
	HorizontalResolutionUnit uint32
	WidthResolutionUnit      uint32
	VerticalResolution       uint32
	VerticalResolutionUnit   uint32
	HeightUnit               uint32
}

func (resolutioninfo *Resolutioninfo) GetTypeID() int {
	return 1005
}

func NewResolutionInfo() *Resolutioninfo {
	return new(Resolutioninfo)
}

func (resolutioninfo *Resolutioninfo) Parse(file *os.File) {
	resolutioninfo.HorizontalResolution = readLittleEndianLong(file)
	// resolutioninfo.HorizontalResolutionUnit = util.ReadBytesLong(file)
	// resolutioninfo.WidthResolutionUnit = util.ReadBytesLong(file)
	// resolutioninfo.VerticalResolution = util.ReadBytesLong(file)
	// resolutioninfo.VerticalResolutionUnit = util.ReadBytesLong(file)
	// resolutioninfo.HeightUnit = util.ReadBytesLong(file)
}

//because little endian is fucking memes
func readLittleEndianLong(file *os.File) uint32 {
	buffer := make([]byte, 4)
	_, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint32(buffer)
}
