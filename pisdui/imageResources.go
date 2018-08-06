package pisdui

type ImageResources struct {
	Length         uint32
	ResourceBlocks []ResourceBlock
}

type ResourceBlock struct {
	Signature    uint32
	Id           uint16
	PascalString string
	DataSize     uint32
	DataBlock    []byte
}

func (pd *Pisdui) ParseImageResources() {
	pd.PSD.ImageResources.Length = ReadBytesLong(pd.FileContents)

}
