package printflaginfo

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	printFlagInfo.Version = fsutil.ReadBytesShort(file)
	printFlagInfo.CenterCrop = fsutil.ReadSingleByte(file)
	printFlagInfo.BufferByte = fsutil.ReadSingleByte(file)
	printFlagInfo.BleedWidth = fsutil.ReadBytesLong(file)
	printFlagInfo.BleedWidthScale = fsutil.ReadBytesShort(file)
}
