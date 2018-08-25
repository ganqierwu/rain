package peermanager

import (
	"github.com/cenkalti/rain/internal/logger"
	"github.com/cenkalti/rain/internal/peer"
	"github.com/cenkalti/rain/internal/peerlist"
	"github.com/cenkalti/rain/internal/peermanager/acceptor"
	"github.com/cenkalti/rain/internal/peermanager/dialer"
	"github.com/cenkalti/rain/internal/peermanager/peerids"
	"github.com/cenkalti/rain/internal/torrentdata"
	"github.com/cenkalti/rain/internal/worker"
)

type PeerManager struct {
	port     int
	peerList *peerlist.PeerList
	peerIDs  *peerids.PeerIDs
	peerID   [20]byte
	infoHash [20]byte
	data     *torrentdata.Data
	messages *peer.Messages
	workers  worker.Workers
	log      logger.Logger
}

func New(port int, pl *peerlist.PeerList, peerID, infoHash [20]byte, d *torrentdata.Data, l logger.Logger) *PeerManager {
	return &PeerManager{
		port:     port,
		peerList: pl,
		peerIDs:  peerids.New(),
		peerID:   peerID,
		infoHash: infoHash,
		data:     d,
		messages: peer.NewMessages(),
		log:      l,
	}
}

func (m *PeerManager) PeerMessages() *peer.Messages {
	return m.messages
}

func (m *PeerManager) Run(stopC chan struct{}) {
	a := acceptor.New(m.port, m.peerIDs, m.peerID, m.infoHash, m.data, m.messages, m.log)
	m.workers.Start(a)

	d := dialer.New(m.peerList, m.peerIDs, m.peerID, m.infoHash, m.data, m.messages, m.log)
	m.workers.Start(d)

	<-stopC
	m.workers.Stop()
}