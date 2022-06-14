package datastore

import (
	"sync"

	"github.com/google/uuid"
	"github.com/lucas-clemente/quic-go"
)

type Tunnel struct {
	Mutex              *sync.Mutex   `json:"-"`
	Uuid               uuid.UUID     `json:"uuid"`
	StreamID           quic.StreamID `json:"streamId"`
	ClientAppAddr      string        `json:"clientAppAddr"`
	ServerAppAddr      string        `json:"serverAppAddr"`
	RemoteEndpointAddr string        `json:"remoteEndpointAddr"`
	CreatedAt          string        `json:"createdAt"`
	Protocol           string        `json:"protocol"`
	ProtocolProperties any           `json:"protocolProperties"`
}

type tunDataStore struct {
	sync.Map
}

func (t *tunDataStore) LoadAllTunnels(tunnels *[]Tunnel) {
	TunDataStore.Range(func(key, value any) bool {
		*tunnels = append(*tunnels, value.(Tunnel))
		return true
	})
}

var TunDataStore = tunDataStore{}

// var tunnelDataStore map[uuid.UUID]Tunnel

// func AddTunnel(tunnel Tunnel) {
// 	tunnelDataStore[tunnel.Uuid] = tunnel
// }

// func RemoveTunnel(id uuid.UUID) {
// 	delete(tunnelDataStore, id)
// }

// func GetAllTunnels() {

// }

// func GetTunnlByID() {}
