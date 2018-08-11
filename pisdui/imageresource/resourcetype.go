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
	unicodeString string
	classID       string
	itemCount     uint32
	items         []descriptorItem
}

type descriptorItem struct {
}

func parseDescriptor(file *os.File) *descriptor {
	fmt.Println("parsing descriptor")

	d := new(descriptor)
	d.unicodeString = fmt.Sprintf("%s", util.ParseUnicodeString(file))
	spew.Dump(d)
	classIdLength := util.ReadBytesLong(file)
	if classIdLength == 0 {
		d.classID = string(util.ReadBytesLong(file))
	} else {
		d.classID = util.ParseUnicodeString(file)
	}

	d.itemCount = util.ReadBytesLong(file)

	return d
}
