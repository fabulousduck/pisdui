package descriptor

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type referenceOsKeyBlock interface {
	getReferenceOsKeyBlockID() string
}

func NewReferenceOsKeyBlock() *referenceOsKeyBlock {
	return new(referenceOsKeyBlock)
}

func parseReferenceOsKeyBlock(file *os.File, osKeyID string) referenceOsKeyBlock {
	fmt.Println("parsing referenceOSKeyBlock : ", osKeyID)
	var r referenceOsKeyBlock
	switch osKeyID {
	case "prop":
		property := NewProperty()
		property.Parse(file)
		r = property
		break
	case "Clss":
		classObject := NewClass()
		classObject.Parse(file)
		r = classObject
		break
	case "Enmr":
		enumReferenceBlock := NewEnumReference()
		enumReferenceBlock.Parse(file)
		r = enumReferenceBlock
		break
	case "rele":
		offsetObject := NewOffset()
		offsetObject.Parse(file)
		r = offsetObject
		break
	case "Idnt":
		//TODO figure out by looking at hex map
		fmt.Printf("idnt index : ")
		spew.Dump(file.Seek(0, 1))
		break
	case "indx":
		fmt.Printf("idx index : ")
		spew.Dump(file.Seek(0, 1))
		//TODO figure out by looking at hex map
		break
	case "name":
		name := NewName()
		name.Parse(file)
		r = name
		break
	default:
		panic("undefined type")
	}

	return r
}

type EnumReference struct {
	ClassIDName string
	ClassID     string
	TypeID      string
	Enum        string
}

func NewEnumReference() *EnumReference {
	return new(EnumReference)
}

func (_ EnumReference) getReferenceOsKeyBlockID() string {
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

type Name struct {
	Value string
}

func (n Name) getReferenceOsKeyBlockID() string {
	return "name"
}

func NewName() *Name {
	return new(Name)
}

func (name *Name) Parse(file *os.File) {
	name.Value = util.ParseUnicodeString(file)
}

type Property struct {
	ClassIDName string
	ClassID     string
	KeyID       string
}

func (p Property) getReferenceOsKeyBlockID() string {
	return "prop"
}

func NewProperty() *Property {
	return new(Property)
}

func (property *Property) Parse(file *os.File) {
	property.ClassIDName = util.ParseUnicodeString(file)
	classIDlength := util.ReadBytesLong(file)
	if classIDlength == 0 {
		property.ClassID = util.ReadBytesString(file, 4)
	} else {
		property.ClassID = util.ReadBytesString(file, int(classIDlength))
	}
	keyIDLength := util.ReadBytesLong(file)
	if keyIDLength == 0 {
		property.KeyID = util.ReadBytesString(file, 4)
	} else {
		property.KeyID = util.ReadBytesString(file, int(keyIDLength))
	}
}

type Class struct {
	ClassIDName string
	ClassID     string
}

func (c Class) getReferenceOsKeyBlockID() string {
	return "Clss"
}
func NewClass() *Class {
	return new(Class)
}

func (class *Class) Parse(file *os.File) {
	class.ClassIDName = util.ParseUnicodeString(file)
	classIDlength := util.ReadBytesLong(file)
	if classIDlength == 0 {
		class.ClassID = util.ReadBytesString(file, 4)
	} else {
		class.ClassID = util.ReadBytesString(file, int(classIDlength))
	}
}

type Offset struct {
	ClassIDName string
	ClassID     string
	Value       uint32
}

func (o Offset) getReferenceOsKeyBlockID() string {
	return "rele"
}

func NewOffset() *Offset {
	return new(Offset)
}

func (offset *Offset) Parse(file *os.File) {
	offset.ClassIDName = util.ParseUnicodeString(file)
	classIDlength := util.ReadBytesLong(file)
	if classIDlength == 0 {
		offset.ClassID = util.ReadBytesString(file, 4)
	} else {
		offset.ClassID = util.ReadBytesString(file, int(classIDlength))
	}

	offset.Value = util.ReadBytesLong(file)
}
