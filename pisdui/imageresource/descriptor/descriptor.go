package descriptor

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type OsKeyBlock interface {
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
	osKeyBlock OsKeyBlock
}

func (descriptor *Descriptor) GetTypeID() int {
	return 1088
}

func (Descriptor Descriptor) getOsKeyBlockID() string {
	return "objc"
}

/*NewDescriptor creates a new descriptor struct*/
func NewDescriptor() *Descriptor {
	return &Descriptor{}
}

/*Parse parses data from a descriptor block in a PSD file into a premade descriptor*/
func (descriptor *Descriptor) Parse(file *os.File) {

	descriptor.UnicodeString = util.ParseUnicodeString(file)

	classIDLength := util.ReadBytesLong(file)

	if classIDLength == 0 {
		descriptor.ClassID = util.ReadBytesString(file, 4)
	} else {
		descriptor.ClassID = util.ReadBytesString(file, int(classIDLength))
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
		descriptorItem.key = util.ReadBytesString(file, int(length))
	}

	descriptorItem.osTypeKey = util.ReadBytesString(file, 4)
	descriptorItem.osKeyBlock = parseOsKeyType(file, descriptorItem.osTypeKey)
	descriptor.Items = append(descriptor.Items, descriptorItem)
}

func parseOsKeyType(file *os.File, osKeyID string) OsKeyBlock {
	spew.Dump(osKeyID)
	var r OsKeyBlock
	switch osKeyID {
	case "obj ":
		referenceObject := NewReference()
		referenceObject.Parse(file)
		break
	case "Objc":
		descriptorObject := NewDescriptor()
		descriptorObject.Parse(file)
		fmt.Printf("------------------------------------------\n")
		spew.Dump(descriptorObject)
		fmt.Printf("------------------------------------------\n")
		r = descriptorObject
		break
	case "VlLS":
		break
	case "doub":
		doubleObject := NewDouble()
		err := doubleObject.Parse(file)
		if err != nil {
			//return this error properly
			panic(err)
		}
		r = doubleObject
		break
	case "UntF":
		break
	case "TEXT":
		textObject := NewText()
		textObject.Parse(file)
		r = textObject
		break
	case "enum":
		enumObject := NewEnum()
		enumObject.Parse(file)
		r = enumObject
		break
	case "long":
		break
	case "comp":
		break
	case "bool":
		boolObject := NewBool()
		boolObject.Parse(file)
		r = boolObject
		break
	case "GlbO":
		break
	case "type": //type and GlbC are both of type class
		fallthrough
	case "GlbC":
		break
	case "alis":
		break
	case "tdta":
		break
	}
	spew.Dump(r)
	return r
}
