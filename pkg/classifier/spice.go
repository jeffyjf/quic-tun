package classifier

import (
	"fmt"

	"github.com/kungze/quic-tun/pkg/datastore"
)

const (
	SPICE_MAGIC         = "REDQ"
	MAJOR_VERSION_INDEX = 4
	MINOR_VERSION_INDEX = 8
	CHANNEL_TYPE_INDEX  = 20
	CHANNEL_MAIN        = 1
	CHANNEL_DISPLAY     = 2
	CHANNEL_INPUTS      = 3
	CHANNEL_CURSOR      = 4
	CHANNEL_PLAYBACK    = 5
	CHANNEL_RECORD      = 6
	CHANNEL_TUNNEL      = 7
	CHANNEL_SMARTCARD   = 8
	CHANNEL_USBREDIR    = 9
	CHANNEL_PORT        = 10
	CHANNEL_WEBDAV      = 11
)

type spiceData struct {
	Version     string `json:"version"`
	SessionId   string `json:"sessionId"`
	ChannelType string `json:"channelType"`
	ServerName  string `json:"serverName"`
}

type SpiceDiscriminator struct{}

func (s *SpiceDiscriminator) Identify(header []byte, tunnelData *datastore.Tunnel) (bool, error) {

	if string(header[:4]) != SPICE_MAGIC {
		return false, nil
	}

	var spice = spiceData{
		Version:   fmt.Sprintf("%x.%x", header[MAJOR_VERSION_INDEX], header[MINOR_VERSION_INDEX]),
		SessionId: fmt.Sprintf("%x", header[16:20]),
	}

	tunnelData.Protocol = "spice"
	// tunnelData.ProtocolVersion = fmt.Sprintf("%x.%x", header[MAJOR_VERSION_INDEX], header[MINOR_VERSION_INDEX])
	// tunnelData.ProtocolSession = fmt.Sprintf("%x", header[16:20])

	switch header[CHANNEL_TYPE_INDEX] {
	case CHANNEL_MAIN:
		spice.ChannelType = "main"
	case CHANNEL_DISPLAY:
		spice.ChannelType = "display"
	case CHANNEL_INPUTS:
		spice.ChannelType = "inputs"
	case CHANNEL_CURSOR:
		spice.ChannelType = "cursor"
	case CHANNEL_PLAYBACK:
		spice.ChannelType = "playback"
	case CHANNEL_RECORD:
		spice.ChannelType = "record"
	case CHANNEL_TUNNEL:
		spice.ChannelType = "tunnel"
	case CHANNEL_SMARTCARD:
		spice.ChannelType = "smartcard"
	case CHANNEL_USBREDIR:
		spice.ChannelType = "usbredir"
	case CHANNEL_PORT:
		spice.ChannelType = "port"
	case CHANNEL_WEBDAV:
		spice.ChannelType = "webdev"
	default:
		spice.ChannelType = "unknow"
	}
	//	json_spice, _ := json.Marshal(spice)
	tunnelData.ProtocolProperties = spice
	return true, nil
}
