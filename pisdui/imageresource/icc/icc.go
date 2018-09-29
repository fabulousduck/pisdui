package icc

import (
	"os"

	"github.com/davecgh/go-spew/spew"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type ICCProfile struct {
	Header  *Header
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
	header := NewHeader()
	tagList := NewTagList()

	header.Parse(file)
	iccProfile.Header = header
	tagList.Parse(file)
	iccProfile.Taglist = tagList
}

func NewTagList() *TagList {
	return new(TagList)
}

func (tagList *TagList) Parse(file *os.File) {
	tagList.Count = util.ReadBytesLong(file)
	spew.Dump(tagList.Count)
	for i := 0; i < int(tagList.Count); i++ {
		tag := NewTag()
		tag.Parse(file)
		tagList.Tags = append(tagList.Tags, tag)
	}
}

func NewTag() *Tag {
	return new(Tag)
}

func (tag *Tag) Parse(file *os.File) {
	tag.Sig = util.ReadBytesString(file, 4)
	tag.Offset = util.ReadBytesLong(file)
	tag.Size = util.ReadBytesLong(file)
}
