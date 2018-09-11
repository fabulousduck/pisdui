package channelinfo

import (
	"os"

	util "github.com/fabulousduck/pisdui/pisdui/util/file"
)

type ChannelInfo struct {
	ID     uint16
	Length uint32
}

func NewChannelInfo() *ChannelInfo {
	return new(ChannelInfo)
}

func (channelInfo *ChannelInfo) Parse(file *os.File) {
	channelInfo.ID = util.ReadBytesShort(file)
	channelInfo.Length = util.ReadBytesLong(file)
}
