package layermaskdata

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type LayerMaskData struct {
	Size              uint32
	Top               uint32
	Left              uint32
	Bottom            uint32
	Right             uint32
	DefaultColor      int
	MaskParameters    int
	MaskParameterData []byte
	RealFlags         []byte
	EnclosingTop      uint32
	EnclosingLeft     uint32
	EnclosingBottom   uint32
	EnclosingRight    uint32
}

func NewLayerMaskData() *LayerMaskData {
	return new(LayerMaskData)
}

func (layerMaskData *LayerMaskData) Parse(file *os.File) {
	layerMaskData.Size = fmtbytes.ReadBytesLong(file)
	layerMaskData.Top = fmtbytes.ReadBytesLong(file)
	layerMaskData.Left = fmtbytes.ReadBytesLong(file)
	layerMaskData.Bottom = fmtbytes.ReadBytesLong(file)
	layerMaskData.Right = fmtbytes.ReadBytesLong(file)
	layerMaskData.DefaultColor = fmtbytes.ReadSingleByte(file)
	layerMaskData.MaskParameters = fmtbytes.ReadSingleByte(file)
}
