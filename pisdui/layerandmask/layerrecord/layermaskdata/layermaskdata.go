package layermaskdata

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	layerMaskData.Size = fsutil.ReadBytesLong(file)
	layerMaskData.Top = fsutil.ReadBytesLong(file)
	layerMaskData.Left = fsutil.ReadBytesLong(file)
	layerMaskData.Bottom = fsutil.ReadBytesLong(file)
	layerMaskData.Right = fsutil.ReadBytesLong(file)
	layerMaskData.DefaultColor = fsutil.ReadSingleByte(file)
	layerMaskData.MaskParameters = fsutil.ReadSingleByte(file)
}
