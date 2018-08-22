package printflaginfo

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type PrintFlagInfo struct {
	Version         uint16
	CenterCrop      int
	BufferByte      int
	BleedWidth      uint32
	BleedWidthScale uint16
}

func (printFlagInfo *PrintFlagInfo) GetTypeID() int {
	return 10000
}

func NewPrintFlagInfo() *PrintFlagInfo {
	return new(PrintFlagInfo)
}

func (printFlagInfo *PrintFlagInfo) Parse(file *os.File) {
	printFlagInfo.Version = util.ReadBytesShort(file)
	printFlagInfo.CenterCrop = util.ReadSingleByte(file)
	printFlagInfo.BufferByte = util.ReadSingleByte(file)
	printFlagInfo.BleedWidth = util.ReadBytesLong(file)
	printFlagInfo.BleedWidthScale = util.ReadBytesShort(file)
}
