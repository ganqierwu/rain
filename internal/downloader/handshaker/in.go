package handshaker

import (
	"net"

	"github.com/cenkalti/rain/internal/bitfield"
	"github.com/cenkalti/rain/internal/btconn"
	"github.com/cenkalti/rain/internal/logger"
	"github.com/cenkalti/rain/internal/peer"
)

type IncomingHandshaker struct {
	conn        net.Conn
	bitfield    *bitfield.Bitfield
	peerID      [20]byte
	sKeyHash    [20]byte
	infoHash    [20]byte
	newPeers    chan *peer.Peer
	disconnectC chan net.Conn
	log         logger.Logger
}

func NewIncoming(conn net.Conn, peerID, sKeyHash, infoHash [20]byte, newPeers chan *peer.Peer, disconnectC chan net.Conn, l logger.Logger) *IncomingHandshaker {
	return &IncomingHandshaker{
		conn:        conn,
		peerID:      peerID,
		sKeyHash:    sKeyHash,
		infoHash:    infoHash,
		newPeers:    newPeers,
		disconnectC: disconnectC,
		log:         l,
	}
}

func (h *IncomingHandshaker) Run(stopC chan struct{}) {
	log := logger.New("peer <- " + h.conn.RemoteAddr().String())

	defer func() {
		_ = h.conn.Close()
		select {
		case h.disconnectC <- h.conn:
		case <-stopC:
		}
	}()

	// TODO get this from config
	encryptionForceIncoming := false

	// TODO get supported extensions from common place
	ourExtensions := [8]byte{}
	ourbf := bitfield.NewBytes(ourExtensions[:], 64)
	ourbf.Set(61) // Fast Extension (BEP 6)
	ourbf.Set(43) // Extension Protocol (BEP 10)

	encConn, cipher, peerExtensions, peerID, _, err := btconn.Accept(
		h.conn, h.getSKey, encryptionForceIncoming, h.checkInfoHash, ourExtensions, h.peerID)
	if err != nil {
		log.Error(err)
		_ = h.conn.Close()
		return
	}
	log.Infof("Connection accepted. (cipher=%s extensions=%x client=%q)", cipher, peerExtensions, peerID[:8])

	peerbf := bitfield.NewBytes(peerExtensions[:], 64)
	extensions := ourbf.And(peerbf)

	p := peer.New(encConn, peerID, extensions, log)
	select {
	case h.newPeers <- p:
		p.Run(stopC)
	case <-stopC:
	}
}

func (h *IncomingHandshaker) getSKey(sKeyHash [20]byte) []byte {
	if sKeyHash == h.sKeyHash {
		return h.infoHash[:]
	}
	return nil
}

func (h *IncomingHandshaker) checkInfoHash(infoHash [20]byte) bool {
	return infoHash == h.infoHash
}