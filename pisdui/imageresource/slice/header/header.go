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

func (headerCS6 HeaderCS6) GetHeaderVersion() uint16 {
	return 6
}

func (headerCS7 HeaderCS7) GetHeaderVersion() uint16 {
	return 7
}

/*
	castBackCS6 and castBackCS7 are done here to prevent cyclic
	imports in the header packages
*/
func CastBackCS6(headerInterface HeaderInterface) HeaderCS6 {
	return headerInterface.(HeaderCS6)
}

func CastBackCS7(headerInterface HeaderInterface) HeaderCS7 {
	return headerInterface.(HeaderCS7)
}
