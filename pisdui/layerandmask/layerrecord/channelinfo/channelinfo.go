package channelinfo

import (
	"os"

	"github.com/pisdhooy/fsutil"
)

type ChannelInfo struct {
	ID     uint16
	Length uint32
}

func NewChannelInfo() *ChannelInfo {
	return new(ChannelInfo)
}

func (channelInfo *ChannelInfo) Parse(file *os.File) {
	channelInfo.ID = fsutil.ReadBytesShort(file)
	channelInfo.Length = fsutil.ReadBytesLong(file)
}
