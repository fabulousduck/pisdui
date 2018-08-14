package descriptor

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor/types/reference"
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
	osKeyBlock osKeyBlock
}

func (descriptor descriptor) getTypeID() int {
	return 1088
}

/*NewDescriptor creates a new descriptor struct*/
func NewDescriptor() *Descriptor {
	return new(descriptor)
}

/*Parse parses data from a descriptor block in a PSD file into a premade descriptor*/
func (descriptor *Descriptor) Parse(file *os.File) {

	descriptor.version = util.ReadBytesLong(file)
	descriptor.unicodeString = util.ParseUnicodeString(file)

	classIDLength := util.ReadBytesLong(file)

	if classIDLength == 0 {
		descriptor.classID = util.ReadBytesString(file, 4)
	} else {
		descriptor.classID = util.ParseUnicodeString(file)
	}

	descriptor.itemCount = util.ReadBytesLong(file)

	var i uint32
	for i = 0; i < d.itemCount; i++ {
		descriptor.parseDescriptorItem(file)
	}

	spew.Dump(descriptor)
	pos, _ := file.Seek(0, 1)
	fmt.Println("pos after descriptor read: ", pos)
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

func parseOsKeyType(file *os.File, osKeyID string) osKeyBlock {

	var r osKeyBlock
	switch osKeyID {
	case "obj ":
		r = reference.parseReference(file)
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
