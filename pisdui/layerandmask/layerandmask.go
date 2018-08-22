package layerandmask

//Data contains the photoshop files
//layers and their respective masks
type Data struct {
	Length                     uint32
	LayerInfo                  LayerInfo
	GlobalLayerMaskInfo        GlobalLayerMaskInfo
	AdditionalLayerInformation AdditionalLayerInformation
}

//LayerInfo contains information about
//the layers in the .psd file
type LayerInfo struct {
	Length           uint32
	LayerCount       uint16
	LayerRecords     []LayerRecord
	ChannelImageData ChannelImageData
}

//LayerRecord represents a layer in
//the photoshop file
type LayerRecord struct {
	Top                 uint32
	Left                uint32
	Bottom              uint32
	Right               uint32
	ChannelCount        uint16
	ChannelInfo         ChannelInfo
	BlendModeSignature  uint32
	BlendModeKey        string
	Opacity             int
	Clipping            int
	Flags               int
	ExtraFieldLength    uint32
	LayermaskData       LayerMaskData
	LayerBlendingRanges LayerBlendingRanges
	LayerName           string
}

type ChannelInfo struct {
	ID     uint16
	Length uint32
}

type GlobalLayerMaskInfo struct {
	Length            uint32
	OverlayColorSpace uint16
	ByteCompOne       uint16
	ByteCompTwo       uint16
	Opacity           uint16
	Kind              int
}

type LayerMaskData struct {
	Size              uint32
	Top               uint32
	Left              uint32
	Bottom            uint32
	Right             int32
	DefaultColor      int
	MaskParameters    int
	MaskParameterData []byte
	RealFlags         []byte
	EnclosingTop      uint32
	EnclosingLeft     uint32
	EnclosingBottom   uint32
	EnclosingRight    uint32
}

type LayerBlendingRanges struct {
	Length                      uint32
	BlendSource                 []byte
	BlendSourceDestinationRange uint32
	BlendSourceRanges           []BlendSourceRange
}

type BlendSourceRange struct {
	Source      uint32
	Destination uint32
}

type AdditionalLayerInformation struct {
	Signature  string
	Key        string
	DataLength uint32
	Data       ADLIDataBlock
}

type ADLIDataBlock interface {
	GetTypeID() string
}

type ChannelImageData struct {
	Compression uint16
	ImageData   []byte
}
