package pisdui

import "fmt"

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
	// fmt.Println(pd.PSD.ImageResources.Length)
	for i = 0; i < pd.PSD.ImageResources.Length; {
		block := pd.ParseResourceBlock()
		pd.PSD.ImageResources.ResourceBlocks = append(pd.PSD.ImageResources.ResourceBlocks, *block)
		i += block.byteSize
		fmt.Println("read ", i, " out of ", pd.PSD.ImageResources.Length)
		// fmt.Printf("%+v\n", block)
	}
}

func (pd *Pisdui) ParseResourceBlock() *ResourceBlock {
	block := new(ResourceBlock)
	block.Signature = ReadBytesString(pd.FileContents, 4)
	p, _ := pd.FileContents.Seek(0, 1)
	fmt.Println("sig : ", block.Signature, " pos after read: ", p)
	block.Id = ReadBytesShort(pd.FileContents)
	pascalString, stringLength := pd.parsePascalString()
	block.PascalString = pascalString
	block.DataSize = ReadBytesLong(pd.FileContents)
	block.DataBlock = ReadBytesNInt(pd.FileContents, block.DataSize)
	if block.DataSize%2 != 0 {
		// fmt.Println("odd data size: ", block.DataSize)
		ReadSingleByte(pd.FileContents)
	}
	block.byteSize = uint32(4 + 2 + stringLength + 4 + int(block.DataSize))
	return block
}

func (pd *Pisdui) parsePascalString() (string, int) {
	b := ReadSingleByte(pd.FileContents)
	if b == 0 {
		ReadSingleByte(pd.FileContents)
		return "", 1
	}

	s := ReadBytesString(pd.FileContents, b)

	if b%2 != 0 {
		ReadSingleByte(pd.FileContents)
	}
	return s, len(s)
}
