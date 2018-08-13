package imageresource

import (
	"fmt"
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
	pos, _ := file.Seek(0, 1)
	fmt.Println("idx before name parse : ", pos)
	n.name = util.ParseUnicodeString(file)
	return *n
}
