package header

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

//HeaderCS7 is for both CS7 and CS8.
//it is denoted to 7 since that changes from 6, where 7 and 8 are the same
type HeaderCS7 struct {
	Version           uint32
	DescriptorVersion uint32
	Descriptor        *descriptor.Descriptor
}

func (headerCS7 HeaderCS7) GetHeaderVersion() uint16 {
	return 7
}

func NewCS7Header() *HeaderCS7 {
	return new(HeaderCS7)
}

func (header *HeaderCS7) Parse(file *os.File) {
	descriptorObject := descriptor.NewDescriptor()

	header.Version = 6
	header.DescriptorVersion = util.ReadBytesLong(file)
	descriptorObject.Parse(file)
	header.Descriptor = descriptorObject
}
