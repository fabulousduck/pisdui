package descriptor

import (
	"fmt"
	"github.com/fabulousduck/pisdui/pisdui/util"
	"os"
)

type Reference struct {
	ItemCount      uint32
	ReferenceItems []*ReferenceItem
}

type ReferenceItem struct {
	OsTypeKey  string
	OsKeyBlock referenceOsKeyBlock

	// OsKeyBlock figure out how to make OsKeyBlock play nice here
}

type referenceOsKeyBlock string

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
		enumReferenceBlock := NewEnumReference()
		enumReferenceBlock.Parse(file)
		r = referenceOsKeyBlock(enumReferenceBlock.TypeID)
		//Idk what you're trying to do here but you can't put a string or a struct on the same type variable-
		// so i replaced it with TypeID
		break
	case "rele":
		break
	case "Idnt":
		break
	case "indx":
		break
	case "name":
		r = referenceOsKeyBlock(parseName(file))
		break
	}
	return r
}

type Name string

func (n Name) getOsKeyBlockID() string {
	return "name"
}

func parseName(file *os.File) Name {
	pos, _ := file.Seek(0, 1)
	fmt.Println("idx before name parse : ", pos)
	n := Name(util.ParseUnicodeString(file))
	return n
}

type EnumReference struct {
	ClassIDName string
	ClassID     string
	TypeID      string
	Enum        string
}

func NewEnumReference() *EnumReference {
	return &EnumReference{}
}

func (_ EnumReference) getOsKeyBlockID() string {
	return "Enmr"
}

func (enumReference *EnumReference) Parse(file *os.File) {
	enumReference.ClassIDName = util.ParseUnicodeString(file)
	classIDlength := util.ReadBytesLong(file)
	if classIDlength == 0 {
		enumReference.ClassID = util.ReadBytesString(file, 4)
	} else {
		enumReference.ClassID = util.ReadBytesString(file, int(classIDlength))
	}
	typeIDLength := util.ReadBytesLong(file)
	if typeIDLength == 0 {
		enumReference.TypeID = util.ReadBytesString(file, 4)
	} else {
		enumReference.TypeID = util.ReadBytesString(file, int(typeIDLength))
	}
	enumLength := util.ReadBytesLong(file)
	if enumLength == 0 {
		enumReference.Enum = util.ReadBytesString(file, 4)
	} else {
		enumReference.Enum = util.ReadBytesString(file, int(enumLength))
	}
}
