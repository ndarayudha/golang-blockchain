package models

// Index is the position of the data record in the blockchain
// Timestamp is automatically determined and is the time the data is written
// BPM or beats per minute, is your pulse rate
// Hash is a SHA256 identifier representing this data record
// PrevHash is the SHA256 identifier of the previous record in the chain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}
