package unicode

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type UnicodeString struct {
	value string
}

func (unicodeString *UnicodeString) GetTypeID() int {
	return 1045
}

func NewUnicodeString() *UnicodeString {
	return new(UnicodeString)
}

func (unicodeString *UnicodeString) Parse(file *os.File) {
	unicodeString.value = fmtbytes.ParseUnicodeString(file)
}
