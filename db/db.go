package db

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/meeron/honey-badger/config"
)

const DbGCPeriodMin = 60

var (
	dbs      = make(map[string]*Database)
	gcTicker = time.NewTicker(24 * time.Hour) // Ticker will be reset for proper duration on init
)

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
	config := config.Get().Badger

	entries, err := os.ReadDir(config.DataDirPath)
	_, ok := err.(*fs.PathError)
	if ok {
		err = os.Mkdir(config.DataDirPath, os.ModeDir)
	}

	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		dbPath := path.Join(config.DataDirPath, name)

		b, err := badger.Open(badger.DefaultOptions(dbPath))
		if err != nil {
			return err
		}

		dbs[name] = &Database{
			b: b,
		}
	}

	startGCRoutine()
	notifySignal()

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
		config := config.Get().Badger

		dbPath := path.Join(config.DataDirPath, options.Name)

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
	gcTicker.Reset(DbGCPeriodMin * time.Minute)

	go func() {
		for range gcTicker.C {
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

func notifySignal() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalChannel
		log.Printf("%s", sig)

		gcTicker.Stop()
		log.Print("GC ticker closed")

		for name, db := range dbs {
			log.Printf("Closing database '%s'", name)
			if err := db.b.Close(); err != nil {
				log.Print(err)
			}
		}

		os.Exit(0)
	}()
}
