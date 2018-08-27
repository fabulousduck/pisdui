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
