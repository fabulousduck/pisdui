package id

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type ID struct {
	value uint32
}

func (id *ID) GetTypeID() int {
	return 1044
}

func NewID() *ID {
	return new(ID)
}

func (id *ID) Parse(file *os.File) {
	id.value = util.ReadBytesLong(file)
}
