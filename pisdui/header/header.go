package header

import (
	"errors"
	"os"

	"github.com/pisdhooy/fmtbytes"
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

/*NewData creates and returns a pointer
to a new FileHeader struct*/
func NewData() *Data {
	return new(Data)
}

/*Parse parses the contents of the photoshop
file header into a Data struct*/
func (fh *Data) Parse(file *os.File) []error {
	errors := []error{}
	errors = append(errors, fh.readSignature(file))
	errors = append(errors, fh.readVersion(file))
	fh.readReserved(file)
	errors = append(errors, fh.readChannels(file))
	errors = append(errors, fh.readDimensions(file))
	fh.readDepth(file)
	fh.readColorMode(file)

	return errors
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

func (fh *Data) readSignature(file *os.File) error {
	signature := fmtbytes.ReadBytesString(file, 4)
	if signature != "8BPS" {
		return errors.New("Invalid header signature. got-" + signature + "-Expected 8BPS")
	}
	fh.Signature = signature
	return nil
}

func (fh *Data) readVersion(file *os.File) error {
	version := fmtbytes.ReadBytesShort(file)
	if version != 1 {
		return errors.New("Invalid file version.")
	}
	fh.Version = version
	return nil
}

func (fh *Data) readReserved(file *os.File) {
	reserved := fmtbytes.ReadBytesNInt(file, 6)
	fh.Reserved = reserved
}

func (fh *Data) readChannels(file *os.File) error {
	channels := fmtbytes.ReadBytesShort(file)
	if channels < 1 || channels > 56 {
		return errors.New("header channels out of range")
	}
	fh.Channels = channels
	return nil
}

func (fh *Data) readDimensions(file *os.File) error {
	height := fmtbytes.ReadBytesLong(file)
	width := fmtbytes.ReadBytesLong(file)
	if width < 1 || width > 30000 || height < 1 || height > 30000 {
		return errors.New("invalid file dimensions")
	}

	fh.Height = height
	fh.Width = width
	return nil
}

func (fh *Data) readDepth(file *os.File) {
	depth := fmtbytes.ReadBytesShort(file)
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

	colorMode := fmtbytes.ReadBytesShort(file)
	fh.ColorMode = colorModes[colorMode]
}
