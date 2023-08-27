package db

import (
	"errors"
	"log"
	"os"
	"path"
	"time"

	"github.com/dgraph-io/badger/v4"
)

const DbBasePath = "./data"
const DbGCPeriodMin = 60

var dbs = make(map[string]*Database)

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

func Init() error {
	entries, err := os.ReadDir(DbBasePath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		dbPath := path.Join(DbBasePath, name)

		b, err := badger.Open(badger.DefaultOptions(dbPath))
		if err != nil {
			return err
		}

		dbs[name] = &Database{
			b: b,
		}
	}

	startGCRoutine()

	return nil
}

func Get(name string) (*Database, error) {
	db := dbs[name]
	if db == nil {
		_, err := Create(NewDbOptions{
			Name:     name,
			InMemory: true,
		})

		if err != nil {
			return nil, err
		}
	}

	return dbs[name], nil
}

func GetAll() map[string]DbInfo {
	result := make(map[string]DbInfo)

	for k, v := range dbs {
		lsm, _ := v.b.Size()
		options := v.b.Opts()

		result[k] = DbInfo{
			Lsm:      lsm,
			InMemory: options.InMemory,
		}
	}

	return result
}

func Drop(name string) error {
	db := dbs[name]
	if db == nil {
		return nil
	}

	dbDir := db.b.Opts().Dir

	// TODO: Block reads and writes
	err := db.b.DropAll()
	if err != nil {
		return err
	}

	err = db.b.Close()
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

func Create(options NewDbOptions) (*Database, error) {
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

	dbs[options.Name] = &Database{
		b: bdb,
	}

	return dbs[options.Name], nil
}

func startGCRoutine() {
	ticker := time.NewTicker(DbGCPeriodMin * time.Minute)

	go func() {
		for range ticker.C {
			for name, itm := range dbs {
				log.Printf("Running GC on database '%s'...", name)
				err := itm.b.RunValueLogGC(0.7)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}()
}
