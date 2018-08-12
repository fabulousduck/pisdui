package imageresource

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/fabulousduck/pisdui/pisdui/util"
)

type osKeyBlock interface {
	getOsKeyBlockID() string
}

type descriptor struct {
	version       uint32
	unicodeString string
	classID       string
	itemCount     uint32
	items         []*descriptorItem
}

type descriptorItem struct {
	key        string
	osTypeKey  string
	osKeyBlock osKeyBlock
}

func parseDescriptor(file *os.File) *descriptor {

	d := new(descriptor)
	d.version = util.ReadBytesLong(file)
	d.unicodeString = util.ParseUnicodeString(file)

	classIDLength := util.ReadBytesLong(file)

	if classIDLength == 0 {
		d.classID = util.ReadBytesString(file, 4)
	} else {
		d.classID = util.ParseUnicodeString(file)
	}

	d.itemCount = util.ReadBytesLong(file)

	var i uint32
	for i = 0; i < d.itemCount; i++ {
		d.items = append(d.items, parseDescriptorItem(file))
	}

	spew.Dump(d)

	return d
}

func parseDescriptorItem(file *os.File) *descriptorItem {
	di := new(descriptorItem)
	length := util.ReadBytesLong(file)
	fmt.Println("length > ", length)
	if length == 0 {
		di.key = util.ReadBytesString(file, 4)
	} else {
		di.key = util.ParseUnicodeString(file)
	}

	di.osTypeKey = util.ReadBytesString(file, 4)
	di.osKeyBlock = parseOsKeyType(file, di.osTypeKey)
	return di
}

func parseOsKeyType(file *os.File, osKeyID string) osKeyBlock {

	var r osKeyBlock
	switch osKeyID {
	case "obj ":
		r = parseReference(file)
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
	case "type":
		break
	case "GlbC":
		break
	case "alis":
		break
	case "tdta":
		break
	}
	return r
}
