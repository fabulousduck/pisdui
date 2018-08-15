package enumreference

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type EnumReference struct {
	ClassIDName string
	ClassID     string
	TypeID      string
	Enum        string
}

func NewEnumReference() *EnumReference {
	return new(EnumReference)
}

func (enumreference EnumReference) getOsKeyBlockID() string {
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
