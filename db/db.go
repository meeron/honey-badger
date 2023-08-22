package db

import (
	"io"
	"os"
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

type DbInfo struct {
	Lsm int64
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

func Init() error {
	entries, err := os.ReadDir(DbBasePath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		dbPath := path.Join(DbBasePath, name)

		bdb, err := badger.Open(badger.DefaultOptions(dbPath))
		if err != nil {
			return err
		}

		dbs[name] = &db_wrapp{
			badger: bdb,
		}
	}

	return nil
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

func GetAll() map[string]DbInfo {
	result := make(map[string]DbInfo)

	for k, v := range dbs {
		lsm, _ := v.badger.Size()

		result[k] = DbInfo{
			Lsm: lsm,
		}
	}

	return result
}

func Drop(name string) error {
	dbWrapp := dbs[name]
	if dbWrapp == nil {
		return nil
	}

	dbDir := dbWrapp.badger.Opts().Dir

	// TODO: Block reads and writes
	err := dbWrapp.badger.DropAll()
	if err != nil {
		return err
	}

	err = dbWrapp.badger.Close()
	if err != nil {
		return err
	}

	err = os.RemoveAll(dbDir)
	if err != nil {
		return err
	}

	delete(dbs, name)

	return nil
}
