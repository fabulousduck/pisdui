package additionalInfo

import "os"

type AdditionalInfoBlock interface {
	getBlockID() string
}

type AdditionalInfo struct {
	Blocks *[]*AdditionalInfoBlock
}

func NewAdditionalInfo() *AdditionalInfo {
	return new(AdditionalInfo)
}

func (additionalInfo *AdditionalInfo) Parse(file *os.File) {

}
