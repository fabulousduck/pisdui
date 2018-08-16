package descriptor

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type osKeyBlock interface {
	getOsKeyBlockID() string
}

/*Descriptor block in an image resource*/
type Descriptor struct {
	Version       uint32
	UnicodeString string
	ClassID       string
	ItemCount     uint32
	Items         []*descriptorItem
}

type descriptorItem struct {
	key        string
	osTypeKey  string
	osKeyBlock referenceOsKeyBlock
}

func (descriptor *Descriptor) GetTypeID() int {
	return 1088
}

/*NewDescriptor creates a new descriptor struct*/
func NewDescriptor() *Descriptor {
	return &Descriptor{}
}

/*Parse parses data from a descriptor block in a PSD file into a premade descriptor*/
func (descriptor *Descriptor) Parse(file *os.File) {

	descriptor.Version = util.ReadBytesLong(file)
	descriptor.UnicodeString = util.ParseUnicodeString(file)

	classIDLength := util.ReadBytesLong(file)

	if classIDLength == 0 {
		descriptor.ClassID = util.ReadBytesString(file, 4)
	} else {
		descriptor.ClassID = util.ParseUnicodeString(file)
	}

	descriptor.ItemCount = util.ReadBytesLong(file)

	var i uint32
	for i = 0; i < descriptor.ItemCount; i++ {
		descriptor.parseDescriptorItem(file)
	}
}

func (descriptor *Descriptor) parseDescriptorItem(file *os.File) {
	descriptorItem := new(descriptorItem)
	length := util.ReadBytesLong(file)
	if length == 0 {
		descriptorItem.key = util.ReadBytesString(file, 4)
	} else {
		descriptorItem.key = util.ParseUnicodeString(file)
	}

	descriptorItem.osTypeKey = util.ReadBytesString(file, 4)
	descriptorItem.osKeyBlock = parseOsKeyType(file, descriptorItem.osTypeKey)
	descriptor.Items = append(descriptor.Items, descriptorItem)
}

func parseOsKeyType(file *os.File, osKeyID string) referenceOsKeyBlock {

	var r referenceOsKeyBlock
	switch osKeyID {
	case "obj ":
		r = parseReferenceOsKeyType(file, osKeyID)
		break
	case "Objc":
		break
	case "VlLS":
		break
	case "doub":
		break
	case "UntF":
		break
	case "TEXT":
		break
	case "enum":
		break
	case "long":
		break
	case "comp":
		break
	case "bool":
		break
	case "GlbO":
		break
	case "type": //type and GlbC are both of type class
	case "GlbC":
		break
	case "alis":
		break
	case "tdta":
		break
	}
	return r
}
