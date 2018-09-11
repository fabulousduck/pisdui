package layerrecord

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/channelinfo"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layerblendingranges"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layermaskdata"
	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

//LayerRecord represents a layer in
//the photoshop file
type LayerRecord struct {
	Top                 uint32
	Left                uint32
	Bottom              uint32
	Right               uint32
	ChannelCount        uint16
	ChannelInfo         *channelinfo.ChannelInfo
	BlendModeSignature  uint32
	BlendModeKey        string
	Opacity             int
	Clipping            int
	Flags               int
	ExtraFieldLength    uint32
	LayermaskData       *layermaskdata.LayerMaskData
	LayerBlendingRanges *layerblendingranges.LayerBlendingRanges
	LayerName           string
}

func NewLayerRecord() *LayerRecord {
	return new(LayerRecord)
}

func (layerRecord *LayerRecord) Parse(file *os.File) {
	layerRecord.Top = util.ReadBytesLong(file)
	layerRecord.Left = util.ReadBytesLong(file)
	layerRecord.Bottom = util.ReadBytesLong(file)
	layerRecord.Right = util.ReadBytesLong(file)
	layerRecord.ChannelCount = util.ReadBytesShort(file)

	layerChannelInfoObject := channelinfo.NewChannelInfo()
	layerChannelInfoObject.Parse(file)
	layerRecord.ChannelInfo = layerChannelInfoObject

	layerRecord.BlendModeSignature = util.ReadBytesLong(file)
	layerRecord.BlendModeKey = util.ReadBytesString(file, 4)
	layerRecord.Opacity = util.ReadSingleByte(file)
	layerRecord.Clipping = util.ReadSingleByte(file)
	layerRecord.Flags = util.ReadSingleByte(file) //TODO: do this properly
	//filler
	util.ReadSingleByte(file)
	layerRecord.ExtraFieldLength = util.ReadBytesLong(file)

	layerMaskDataObject := layermaskdata.NewLayerMaskData()
	layerMaskDataObject.Parse(file)
	layerRecord.LayermaskData = layerMaskDataObject
}
