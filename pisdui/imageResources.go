package pisdui

import (
	"fmt"
)

type ImageResources struct {
	Length         uint32
	ResourceBlocks []ResourceBlock
}

type ResourceBlock struct {
	byteSize     uint32
	Signature    string
	Id           uint16
	PascalString string
	DataSize     uint32
	DataBlock    []byte
}

func (pd *Pisdui) ParseImageResources() {
	pd.PSD.ImageResources.Length = ReadBytesLong(pd.FileContents)
	var i uint32
	fmt.Println(pd.PSD.ImageResources.Length)
	for i = 0; i < pd.PSD.ImageResources.Length; {
		block := pd.ParseResourceBlock()
		pd.PSD.ImageResources.ResourceBlocks = append(pd.PSD.ImageResources.ResourceBlocks, *block)
		i += block.byteSize
		fmt.Printf("%+v\n", block)
	}
}

func (pd *Pisdui) ParseResourceBlock() *ResourceBlock {
	block := new(ResourceBlock)
	block.Signature = ReadBytesString(pd.FileContents, 4)
	block.Id = ReadBytesShort(pd.FileContents)
	pascalString, stringLength := pd.parsePascalString()
	block.PascalString = pascalString
	block.DataSize = ReadBytesLong(pd.FileContents)
	block.DataBlock = ReadBytesNInt(pd.FileContents, block.DataSize)
	block.byteSize = uint32(4 + 2 + stringLength + 4 + int(block.DataSize))
	return block
}

func (pd *Pisdui) parsePascalString() (string, int) {
	stringLength := ReadBytesShort(pd.FileContents)
	if stringLength%2 != 0 {
		stringLength++
	}
	if stringLength == 0 {
		return "", 2
	}
	pascalStringContents := ReadBytesString(pd.FileContents, int(stringLength))
	return pascalStringContents, len(pascalStringContents)
}
