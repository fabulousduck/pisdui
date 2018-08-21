package printflags

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
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
	printFlags.Labels = util.ReadSingleByte(file) != 0
	printFlags.CropMarks = util.ReadSingleByte(file) != 0
	printFlags.ColorBars = util.ReadSingleByte(file) != 0
	printFlags.RegistrationMarks = util.ReadSingleByte(file) != 0
	printFlags.Negative = util.ReadSingleByte(file) != 0
	printFlags.Flip = util.ReadSingleByte(file) != 0
	printFlags.Interpolate = util.ReadSingleByte(file) != 0
	printFlags.Caption = util.ReadSingleByte(file) != 0
	printFlags.MysteryByte = util.ReadSingleByte(file) != 0
}
