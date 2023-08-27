package db

import (
	"errors"
	"io"
	"os"
	"path"
	"time"

	"github.com/dgraph-io/badger/v4"
)

const DbBasePath = "./data"

var dbs = make(map[string]*db_wrapp)

type db_wrapp struct {
	badger *badger.DB
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

type DbInfo struct {
	Lsm      int64
	InMemory bool
}

type NewDbOptions struct {
	Name     string
	InMemory bool
}

func (o NewDbOptions) Validate() error {
	if o.Name == "" {
		return errors.New("Name cannot be empty")
	}

	return nil
}

func (db *db_wrapp) Stats() DbStats {
	lsm, vlog := db.badger.Size()
	options := db.badger.Opts()

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
		Lsm:      lsm,
		Vlog:     vlog,
		InMemory: options.InMemory,
		Tables:   tables,
	}
}

func (db *db_wrapp) Get(key string) ([]byte, byte, error) {
	txn := db.badger.NewTransaction(false)
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

func (db *db_wrapp) Set(key string, reader io.ReadCloser, meta byte, ttl uint) error {
	return db.badger.Update(func(txn *badger.Txn) error {
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

func (db *db_wrapp) Sync() error {
	// Cannot sync in memory databases
	if db.badger.Opts().InMemory {
		return nil
	}

	return db.badger.Sync()
}

func (db *db_wrapp) DeleteByKey(key string) error {
	return db.badger.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (db *db_wrapp) DeleteByPrefix(prefix string) error {
	return db.badger.DropPrefix([]byte(prefix))
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
		opt := badger.DefaultOptions("").WithInMemory(true)

		bdb, err := badger.Open(opt)
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
		options := v.badger.Opts()

		result[k] = DbInfo{
			Lsm:      lsm,
			InMemory: options.InMemory,
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

func Create(options NewDbOptions) (*db_wrapp, error) {
	if dbs[options.Name] != nil {
		return nil, errors.New("Db already exists")
	}

	var opt badger.Options

	if !options.InMemory {
		dbPath := path.Join(DbBasePath, options.Name)

		opt = badger.DefaultOptions(dbPath)
	} else {
		opt = badger.DefaultOptions("").
			WithInMemory(options.InMemory)
	}

	bdb, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	dbs[options.Name] = &db_wrapp{
		badger: bdb,
	}

	return dbs[options.Name], nil
}
