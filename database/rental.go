package database

import "time"

func (db *Db) SaveDeployNft(contractOwner, name, symbol, description, uri, jsonData,
	data string, height int64, timestamp time.Time) error {
	return nil
}

func (db *Db) SaveMintNft(contractOwner, classId, nftId, reciever, uri, jsonData,
	data string, height int64, timestamp time.Time) error {
	return nil
}

func (db *Db) SaveBurnNft(contractOwner, classId, nftId,
	data string, height int64, timestamp time.Time) error {
	return nil
}

func (db *Db) SaveRentMintNft(contractOwner, classId, nftId string, startDate, endDate int64,
	data string, height int64, timestamp time.Time) error {
	return nil
}

func (db *Db) SaveRentBurnNft(contractOwner, classId, nftId, sessionId,
	data string, height int64, timestamp time.Time) error {
	return nil
}

func (db *Db) SaveAccessNft(classId, nftId, renter,
	data string, height int64, timestamp time.Time) error {
	return nil
}

func (db *Db) SaveGiveAccessNft(classId, nftId, sessionId, renter, newRenter,
	data string, height int64, timestamp time.Time) error {
	return nil
}

func (db *Db) SaveSendRent(classId, nftId, sessionId, fromRenter, toRenter,
	data string, height int64, timestamp time.Time) error {
	return nil
}
