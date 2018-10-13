package exif

import (
	"os"

	gexif "github.com/rwcarlsen/goexif/exif"
)

type Exif struct {
	data *gexif.Exif
}

func (exif *Exif) GetTypeID() int {
	return 1058
}

func NewExif() *Exif {
	return new(Exif)
}

func (exif *Exif) Parse(file *os.File, size uint32) {
	tmpFP, _ := file.Seek(0, 1)
	e, err := gexif.Decode(file)
	if err != nil {
		//TODO make this return an error instead of panicing
		panic(err)
	}
	exif.data = e
	file.Seek(tmpFP+int64(size), 0)
}
