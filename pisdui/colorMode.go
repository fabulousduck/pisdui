package pisdui

/*ColorModeData contains data related to the files colors mode

Only really important when the color mode is set to "Indexed" or "Duotone"
*/
type ColorModeData struct {
	Length uint32
	Data   uint32
}

/*ParseColorModeData interprets the colormode data in the file

Only really interesting when color mode in the header is either "Indexed" or "Duotone"
*/
func (pd *Pisdui) ParseColorModeData() {
	pd.PSD.ColorModeData.Length = ReadBytesLong(pd.FileContents)
	fileColorMode := pd.PSD.Header.colorMode
	switch fileColorMode {
	case "Indexed":
		pd.PSD.ColorModeData.parseIndexedColorMode()
		break
	case "Duotone":
		pd.PSD.ColorModeData.parseDuotoneColorMode()
		break
	default:
		pd.PSD.ColorModeData.parseDefaultColorMode()
	}
}

func (cmd *ColorModeData) parseIndexedColorMode() {

}

func (cmd *ColorModeData) parseDuotoneColorMode() {

}

func (cmd *ColorModeData) parseDefaultColorMode() {

}
