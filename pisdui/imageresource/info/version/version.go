package version

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	version.Key = fsutil.ReadBytesLong(file)
	version.HasRealMergedData = fsutil.ReadSingleByte(file)
	version.WriterName = fsutil.ParseUnicodeString(file)
	version.ReaderName = fsutil.ParseUnicodeString(file)
	version.FileVersion = fsutil.ReadBytesLong(file)
}
