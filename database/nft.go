package database

import "time"

type NftClass struct {
	Height int64
}

type Nft struct {
	Height int64
}

type NftSend struct {
	Sender   string
	Receiver string
	ClassId  string
	NftId    string
	Height   int64
}

func (db *Db) SaveNftSend(Sender, Receiver, ClassId, NftId, data string, Height int64, executedAt time.Time) error {
	return nil
}
