package printflags

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type PrintFlags struct {
	Labels            bool
	CropMarks         bool
	ColorBars         bool
	RegistrationMarks bool
	Negative          bool
	Flip              bool
	Interpolate       bool
	Caption           bool
	MysteryByte       bool //TODO: figure out what this is

}

func (printFlags *PrintFlags) GetTypeID() int {
	return 1011
}

func NewPrintFlags() *PrintFlags {
	return new(PrintFlags)
}

func (printFlags *PrintFlags) Parse(file *os.File) {
	printFlags.Labels = fmtbytes.ReadSingleByte(file) != 0
	printFlags.CropMarks = fmtbytes.ReadSingleByte(file) != 0
	printFlags.ColorBars = fmtbytes.ReadSingleByte(file) != 0
	printFlags.RegistrationMarks = fmtbytes.ReadSingleByte(file) != 0
	printFlags.Negative = fmtbytes.ReadSingleByte(file) != 0
	printFlags.Flip = fmtbytes.ReadSingleByte(file) != 0
	printFlags.Interpolate = fmtbytes.ReadSingleByte(file) != 0
	printFlags.Caption = fmtbytes.ReadSingleByte(file) != 0
	printFlags.MysteryByte = fmtbytes.ReadSingleByte(file) != 0
}
