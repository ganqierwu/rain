package session

import (
	"time"

	"github.com/cenkalti/rain/internal/clientversion"
)

// Config for Session.
type Config struct {
	// Database file to save resume data.
	Database string
	// DataDir is where files are downloaded.
	DataDir string
	// New torrents will be listened at selected port in this range.
	PortBegin, PortEnd uint16
	// DHT node will listen on this IP.
	DHTAddress string
	// DHT node will listen on this UDP port.
	DHTPort uint16
	// At start, client will set max open files limit to this number. (like "ulimit -n" command)
	MaxOpenFiles uint64
	// URL to the blocklist file in CIDR format.
	BlocklistURL string
	// When to refresh blocklist
	BlocklistUpdateInterval time.Duration
	// Host to listen for RPC server
	RPCHost string
	// Listen port for RPC server
	RPCPort int
	// Time to wait for ongoing requests before shutting down RPC HTTP server.
	RPCShutdownTimeout time.Duration

	// Number of unchoked peers.
	UnchokedPeers int
	// Number of optimistic unchoked peers.
	OptimisticUnchokedPeers int
	// Max number of blocks requested from a peer but not received yet
	RequestQueueLength int
	// Time to wait for a requested block to be received before marking peer as snubbed
	RequestTimeout time.Duration
	// Max number of running downloads on piece in endgame mode, snubbed and choed peers don't count
	EndgameParallelDownloadsPerPiece int
	// Max number of outgoing connections to dial
	MaxPeerDial int
	// Max number of incoming connections to accept
	MaxPeerAccept int
	// Running piece downloads, snubbed and choked peers don't count
	ParallelPieceDownloads int
	// Running metadata downloads, snubbed peers don't count
	ParallelMetadataDownloads int
	// Time to wait for TCP connection to open.
	PeerConnectTimeout time.Duration
	// Time to wait for BitTorrent handshake to complete.
	PeerHandshakeTimeout time.Duration
	// When peer has started to send piece block, if it does not send any bytes in PieceTimeout, the connection is closed.
	PieceTimeout time.Duration
	// Buffer size for messages read from a single peer
	PeerReadBufferSize int
	// When the client want to connect a peer, first it tries to do encrypted handshake.
	// If it does not work, it connects to same peer again and does unencrypted handshake.
	// This behavior can be changed via this variable.
	DisableOutgoingEncryption bool
	// Dial only encrypted connections.
	ForceOutgoingEncryption bool
	// Do not accept unencrypted connections.
	ForceIncomingEncryption bool
	// Number of peer addresses to request in announce request.
	TrackerNumWant int
	// Time to wait for announcing stopped event.
	// Stopped event is sent to the tracker when torrent is stopped.
	TrackerStopTimeout time.Duration
	// When the client needs new peer addresses to connect, it ask to the tracker.
	// To prevent spamming the tracker an interval is set to wait before the next announce.
	MinAnnounceInterval time.Duration
	// Total time to wait for response to be read.
	// This includes ConnectTimeout and TLSHandshakeTimeout.
	HTTPTrackerTimeout time.Duration
	// User agent sent when communicating with HTTP trackers.
	HTTPTrackerUserAgent string
	// DHT announce interval
	DHTAnnounceInterval time.Duration
	// Minimum announce interval when announcing to DHT.
	MinDHTAnnounceInterval time.Duration
	// Bitfield is saved to disk for fast resume without hash checking.
	// There is an interval to keep IO lower.
	BitfieldWriteInterval time.Duration
	// Stats are written at interval to reduce write operations.
	StatsWriteInterval time.Duration
	// Number of bytes to read when a piece is requested by a peer.
	PieceReadSize int64
	// Number of cached bytes for piece read requests.
	PieceCacheSize int64
	// Read bytes for a piece part expires after duration.
	PieceCacheTTL time.Duration
	// Max number of peer addresses to keep in connect queue.
	MaxPeerAddresses int
	// Client version that is sent in BEP 10 handshake message.
	ExtensionHandshakeClientVersion string
	// Peer id is prefixed with this string. See BEP 20. Remaining bytes of peer id will be randomized.
	PeerIDPrefix string
}

var DefaultConfig = Config{
	Database:           "~/rain/session.db",
	DataDir:            "~/rain/data",
	PortBegin:          50000,
	PortEnd:            60000,
	DHTAddress:         "0.0.0.0",
	DHTPort:            7246,
	MaxOpenFiles:       1024 * 1024,
	RPCHost:            "127.0.0.1",
	RPCPort:            7246,
	RPCShutdownTimeout: 5 * time.Second,

	UnchokedPeers:                    3,
	OptimisticUnchokedPeers:          1,
	RequestQueueLength:               50,
	RequestTimeout:                   20 * time.Second,
	EndgameParallelDownloadsPerPiece: 2,
	MaxPeerDial:                      20,
	MaxPeerAccept:                    20,
	ParallelPieceDownloads:           10,
	ParallelMetadataDownloads:        2,
	PeerConnectTimeout:               5 * time.Second,
	PeerHandshakeTimeout:             10 * time.Second,
	PieceTimeout:                     30 * time.Second,
	PeerReadBufferSize:               32 * 1024,
	TrackerNumWant:                   100,
	TrackerStopTimeout:               5 * time.Second,
	MinAnnounceInterval:              time.Minute,
	HTTPTrackerTimeout:               10 * time.Second,
	HTTPTrackerUserAgent:             "Rain/" + clientversion.Version,
	DHTAnnounceInterval:              30 * time.Minute,
	MinDHTAnnounceInterval:           time.Minute,
	BitfieldWriteInterval:            30 * time.Second,
	StatsWriteInterval:               30 * time.Second,
	PieceReadSize:                    256 * 1024,
	PieceCacheSize:                   50 * 256 * 1024,
	PieceCacheTTL:                    5 * time.Minute,
	MaxPeerAddresses:                 2000,
	ExtensionHandshakeClientVersion:  "Rain " + clientversion.Version,
	PeerIDPrefix:                     "-RN" + clientversion.Version + "-",
}