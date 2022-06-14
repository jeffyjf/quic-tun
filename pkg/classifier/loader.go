package classifier

import (
	"bytes"
	"io"
	"net"
	"strings"

	//	"github.com/kungze/quic-tun/pkg/restfulapi"
	"github.com/kungze/quic-tun/pkg/datastore"
	"github.com/lucas-clemente/quic-go"
)

var discriminators []DiscriminatorPlugin

func LoadDiscriminatorPlugin(plugins []string) {
	for _, plugin := range plugins {
		switch plugin = strings.ToLower(plugin); plugin {
		case "spice":
			discriminators = append(discriminators, &SpiceDiscriminator{})
		}
	}
}

func ProcessHeader(tunnelData *datastore.Tunnel, conn *net.Conn, stream *quic.Stream) error {
	header := bytes.Buffer{}
	io.CopyN(io.MultiWriter(&header, *stream), *conn, 64)
	for _, discr := range discriminators {
		res, err := discr.Identify(header.Bytes(), tunnelData)
		if err != nil {
			return err
		}
		if res {
			return nil
		}
	}
	return nil
}
