package db

import (
	"io"
	"path"

	"github.com/dgraph-io/badger/v4"
)

const DbBasePath = "./data"

var dbs = make(map[string]*db_wrapp)

type db_wrapp struct {
	badger *badger.DB
}

type DbStats struct {
	Lsm    int64
	Vlog   int64
	Tables []TableInfo
}

type TableInfo struct {
	Id            uint64
	KeyCount      uint32
	OnDiskSize    uint32
	StaleDataSize uint32
}

func (db *db_wrapp) Stats() DbStats {
	lsm, vlog := db.badger.Size()

	tables := make([]TableInfo, 0)
	for _, t := range db.badger.Tables() {
		tables = append(tables, TableInfo{
			Id:            t.ID,
			KeyCount:      t.KeyCount,
			OnDiskSize:    t.OnDiskSize,
			StaleDataSize: t.StaleDataSize,
		})
	}

	return DbStats{
		Lsm:    lsm,
		Vlog:   vlog,
		Tables: tables,
	}
}

func (db *db_wrapp) WriteValue(key string, w io.Writer) error {
	return db.badger.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		value, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		_, err = w.Write(value)

		return err
	})
}

func (db *db_wrapp) Set(key string, reader io.ReadCloser) error {
	return db.badger.Update(func(txn *badger.Txn) error {
		defer reader.Close()

		data, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		return txn.Set([]byte(key), data)
	})
}

func Get(name string) (*db_wrapp, error) {
	db := dbs[name]
	if db == nil {
		dbPath := path.Join(DbBasePath, name)

		bdb, err := badger.Open(badger.DefaultOptions(dbPath))
		if err != nil {
			return nil, err
		}

		dbs[name] = &db_wrapp{
			badger: bdb,
		}
	}

	return dbs[name], nil
}
