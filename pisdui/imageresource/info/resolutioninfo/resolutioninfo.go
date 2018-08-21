package resolutioninfo

import (
	"fmt"
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

//Note: HorizontalResolution and VerticalResolution are byte buffers as they are fixed point integers
type Resolutioninfo struct {
	HorizontalResolution     []byte
	HorizontalResolutionUnit string
	WidthResolutionUnit      string
	VerticalResolution       []byte
	VerticalResolutionUnit   string
	HeightUnit               uint16
}

func (resolutioninfo *Resolutioninfo) GetTypeID() int {
	return 1005
}

func NewResolutionInfo() *Resolutioninfo {
	return new(Resolutioninfo)
}

func (resolutioninfo *Resolutioninfo) Parse(file *os.File) {
	pos, _ := file.Seek(0, 1)
	fmt.Println("file pointer pos : ", pos)
	resolutioninfo.HorizontalResolution = util.ReadRawBytes(file, 4)
	resolutioninfo.HorizontalResolutionUnit = parseUnit(util.ReadBytesShort(file))
	resolutioninfo.WidthResolutionUnit = parseUnit(util.ReadBytesShort(file))
	resolutioninfo.VerticalResolution = util.ReadRawBytes(file, 4)
	resolutioninfo.VerticalResolutionUnit = parseUnit(util.ReadBytesShort(file))
	resolutioninfo.HeightUnit = util.ReadBytesShort(file)
}

func parseUnit(unit uint16) string {
	opts := map[uint16]string{
		1: "inches",
		2: "cm",
		3: "points",
		4: "picas",
		5: "columns",
	}
	return opts[unit]
}
