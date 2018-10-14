package angle

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Angle struct {
	Value uint32
}

func (angle *Angle) GetTypeID() int {
	return 1037
}

func NewAngle() *Angle {
	return new(Angle)
}

func (angle *Angle) Parse(file *os.File) {
	angle.Value = fmtbytes.ReadBytesLong(file)
}
