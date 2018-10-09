package layerrecord

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/channelinfo"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layerblendingranges"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layermaskdata"
	"github.com/pisdhooy/fmtbytes"
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
	layerRecord.Top = fmtbytes.ReadBytesLong(file)
	layerRecord.Left = fmtbytes.ReadBytesLong(file)
	layerRecord.Bottom = fmtbytes.ReadBytesLong(file)
	layerRecord.Right = fmtbytes.ReadBytesLong(file)
	layerRecord.ChannelCount = fmtbytes.ReadBytesShort(file)

	layerChannelInfoObject := channelinfo.NewChannelInfo()
	layerChannelInfoObject.Parse(file)
	layerRecord.ChannelInfo = layerChannelInfoObject

	layerRecord.BlendModeSignature = fmtbytes.ReadBytesLong(file)
	layerRecord.BlendModeKey = fmtbytes.ReadBytesString(file, 4)
	layerRecord.Opacity = fmtbytes.ReadSingleByte(file)
	layerRecord.Clipping = fmtbytes.ReadSingleByte(file)
	layerRecord.Flags = fmtbytes.ReadSingleByte(file) //TODO: do this properly filler
	fmtbytes.ReadSingleByte(file)
	layerRecord.ExtraFieldLength = fmtbytes.ReadBytesLong(file)

	layerMaskDataObject := layermaskdata.NewLayerMaskData()
	layerMaskDataObject.Parse(file)
	layerRecord.LayermaskData = layerMaskDataObject
}
