package id

import (
	"os"

	"github.com/pisdhooy/fsutil"
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
	id.value = fsutil.ReadBytesLong(file)
}
