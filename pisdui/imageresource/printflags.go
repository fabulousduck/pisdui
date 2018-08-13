package imageresource

type printflags struct {
}

func (n printflags) getOsKeyBlockID() string {
	return "name"
}
