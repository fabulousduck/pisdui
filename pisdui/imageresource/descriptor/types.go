//note: this entire file is a fuck fest of type aliases and nested types
//it makes no sense at all, but thats just how psd files work.

package descriptor

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type Reference struct {
	ItemCount      uint32
	ReferenceItems []*ReferenceItem
}

type ReferenceItem struct {
	OsTypeKey  string
	OsKeyBlock referenceOsKeyBlock
}

func (reference Reference) getOsKeyBlockID() string {
	return "obj "
}

func NewReference() *Reference {
	return new(Reference)
}

func (reference *Reference) Parse(file *os.File) {
	reference.ItemCount = util.ReadBytesLong(file)
	var i uint32
	for i = 0; i < reference.ItemCount; i++ {
		referenceItem := new(ReferenceItem)
		referenceItem.Parse(file)
		reference.ReferenceItems = append(reference.ReferenceItems, referenceItem)
	}
}

func (referenceItem *ReferenceItem) Parse(file *os.File) {
	referenceItem.OsTypeKey = util.ReadBytesString(file, 4)
	referenceItem.OsKeyBlock = parseReferenceOsKeyBlock(file, referenceItem.OsTypeKey)
}

type Bool struct {
	Value bool
}

func (Bool Bool) getOsKeyBlockID() string {
	return "bool"
}

func NewBool() *Bool {
	return new(Bool)
}

func (Bool *Bool) Parse(file *os.File) {
	Bool.Value = util.ReadSingleByte(file) == 1
}

type Enum struct {
	Type  string
	Value string
}

func (enum Enum) getOsKeyBlockID() string {
	return "enum"
}

func NewEnum() *Enum {
	return new(Enum)
}

func (enum *Enum) Parse(file *os.File) {
	typeLength := util.ReadBytesLong(file)
	if typeLength < 1 {
		enum.Type = util.ReadBytesString(file, 4)
	} else {
		enum.Type = util.ReadBytesString(file, int(typeLength))
	}

	enumLength := util.ReadBytesLong(file)
	if enumLength < 0 {
		enum.Value = util.ReadBytesString(file, 4)
	} else {
		enum.Value = util.ReadBytesString(file, int(enumLength))
	}

}
