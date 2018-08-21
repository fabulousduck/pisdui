package header

import (
	"errors"
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

/*Data contains all header information
about the psd file*/
type Data struct {
	Signature string
	Version   uint16
	Reserved  []byte
	Channels  uint16
	Height    uint32
	Width     uint32
	Depth     uint16
	ColorMode string
}

/*NewFileHeader creates and returns a pointer
to a new FileHeader struct*/
func NewData() *Data {
	return new(Data)
}

/*Parse parses the contents of the photoshop
file header into a Data struct*/
func (fh *Data) Parse(file *os.File) {
	fh.readSignature(file)
	fh.readVersion(file)
	fh.readReserved(file)
	fh.readChannels(file)
	fh.readDimensions(file)
	fh.readDepth(file)
	fh.readColorMode(file)
}

func (fh *Data) Save(file *os.File) error {
	err := fh.writeSignature(file)
	if err != nil {
		return err
	}
	return nil
}

func (fh *Data) writeSignature(fp *os.File) error {
	if len(fh.Signature) < 0 {
		return errors.New("trying to write null header signature")
	}
	fp.WriteAt([]byte(fh.Signature), 0)
	return nil
}

func (fh *Data) readSignature(file *os.File) {
	signature := util.ReadBytesString(file, 4)
	if signature != "8BPS" {
		panic("Invalid header signature. got-" + signature + "-Expected 8BPS")
	}
	fh.Signature = signature
}

func (fh *Data) readVersion(file *os.File) {
	version := util.ReadBytesShort(file)
	if version != 1 {
		panic("Invalid file version.")
	}
	fh.Version = version
}

func (fh *Data) readReserved(file *os.File) {
	reserved := util.ReadBytesNInt(file, 6)
	fh.Reserved = reserved
}

func (fh *Data) readChannels(file *os.File) {
	channels := util.ReadBytesShort(file)
	if channels < 1 || channels > 56 {
		panic("header channels out of range")
	}
	fh.Channels = channels
}

func (fh *Data) readDimensions(file *os.File) {
	height := util.ReadBytesLong(file)
	width := util.ReadBytesLong(file)
	if width < 1 || width > 30000 || height < 1 || height > 30000 {
		panic("invalid file dimensions")
	}

	fh.Height = height
	fh.Width = width
}

func (fh *Data) readDepth(file *os.File) {
	depth := util.ReadBytesShort(file)
	fh.Depth = depth
}

func (fh *Data) readColorMode(file *os.File) {

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
	fh.ColorMode = colorModes[colorMode]
}
