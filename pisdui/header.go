package pisdui

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

func (interpreter *Pisdui) ParseHeader() {
	interpreter.readSignature()
	interpreter.readVersion()
	interpreter.readReserved()
	interpreter.readChannels()
	interpreter.readDimensions()
	interpreter.readDepth()
	interpreter.readColorMode()
}

func (pd *Pisdui) readSignature() {
	signature := ReadBytesString(pd.FileContents, 4)
	if signature != "8BPS" {
		panic("Invalid header signature. got-" + signature + "-Expected 8BPS")
	}
	pd.PSD.Header.signature = signature
}

func (pd *Pisdui) readVersion() {
	version := ReadBytesShort(pd.FileContents)
	if version != 1 {
		panic("Invalid file version.")
	}
	pd.PSD.Header.version = version
}

func (pd *Pisdui) readReserved() {
	reserved := ReadBytesNInt(pd.FileContents, 6)
	pd.PSD.Header.reserved = reserved
}

func (pd *Pisdui) readChannels() {
	channels := ReadBytesShort(pd.FileContents)
	if channels < 1 || channels > 56 {
		panic("header channels out of range")
	}
	pd.PSD.Header.channels = channels
}

func (pd *Pisdui) readDimensions() {
	height := ReadBytesLong(pd.FileContents)
	width := ReadBytesLong(pd.FileContents)
	if width < 1 || width > 30000 || height < 1 || height > 30000 {
		panic("invalid file dimensions")
	}

	pd.PSD.Header.height = height
	pd.PSD.Header.width = width
}

func (pd *Pisdui) readDepth() {
	depth := ReadBytesShort(pd.FileContents)
	pd.PSD.Header.depth = depth
}

func (pd *Pisdui) readColorMode() {

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

	colorMode := ReadBytesShort(pd.FileContents)
	pd.PSD.Header.colorMode = colorModes[colorMode]
}
