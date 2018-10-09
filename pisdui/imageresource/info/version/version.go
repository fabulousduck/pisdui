package version

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type Version struct {
	Key               uint32
	HasRealMergedData int
	WriterName        string
	ReaderName        string
	FileVersion       uint32
}

func (version *Version) GetTypeID() int {
	return 1057
}

func NewVersion() *Version {
	return new(Version)
}

func (version *Version) Parse(file *os.File) {
	version.Key = fmtbytes.ReadBytesLong(file)
	version.HasRealMergedData = fmtbytes.ReadSingleByte(file)
	version.WriterName = fmtbytes.ParseUnicodeString(file)
	version.ReaderName = fmtbytes.ParseUnicodeString(file)
	version.FileVersion = fmtbytes.ReadBytesLong(file)
}
