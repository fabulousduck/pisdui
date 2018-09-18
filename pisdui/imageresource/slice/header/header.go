package header

import (
	"os"
)

type HeaderInterface interface {
	GetHeaderVersion() uint16
}

func ParseHeader(file *os.File, version uint32) *HeaderInterface {
	var header HeaderInterface
	switch version {
	case 6:
		CS6Header := NewCS6Header()
		CS6Header.Version = version
		CS6Header.Parse(file)
		header = CS6Header
		break
	case 7:
		fallthrough
	case 8:
		CS7Header := NewCS7Header()
		CS7Header.Version = version
		CS7Header.Parse(file)
		header = CS7Header
		break
	}
	return &header
}
