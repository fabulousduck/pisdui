package layerrecord

import (
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/channelinfo"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layerblendingranges"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layermaskdata"
)

//LayerRecord represents a layer in
//the photoshop file
type LayerRecord struct {
	Top                 uint32
	Left                uint32
	Bottom              uint32
	Right               uint32
	ChannelCount        uint16
	ChannelInfo         channelinfo.ChannelInfo
	BlendModeSignature  uint32
	BlendModeKey        string
	Opacity             int
	Clipping            int
	Flags               int
	ExtraFieldLength    uint32
	LayermaskData       layermaskdata.LayerMaskData
	LayerBlendingRanges layerblendingranges.LayerBlendingRanges
	LayerName           string
}
