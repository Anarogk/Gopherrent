package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/url"

	bencode "github.com/jackpal/bencode-go"
)

const port uint16 = 6881

type TorrentFile struct {
	Announce     string
	InfoHash     [20]byte
	PieceHashes  [][20]byte
	PieceLengths int64
	Name         string
	Length       int64
}

// bencode structs
type bencodeInfo struct {
	Pieces       string `bencode:"pieces"`
	PiecesLength int64  `bencode:"piece length"`
	Length       int64  `bencode:"length"`
	Name         string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

type Peer struct {
	IP   net.IP
	Port uint16
}

// parse torrent file
func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

// announcing us as peers and getting a list of peers
func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}
	params := url.Values{
		"info_hash":  {fmt.Sprintf("%x", t.InfoHash)},
		"peer_id":    {fmt.Sprintf("%x", peerID)},
		"port":       {fmt.Sprintf("%d", port)},
		"uploaded":   {fmt.Sprintf("%d", t.PieceLengths)},
		"downloaded": {fmt.Sprintf("%d", t.PieceLengths)},
		"left":       {fmt.Sprintf("%d", t.PieceLengths)},
		"numwant":    {"20"},
		"compact":    {"1"},
		"no_peer_id": {"1"},
		"event":      {"started"},
	}
	base.RawQuery = params.Encode()
	return base.String(), nil
}

func Unmarshal(peersBin []byte) ([]Peer, error) {
	const peerSize = 6
	numPeers := len(peersBin) / peerSize
	if len(peersBin)%peerSize != 0 {
		return nil, fmt.Errorf("invalid peers length")
	}
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		peers[i].IP = net.IP(peersBin[i*peerSize : i*peerSize+4])
		peers[i].Port = binary.BigEndian.Uint16(peersBin[i*peerSize+4 : i*peerSize+6])
	}
	return peers, nil

}
