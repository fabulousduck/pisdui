package descriptor

import (
	"fmt"
	"os"

	"github.com/pisdhooy/fsutil"
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
		break
	case "indx":
		fmt.Printf("idx index : ")
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
	enumReference.ClassIDName = fsutil.ParseUnicodeString(file)
	classIDlength := fsutil.ReadBytesLong(file)
	if classIDlength == 0 {
		enumReference.ClassID = fsutil.ReadBytesString(file, 4)
	} else {
		enumReference.ClassID = fsutil.ReadBytesString(file, int(classIDlength))
	}
	typeIDLength := fsutil.ReadBytesLong(file)
	if typeIDLength == 0 {
		enumReference.TypeID = fsutil.ReadBytesString(file, 4)
	} else {
		enumReference.TypeID = fsutil.ReadBytesString(file, int(typeIDLength))
	}
	enumLength := fsutil.ReadBytesLong(file)
	if enumLength == 0 {
		enumReference.Enum = fsutil.ReadBytesString(file, 4)
	} else {
		enumReference.Enum = fsutil.ReadBytesString(file, int(enumLength))
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
	name.Value = fsutil.ParseUnicodeString(file)
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
	property.ClassIDName = fsutil.ParseUnicodeString(file)
	classIDlength := fsutil.ReadBytesLong(file)
	if classIDlength == 0 {
		property.ClassID = fsutil.ReadBytesString(file, 4)
	} else {
		property.ClassID = fsutil.ReadBytesString(file, int(classIDlength))
	}
	keyIDLength := fsutil.ReadBytesLong(file)
	if keyIDLength == 0 {
		property.KeyID = fsutil.ReadBytesString(file, 4)
	} else {
		property.KeyID = fsutil.ReadBytesString(file, int(keyIDLength))
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
	class.ClassIDName = fsutil.ParseUnicodeString(file)
	classIDlength := fsutil.ReadBytesLong(file)
	if classIDlength == 0 {
		class.ClassID = fsutil.ReadBytesString(file, 4)
	} else {
		class.ClassID = fsutil.ReadBytesString(file, int(classIDlength))
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
	offset.ClassIDName = fsutil.ParseUnicodeString(file)
	classIDlength := fsutil.ReadBytesLong(file)
	if classIDlength == 0 {
		offset.ClassID = fsutil.ReadBytesString(file, 4)
	} else {
		offset.ClassID = fsutil.ReadBytesString(file, int(classIDlength))
	}

	offset.Value = fsutil.ReadBytesLong(file)
}
