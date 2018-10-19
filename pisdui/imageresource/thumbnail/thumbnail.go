package thumbnail

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Thumbnail struct {
	Format         uint32
	Width          uint32
	Height         uint32
	WidthBytes     uint32
	TotalSize      uint32
	CompressedSize uint32
	BitsPerPixel   uint16
	NumberOfPlanes uint16
	JIFFData       []byte
}

func (thumbnail *Thumbnail) GetTypeID() int {
	return 1036
}

func NewThumbnail() *Thumbnail {
	return new(Thumbnail)
}

func (thumbnail *Thumbnail) Parse(file *os.File) {
	thumbnail.Format = fmtbytes.ReadBytesLong(file)
	thumbnail.Width = fmtbytes.ReadBytesLong(file)
	thumbnail.Height = fmtbytes.ReadBytesLong(file)
	thumbnail.WidthBytes = fmtbytes.ReadBytesLong(file)
	thumbnail.TotalSize = fmtbytes.ReadBytesLong(file)
	thumbnail.CompressedSize = fmtbytes.ReadBytesLong(file)
	thumbnail.BitsPerPixel = fmtbytes.ReadBytesShort(file)
	thumbnail.NumberOfPlanes = fmtbytes.ReadBytesShort(file)
	thumbnail.JIFFData = fmtbytes.ReadRawBytes(file, int(thumbnail.CompressedSize))
}
