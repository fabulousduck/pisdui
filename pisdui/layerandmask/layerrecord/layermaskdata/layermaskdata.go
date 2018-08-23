package layermaskdata

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
