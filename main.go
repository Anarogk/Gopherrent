package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/anarogk/bit-tor-og/torrentfile"
)

func main() {
	// Open torrent file
	torrentFile, err := torrentfile.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening torrent file:", err)
		return
	}
	fmt.Println("Torrent file opened successfully")

	// Get peers from tracker
	peers, err := torrentFile.GetPeers()
	if err != nil {
		fmt.Println("Error getting peers from tracker:", err)
		return
	}
	fmt.Println("Peers obtained successfully")
	for _, peer := range peers {
		fmt.Printf("Peer: %s:%d\n", peer.IP, peer.Port)
	}

	// Connect to peers
	for _, peer := range peers {
		conn, err := net.Dial("tcp", peer.IP.String()+":"+fmt.Sprintf("%d", peer.Port))
		if err != nil {
			fmt.Println("Error connecting to peer:", err)
			return
		}
		fmt.Println("Connected to peer successfully")
		go func() {
			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("Error reading from peer:", err)
					return
				}
				fmt.Println("Received", n, "bytes from peer")
			}
		}()
		time.Sleep(time.Second)
	}
}
