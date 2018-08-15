package reference

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor/types/enumreference"
	"github.com/fabulousduck/pisdui/pisdui/util"
)

type Reference struct {
	ItemCount      uint32
	ReferenceItems []*ReferenceItem
}

type ReferenceItem struct {
	OsTypeKey string
	// OsKeyBlock figure out how to make OsKeyBlock play nice here
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

func (reference Reference) getOsKeyBlockID() string {
	return "obj "
}

func (referenceItem *ReferenceItem) Parse(file *os.File) {
	referenceItem.OsTypeKey = util.ReadBytesString(file, 4)
	referenceItem.OsKeyBlock = parseReferenceOsKeyType(file, referenceItem.OsTypeKey)
}

func parseReferenceOsKeyType(file *os.File, osKeyID string) referenceOsKeyBlock {
	var r referenceOsKeyBlock
	switch osKeyID {
	case "prop":
		break
	case "Clss":
		break
	case "Enmr":
		enumReferenceBlock := enumreference.NewEnumReference()
		enumReferenceBlock.Parse(file)
		r = enumReferenceBlock
		break
	case "rele":
		break
	case "Idnt":
		break
	case "indx":
		break
	case "name":
		r = parseName(file)
		break
	}
	return r
}
