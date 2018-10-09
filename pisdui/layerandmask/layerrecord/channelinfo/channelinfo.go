package channelinfo

import (
	"os"

	"github.com/pisdhooy/fmtbytes"
)

type ChannelInfo struct {
	ID     uint16
	Length uint32
}

func NewChannelInfo() *ChannelInfo {
	return new(ChannelInfo)
}

func (channelInfo *ChannelInfo) Parse(file *os.File) {
	channelInfo.ID = fmtbytes.ReadBytesShort(file)
	channelInfo.Length = fmtbytes.ReadBytesLong(file)
}
