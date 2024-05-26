
# TorrentFile struct

```go
type TorrentFile struct {
	Announce     string
	InfoHash     [20]byte
	PieceHashes  [][20]byte
	PieceLengths int64
	Name         string
	Length       int64
}
```

## Fields

### Detailed struct definition 
```go
type TorrentInfo struct {
  Announce string `bencode:"announce"`
  Info      struct {
    Length  int    `bencode:"length"`
    Name     string `bencode:"name"`
    PieceLength int  `bencode:"piece length"`
    Pieces   []byte `bencode:"pieces"`
    Files    []struct { // Optional for multi-file torrents
      Length int    `bencode:"length"`
      Path   []string `bencode:"path"`
    } `bencode:"files"`
  } `bencode:"info"`
  // Info hash is not part of the actual torrent file, but calculated from info dict
  InfoHash [20]byte 
}
```

### Announce
```go
Announce string
```
The announce field points to the tracker of the content file(s) that we are uploading as a torrent. This is a ‘UDP tracker protocol’. Note that the above value of ‘announce’ is of the form udp://exampletracker.com:port. This syntax might change depending on different torrent clients and different tracker providers.


### InfoHash

```go
InfoHash [20]byte
```

### PieceHashes

```go
PieceHashes [][20]byte
```

### PieceLengths

```go
PieceLengths int64
```
Size of each individual piece the file is divided into (in bytes). This is typically 256KB or 1MB.

### Name

```go
Name string
```
Name of the file (or directory for multiple files).

### Length

```go
Length int64
```
 Total size of the file(s) being shared (in bytes).
```rust

```
