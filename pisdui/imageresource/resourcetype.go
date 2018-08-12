package imageresource

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type parsedResourceBlock interface {
	getTypeId() int
}

type descriptor struct {
	version       uint32
	unicodeString string
	classID       string
	itemCount     uint32
	items         []descriptorItem
}

type descriptorItem struct {
}

func parseDescriptor(file *os.File) *descriptor {

	d := new(descriptor)
	d.version = util.ReadBytesLong(file)
	d.unicodeString = util.ParseUnicodeString(file)

	classIDLength := util.ReadBytesLong(file)

	fmt.Println("classIdLength : ", classIDLength)

	if classIDLength == 0 {
		d.classID = util.ReadBytesString(file, 4)
	} else {
		d.classID = util.ParseUnicodeString(file)
	}

	d.itemCount = util.ReadBytesLong(file)
	spew.Dump(d)

	return d
}
