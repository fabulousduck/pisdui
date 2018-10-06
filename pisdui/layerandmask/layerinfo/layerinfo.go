package layerinfo

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord"
	"github.com/pisdhooy/fsutil"
)

//LayerInfo contains information about
//the layers in the .psd file
type LayerInfo struct {
	Length           uint32
	LayerCount       uint16
	LayerRecords     []*layerrecord.LayerRecord
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
	layerinfo.Length = fsutil.ReadBytesLong(file)
	layerinfo.LayerCount = fsutil.ReadBytesShort(file)
	for i := 0; i < int(layerinfo.LayerCount); i++ {
		layer := layerrecord.NewLayerRecord()
		layer.Parse(file)
		layerinfo.LayerRecords = append(layerinfo.LayerRecords, layer)
	}
}
