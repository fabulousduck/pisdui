package measurementscale

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"

	"github.com/fabulousduck/pisdui/pisdui/imageresource/descriptor"
)

type MeasurementScale struct {
	DescriptorVersion uint32
	Descriptor        *descriptor.Descriptor
}

func (measurementScale *MeasurementScale) GetTypeID() int {
	return 1074
}

func NewMeasurementScale() *MeasurementScale {
	return new(MeasurementScale)
}

func (measurementScale *MeasurementScale) Parse(file *os.File) {
	descriptorObject := descriptor.NewDescriptor()

	measurementScale.DescriptorVersion = util.ReadBytesLong(file)
	descriptorObject.Parse(file)
	measurementScale.Descriptor = descriptorObject
}
