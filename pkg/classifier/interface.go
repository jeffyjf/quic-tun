package classifier

import (
	"github.com/kungze/quic-tun/pkg/datastore"
)

type DiscriminatorPlugin interface {
	Identify([]byte, *datastore.Tunnel) (bool, error)
}
