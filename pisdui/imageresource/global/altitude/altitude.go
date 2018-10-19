package altitude

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Altitude struct {
	value uint32
}

func (altitude *Altitude) GetTypeID() int {
	return 1049
}

func NewAltitude() *Altitude {
	return new(Altitude)
}

func (altitude *Altitude) Parse(file *os.File) {
	altitude.value = fmtbytes.ReadBytesLong(file)
}
