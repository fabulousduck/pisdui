package layerinfo

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord"
	"github.com/fabulousduck/pisdui/pisdui/util"
)

//LayerInfo contains information about
//the layers in the .psd file
type LayerInfo struct {
	Length           uint32
	LayerCount       uint16
	LayerRecords     []layerrecord.LayerRecord
	ChannelImageData ChannelImageData
}

type ChannelImageData struct {
	Compression uint16
	ImageData   []byte
}

func NewLayerInfo() *LayerInfo {
	return new(LayerInfo)
}

func (layerinfo *LayerInfo) Parse(file *os.File) {
	layerinfo.Length = util.ReadBytesLong(file)
	layerinfo.LayerCount = util.ReadBytesShort(file)
}
