package imageresource

import (
	"os"

	"github.com/fabulousduck/pisdui/pisdui/util"
)

type name struct {
	name string
}

func (n name) getOsKeyBlockID() string {
	return "name"
}

func parseName(file *os.File) osKeyBlock {
	n := new(name)
	n.name = util.ParseUnicodeString(file)
	return *n
}
