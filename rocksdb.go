package gorocksdb

import "gorocksdb/wrapper"

type RocksDB interface{}

type rocksDB struct {
	db *wrapper.DB
}

func New() (RocksDB, error) {
	db, err := wrapper.OpenDb(nil, "123")
	if err != nil {
		return nil, err
	}

	return &rocksDB{
		db: db,
	}, nil
}
