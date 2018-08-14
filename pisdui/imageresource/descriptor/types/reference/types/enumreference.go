package imageresource

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type enumReference struct {
	classIDName string
	classID     string
	typeID      string
	enum        string
}

func (er enumReference) getOsKeyBlockID() string {
	return "Enmr"
}

func parseEnumReference(file *os.File) enumReference {
	reference := new(enumReference)
	reference.classIDName = util.ParseUnicodeString(file)
	classIDlength := util.ReadBytesLong(file)
	if classIDlength == 0 {
		reference.classID = util.ReadBytesString(file, 4)
	} else {
		reference.classID = util.ReadBytesString(file, int(classIDlength))
	}
	typeIDLength := util.ReadBytesLong(file)
	if typeIDLength == 0 {
		reference.typeID = util.ReadBytesString(file, 4)
	} else {
		reference.typeID = util.ReadBytesString(file, int(typeIDLength))
	}
	enumLength := util.ReadBytesLong(file)
	if enumLength == 0 {
		reference.enum = util.ReadBytesString(file, 4)
	} else {
		reference.enum = util.ReadBytesString(file, int(enumLength))
	}
	return *reference
}
