package layerblendingranges

type LayerBlendingRanges struct {
	Length                      uint32
	BlendSource                 []byte
	BlendSourceDestinationRange uint32
	BlendSourceRanges           []BlendSourceRange
}

type BlendSourceRange struct {
	Source      uint32
	Destination uint32
}
