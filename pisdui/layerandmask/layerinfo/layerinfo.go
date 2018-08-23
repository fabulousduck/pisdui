package layerinfo

import "github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord"

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
