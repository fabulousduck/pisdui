package layerrecord

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/channelinfo"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layerblendingranges"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerrecord/layermaskdata"
	"github.com/pisdhooy/fsutil"
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
	layerRecord.Top = fsutil.ReadBytesLong(file)
	layerRecord.Left = fsutil.ReadBytesLong(file)
	layerRecord.Bottom = fsutil.ReadBytesLong(file)
	layerRecord.Right = fsutil.ReadBytesLong(file)
	layerRecord.ChannelCount = fsutil.ReadBytesShort(file)

	layerChannelInfoObject := channelinfo.NewChannelInfo()
	layerChannelInfoObject.Parse(file)
	layerRecord.ChannelInfo = layerChannelInfoObject

	layerRecord.BlendModeSignature = fsutil.ReadBytesLong(file)
	layerRecord.BlendModeKey = fsutil.ReadBytesString(file, 4)
	layerRecord.Opacity = fsutil.ReadSingleByte(file)
	layerRecord.Clipping = fsutil.ReadSingleByte(file)
	layerRecord.Flags = fsutil.ReadSingleByte(file) //TODO: do this properly filler
	fsutil.ReadSingleByte(file)
	layerRecord.ExtraFieldLength = fsutil.ReadBytesLong(file)

	layerMaskDataObject := layermaskdata.NewLayerMaskData()
	layerMaskDataObject.Parse(file)
	layerRecord.LayermaskData = layerMaskDataObject
}
