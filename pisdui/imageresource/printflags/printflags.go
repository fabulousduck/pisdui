package printflags

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	printFlags.Labels = fsutil.ReadSingleByte(file) != 0
	printFlags.CropMarks = fsutil.ReadSingleByte(file) != 0
	printFlags.ColorBars = fsutil.ReadSingleByte(file) != 0
	printFlags.RegistrationMarks = fsutil.ReadSingleByte(file) != 0
	printFlags.Negative = fsutil.ReadSingleByte(file) != 0
	printFlags.Flip = fsutil.ReadSingleByte(file) != 0
	printFlags.Interpolate = fsutil.ReadSingleByte(file) != 0
	printFlags.Caption = fsutil.ReadSingleByte(file) != 0
	printFlags.MysteryByte = fsutil.ReadSingleByte(file) != 0
}
