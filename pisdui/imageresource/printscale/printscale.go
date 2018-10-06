package printscale

import (
	"os"

	"github.com/pisdhooy/fsutil"
)

type PrintScale struct {
	Style     string
	XLocation float32
	YLocation float32
	Scale     float32
}

func (printScale *PrintScale) GetTypeID() int {
	return 1062
}

func NewPrintScale() *PrintScale {
	return new(PrintScale)
}

func (printScale *PrintScale) Parse(file *os.File) {
	printScale.parseStyle(file)
	printScale.XLocation = fsutil.ReadFloat(file)
	printScale.YLocation = fsutil.ReadFloat(file)
	printScale.Scale = fsutil.ReadFloat(file)
}

func (printscale *PrintScale) parseStyle(file *os.File) {
	style := fsutil.ReadBytesShort(file)
	switch style {
	case 0:
		printscale.Style = "centered"
	case 1:
		printscale.Style = "size to fit"
	case 2:
		printscale.Style = "user defined"
	}
	return
}
