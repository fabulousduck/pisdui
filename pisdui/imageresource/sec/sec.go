package sec

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Sec struct {
	Digest []byte
}

func (sec *Sec) GetTypeID() int {
	return 1061
}

func NewSec() *Sec {
	return new(Sec)
}

func (sec *Sec) Parse(file *os.File) {
	sec.Digest = fmtbytes.ReadRawBytes(file, 16)
}
