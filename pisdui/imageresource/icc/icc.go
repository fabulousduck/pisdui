package icc

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/icc/header"
)

type ICCProfile struct {
	Header  *header.Header
	Taglist *TagList
}

type TagList struct {
	Count uint32
	Tags  []*Tag
}

type Tag struct {
	Sig    string
	Offset uint32
	Size   uint32
}

func (iccProfile *ICCProfile) GetTypeID() int {
	return 1039
}

func NewICCProfile() *ICCProfile {
	return new(ICCProfile)
}

func (iccProfile *ICCProfile) Parse(file *os.File) {

}
