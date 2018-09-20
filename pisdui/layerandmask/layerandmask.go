package layerandmask

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerinfo"
	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

//Data contains the photoshop files
//layers and their respective masks
type Data struct {
	Length                     uint32
	LayerInfo                  *layerinfo.LayerInfo
	GlobalLayerMaskInfo        *GlobalLayerMaskInfo
	AdditionalLayerInformation *AdditionalLayerInformation
}

type GlobalLayerMaskInfo struct {
	Length            uint32
	OverlayColorSpace uint16
	ByteCompOne       uint16
	ByteCompTwo       uint16
	Opacity           uint16
	Kind              int
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

func NewData() *Data {
	return new(Data)
}

func (layerandmaskobject *Data) Parse(file *os.File) {
	spew.Dump(file.Seek(0, 1))
	layerandmaskobject.Length = util.ReadBytesLong(file)
	layerInfoObject := layerinfo.NewLayerInfo()
	layerInfoObject.Parse(file)
	layerandmaskobject.LayerInfo = layerInfoObject
}
