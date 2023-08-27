package db

import (
	"io"
	"time"

	"github.com/dgraph-io/badger/v4"
)

type Database struct {
	b *badger.DB
}

type DbStats struct {
	Lsm      int64
	Vlog     int64
	InMemory bool
	Tables   []TableInfo
}

type TableInfo struct {
	Id            uint64
	KeyCount      uint32
	OnDiskSize    uint32
	StaleDataSize uint32
}

func (db *Database) Get(key string) ([]byte, byte, error) {
	txn := db.b.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(key))
	if err != nil {
		return nil, 0, err
	}

	meta := item.UserMeta()
	value, err := item.ValueCopy(nil)

	if err != nil {
		return nil, 0, err
	}

	return value, meta, nil
}

func (db *Database) Stats() DbStats {
	lsm, vlog := db.b.Size()
	options := db.b.Opts()

	tables := make([]TableInfo, 0)
	for _, t := range db.b.Tables() {
		tables = append(tables, TableInfo{
			Id:            t.ID,
			KeyCount:      t.KeyCount,
			OnDiskSize:    t.OnDiskSize,
			StaleDataSize: t.StaleDataSize,
		})
	}

	return DbStats{
		Lsm:      lsm,
		Vlog:     vlog,
		InMemory: options.InMemory,
		Tables:   tables,
	}
}

func (db *Database) Set(key string, reader io.ReadCloser, meta byte, ttl uint) error {
	return db.b.Update(func(txn *badger.Txn) error {
		defer reader.Close()

		data, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		entry := badger.NewEntry([]byte(key), data).WithMeta(meta)

		if ttl > 0 {
			entry = entry.WithTTL(time.Duration(ttl) * time.Second)
		}

		return txn.SetEntry(entry)
	})
}

func (db *Database) Sync() error {
	// Cannot sync in memory databases
	if db.b.Opts().InMemory {
		return nil
	}

	return db.b.Sync()
}

func (db *Database) DeleteByKey(key string) error {
	return db.b.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (db *Database) DeleteByPrefix(prefix string) error {
	return db.b.DropPrefix([]byte(prefix))
}
