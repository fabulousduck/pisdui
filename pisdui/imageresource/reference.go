package imageresource

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type reference struct {
	itemCount      uint32
	referenceItems []*referenceItem
}

type referenceItem struct {
	osTypeKey  string
	osKeyBlock osKeyBlock
}

func (rf reference) getOsKeyBlockID() string {
	return "obj "
}

func parseReference(file *os.File) *reference {
	r := new(reference)

	r.itemCount = util.ReadBytesLong(file)
	var i uint32
	for i = 0; i < r.itemCount; i++ {
		r.referenceItems = append(r.referenceItems, parseReferenceItem(file))
	}

	return r
}

func parseReferenceItem(file *os.File) *referenceItem {
	referenceItem := new(referenceItem)
	referenceItem.osTypeKey = util.ReadBytesString(file, 4)
	referenceItem.osKeyBlock = parseReferenceOsKeyType(file, referenceItem.osTypeKey)
	return referenceItem
}

func parseReferenceOsKeyType(file *os.File, osKeyID string) osKeyBlock {
	var r osKeyBlock
	switch osKeyID {
	case "prop":
		break
	case "Clss":
		break
	case "Enmr":
		r = parseEnumReference(file)
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
