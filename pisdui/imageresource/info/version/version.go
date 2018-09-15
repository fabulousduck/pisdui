package version

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
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
	version.Key = util.ReadBytesLong(file)
	version.HasRealMergedData = util.ReadSingleByte(file)
	version.WriterName = util.ParseUnicodeString(file)
	version.ReaderName = util.ParseUnicodeString(file)
	version.FileVersion = util.ReadBytesLong(file)
}
