package xmp

import (
	"os"

	"github.com/pisdhooy/fmtbytes"

	"trimmer.io/go-xmp/xmp"
)

type XMP struct {
	Values *xmp.Document
}

func (XMP *XMP) GetTypeID() int {
	return 1060
}

func NewXMP() *XMP {
	return new(XMP)
}

func (XMP *XMP) Parse(file *os.File, size uint32) {

	tmpFP, _ := file.Seek(0, 1)
	document, err := xmp.Read(file)
	if err != nil {
		panic(err)
	}

	file.Seek(tmpFP, 1)
	data := fmtbytes.ReadRawBytes(file, int(size))

	newOffset := int(tmpFP) + len(data)
	file.Seek(int64(newOffset), 0)
	XMP.Values = document

}
