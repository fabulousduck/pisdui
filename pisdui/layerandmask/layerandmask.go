package layerandmask

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	additionalinfo "github.com/fabulousduck/pisdui/pisdui/layerandmask/additionalInfo"
	"github.com/fabulousduck/pisdui/pisdui/layerandmask/layerinfo"
	"github.com/pisdhooy/fmtbytes"
)

//Data contains the photoshop files
//layers and their respective masks
type Data struct {
	Length                     uint32
	LayerInfo                  *layerinfo.LayerInfo
	GlobalLayerMaskInfo        *GlobalLayerMaskInfo
	AdditionalLayerInformation *additionalinfo.AdditionalInfo
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
	fmt.Println("LAYER INFO INDEX FP")
	spew.Dump(file.Seek(0, 1))
	layerandmaskobject.Length = fmtbytes.ReadBytesLong(file)

	if layerandmaskobject.Length%2 != 0 {
		layerandmaskobject.Length++
	}

	layerInfoObject := layerinfo.NewLayerInfo()
	layerInfoObject.Parse(file)
	layerandmaskobject.LayerInfo = layerInfoObject
	spew.Dump(layerandmaskobject)
}
