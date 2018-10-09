package colormode

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

/*Data contains data related to the files colors mode

Only really important when the color mode is set to "Indexed" or "Duotone"
*/
type Data struct {
	Length      uint32
	Data        uint32
	Palette     []uint16
	DuotoneData []byte
}

/*NewData creates a new ColorMode struct
and returns a pointer to it*/
func NewData() *Data {
	return new(Data)
}

/*Parse interprets the colormode data in the file

Only really interesting when color mode in the header is either "Indexed" or "Duotone"
*/
func (cm *Data) Parse(file *os.File, colorMode string) {
	cm.Length = fmtbytes.ReadBytesLong(file)
	switch colorMode {
	case "Indexed":
		cm.parseIndexedColorMode(file)
		break
	case "Duotone":
		cm.parseDuotoneColorMode(file)
		break
	default:
		break
	}
}

func (cm *Data) parseIndexedColorMode(file *os.File) {
	palette := fmtbytes.ReadIntoArray16(file, cm.Length)
	cm.Palette = palette
}

func (cm *Data) parseDuotoneColorMode(file *os.File) {
	duotoneData := fmtbytes.ReadBytesNInt(file, cm.Length)
	cm.DuotoneData = duotoneData
}
