package header

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type HeaderInterface interface {
	GetHeaderVersion() uint16
}

func ParseHeader(file *os.File, version uint32) *HeaderInterface {
	var header HeaderInterface
	switch util.ReadBytesLong(file) {
	case 6:
		CS6Header := NewCS6Header()
		CS6Header.Parse(file)
		header = CS6Header
	case 7:
		fallthrough
	case 8:
		CS7Header := NewCS7Header()
		CS7Header.Parse(file)
		header = CS7Header
	}
	return &header
}
