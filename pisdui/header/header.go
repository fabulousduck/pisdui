package header

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

/*FileHeader contains all header information
about the psd file*/
type FileHeader struct {
	signature string
	version   uint16
	reserved  []byte
	channels  uint16
	height    uint32
	width     uint32
	depth     uint16
	colorMode string
}

/*NewFileHeader creates and returns a pointer
to a new FileHeader struct*/
func NewFileHeader() *FileHeader {
	return new(FileHeader)
}

/*Parse parses the contents of the photoshop
file header into a FileHeader struct*/
func (fh *FileHeader) Parse(file *os.File) {
	fh.readSignature(file)
	fh.readVersion(file)
	fh.readReserved(file)
	fh.readChannels(file)
	fh.readDimensions(file)
	fh.readDepth(file)
	fh.readColorMode(file)
}

func (fh *FileHeader) readSignature(file *os.File) {
	signature := util.ReadBytesString(file, 4)
	if signature != "8BPS" {
		panic("Invalid header signature. got-" + signature + "-Expected 8BPS")
	}
	fh.signature = signature
}

func (fh *FileHeader) readVersion(file *os.File) {
	version := util.ReadBytesShort(file)
	if version != 1 {
		panic("Invalid file version.")
	}
	fh.version = version
}

func (fh *FileHeader) readReserved(file *os.File) {
	reserved := util.ReadBytesNInt(file, 6)
	fh.reserved = reserved
}

func (fh *FileHeader) readChannels(file *os.File) {
	channels := util.ReadBytesShort(file)
	if channels < 1 || channels > 56 {
		panic("header channels out of range")
	}
	fh.channels = channels
}

func (fh *FileHeader) readDimensions(file *os.File) {
	height := util.ReadBytesLong(file)
	width := util.ReadBytesLong(file)
	if width < 1 || width > 30000 || height < 1 || height > 30000 {
		panic("invalid file dimensions")
	}

	fh.height = height
	fh.width = width
}

func (fh *FileHeader) readDepth(file *os.File) {
	depth := util.ReadBytesShort(file)
	fh.depth = depth
}

func (fh *FileHeader) readColorMode(file *os.File) {

	colorModes := map[uint16]string{
		0: "Bitmap",
		1: "Greyscale",
		2: "Indexed",
		3: "RGB",
		4: "CYMK",
		7: "Multichannel",
		8: "Duotone",
		9: "Lab",
	}

	colorMode := util.ReadBytesShort(file)
	fh.colorMode = colorModes[colorMode]
}
