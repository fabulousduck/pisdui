package layermaskdata

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
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
	layerMaskData.Size = util.ReadBytesLong(file)
	layerMaskData.Top = util.ReadBytesLong(file)
	layerMaskData.Left = util.ReadBytesLong(file)
	layerMaskData.Bottom = util.ReadBytesLong(file)
	layerMaskData.Right = util.ReadBytesLong(file)
	layerMaskData.DefaultColor = util.ReadSingleByte(file)
	layerMaskData.MaskParameters = util.ReadSingleByte(file)
}
