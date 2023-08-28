package db

import (
	"context"
	"encoding/binary"
	"io"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto/z"
)

type Database struct {
	b *badger.DB
}

type DbStats struct {
	Lsm           int64
	Vlog          int64
	InMemory      bool
	KeyCount      uint32
	Size          int64
	OnDiskSize    uint32
	StaleDataSize uint32
	Metrics       string
}

type DbMetrics struct {
	KeysAdded uint64
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
	metrics := db.b.BlockCacheMetrics()

	stats := DbStats{
		Lsm:      lsm,
		Vlog:     vlog,
		InMemory: options.InMemory,
		Metrics:  metrics.String(),
	}

	for _, t := range db.b.Tables() {
		stats.KeyCount += t.KeyCount
		stats.OnDiskSize += t.OnDiskSize
		stats.StaleDataSize += t.StaleDataSize
	}

	for _, l := range db.b.Levels() {
		stats.Size += l.Size
	}

	return stats
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

func (db *Database) GetByPrefix(prefix string, w io.Writer) error {
	stream := db.b.NewStream()

	stream.Prefix = []byte(prefix)
	stream.Send = func(buf *z.Buffer) error {
		list, err := badger.BufferToKVList(buf)
		if err != nil {
			return err
		}

		sizeBytes := make([]byte, 4)
		for _, kv := range list.Kv {
			binary.LittleEndian.PutUint32(sizeBytes, uint32(len(kv.Key)))
			w.Write(sizeBytes)
			w.Write(kv.Key)

			binary.LittleEndian.PutUint32(sizeBytes, uint32(len(kv.Value)))
			w.Write(sizeBytes)
			w.Write(kv.Value)
		}

		return nil
	}

	return stream.Orchestrate(context.TODO())
}
