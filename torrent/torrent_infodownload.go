package torrent

import (
	"github.com/ganqierwu/rain/internal/infodownloader"
	"github.com/ganqierwu/rain/internal/peerprotocol"
)

func (t *torrent) nextInfoDownload() *infodownloader.InfoDownloader {
	for pe := range t.peers {
		if _, ok := t.infoDownloaders[pe]; ok {
			continue
		}
		if pe.ExtensionHandshake == nil {
			continue
		}
		if pe.ExtensionHandshake.MetadataSize == 0 {
			continue
		}
		if pe.ExtensionHandshake.MetadataSize > int(t.session.config.MaxMetadataSize) {
			t.log.Debugf("metadata size larger than allowed: %d", pe.ExtensionHandshake.MetadataSize)
			continue
		}
		_, ok := pe.ExtensionHandshake.M[peerprotocol.ExtensionKeyMetadata]
		if !ok {
			continue
		}
		t.log.Debugln("downloading info from", pe.String())
		return infodownloader.New(pe)
	}
	return nil
}
