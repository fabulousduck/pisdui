package pisdui

/*ColorModeData contains data related to the files colors mode

Only really important when the color mode is set to "Indexed" or "Duotone"
*/
type ColorModeData struct {
	Length      uint32
	Data        uint32
	Palette     []uint16
	DuotoneData []byte
}

/*ParseColorModeData interprets the colormode data in the file

Only really interesting when color mode in the header is either "Indexed" or "Duotone"
*/
func (pd *Pisdui) ParseColorModeData() {
	pd.PSD.ColorModeData.Length = ReadBytesLong(pd.FileContents)
	fileColorMode := pd.PSD.Header.colorMode
	switch fileColorMode {
	case "Indexed":
		pd.parseIndexedColorMode()
		break
	case "Duotone":
		pd.parseDuotoneColorMode()
		break
	default:
		break
	}
}

func (pd *Pisdui) parseIndexedColorMode() {
	palette := ReadIntoArray16(pd.FileContents, pd.PSD.ColorModeData.Length)
	pd.PSD.ColorModeData.Palette = palette
}

func (pd *Pisdui) parseDuotoneColorMode() {
	duotoneData := ReadBytesNInt(pd.FileContents, pd.PSD.ColorModeData.Length)
	pd.PSD.ColorModeData.DuotoneData = duotoneData
}
