package xmp

import (
	"fmt"
	"os"

	"github.com/pisdhooy/fmtbytes"

	"github.com/davecgh/go-spew/spew"

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

func (XMP *XMP) Parse(file *os.File) {
	document, err := xmp.Read(file)
	if err != nil {
		panic(err)
	}
	fmtbytes.ReadRawBytes(file, 10112)
	fmt.Println("XMP FP POS")
	spew.Dump(file.Seek(0, 1))
	XMP.Values = document
}
