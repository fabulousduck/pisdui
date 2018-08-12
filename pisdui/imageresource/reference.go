package imageresource

import "os"

type reference struct {
	itemCount uint32
}

type referenceItem struct {
	osTypeKey string
}

func (rf reference) getOsKeyBlockID() string {
	return "obj "
}

func parseReference(file *os.File) *reference {
	r := new(reference)
	return r
}
