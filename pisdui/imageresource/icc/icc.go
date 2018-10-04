package icc

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type ICCProfile struct {
	Header   *Header
	TagTable *TagTable
	TagData  [][]byte
}

type TagTable struct {
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
	tagTable := NewTagList()

	header.Parse(file)
	iccProfile.Header = header
	tagTable.Parse(file)
	iccProfile.TagTable = tagTable
	// for i := 0; i < int(iccProfile.TagTable.Count); i++ {
	// 	buffer := util.ReadBytesNInt(file, iccProfile.TagTable.Tags[i].Size)
	// 	iccProfile.TagData = append(iccProfile.TagData, buffer)
	// }
}

func NewTagList() *TagTable {
	return new(TagTable)
}

func (tagTable *TagTable) Parse(file *os.File) {
	tagTable.Count = util.ReadBytesLong(file)
	for i := 0; i < int(tagTable.Count); i++ {
		tag := NewTag()
		tag.Parse(file)
		tagTable.Tags = append(tagTable.Tags, tag)
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
