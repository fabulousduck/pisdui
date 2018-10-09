package printflaginfo

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
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
	printFlagInfo.Version = fmtbytes.ReadBytesShort(file)
	printFlagInfo.CenterCrop = fmtbytes.ReadSingleByte(file)
	printFlagInfo.BufferByte = fmtbytes.ReadSingleByte(file)
	printFlagInfo.BleedWidth = fmtbytes.ReadBytesLong(file)
	printFlagInfo.BleedWidthScale = fmtbytes.ReadBytesShort(file)
}
