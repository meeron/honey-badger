package db

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/logger"
)

type DbContext struct {
	dbs      map[string]*Database
	gcTicker *time.Ticker
	config   config.BadgerConfig
	logger   *logger.Logger
}

func CreateCtx(c config.BadgerConfig) *DbContext {
	ctx := &DbContext{
		dbs:    make(map[string]*Database),
		config: c,
		logger: logger.Get("DbContext"),
	}

	return ctx
}

func (ctx *DbContext) LoadDbs() error {
	entries, err := os.ReadDir(ctx.config.DataDirPath)
	_, ok := err.(*fs.PathError)
	if ok {
		err = os.Mkdir(ctx.config.DataDirPath, 0777)
	}

	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		dbPath := path.Join(ctx.config.DataDirPath, name)

		opt := badger.DefaultOptions(dbPath).
			WithLogger(logger.Badger())

		b, err := badger.Open(opt)
		if err != nil {
			return err
		}

		ctx.dbs[name] = &Database{
			b: b,
		}
	}

	ctx.gcTicker = time.NewTicker(time.Duration(ctx.config.GCPeriodMin) * time.Minute)

	startGCRoutine(ctx)

	return nil
}

func (ctx *DbContext) GetDb(name string) (*Database, error) {
	db := ctx.dbs[name]
	if db == nil {
		_, err := ctx.CreateDb(name, true)

		if err != nil {
			return nil, err
		}
	}

	return ctx.dbs[name], nil
}

func (ctx *DbContext) DropDb(name string) error {
	db := ctx.dbs[name]
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

	delete(ctx.dbs, name)

	return nil
}

func (ctx *DbContext) CreateDb(name string, inMemory bool) (*Database, error) {
	if name == "" {
		return nil, errors.New("'name' cannot be empty")
	}

	if ctx.dbs[name] != nil {
		return nil, errors.New("Db already exists")
	}

	var opt badger.Options

	if !inMemory {
		config := config.Get().Badger

		dbPath := path.Join(config.DataDirPath, name)

		opt = badger.DefaultOptions(dbPath)
	} else {
		opt = badger.DefaultOptions("").
			WithInMemory(inMemory)
	}

	opt = opt.WithLogger(logger.Badger())

	bdb, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	ctx.dbs[name] = &Database{
		b: bdb,
	}

	return ctx.dbs[name], nil
}

func (ctx *DbContext) Close() {
	ctx.gcTicker.Stop()
	ctx.logger.Infof("GC ticker closed")

	for name, db := range ctx.dbs {
		ctx.logger.Infof("Closing database '%s'", name)
		if err := db.b.Close(); err != nil {
			ctx.logger.Error(err)
		}
	}
}

func startGCRoutine(ctx *DbContext) {
	period := time.Duration(ctx.config.GCPeriodMin) * time.Minute
	ctx.gcTicker.Reset(period)
	ctx.logger.Infof("GC tick set to: %v\n", period)

	go func() {
		for range ctx.gcTicker.C {
			for name, itm := range ctx.dbs {
				ctx.logger.Infof("Running GC on database '%s'...", name)
				err := itm.b.RunValueLogGC(0.7)
				if err != nil {
					ctx.logger.Error(err)
				}
			}
		}
	}()
}

/*
func notifySignal() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		log := logger.Get("db")

		sig := <-signalChannel
		log.Infof("%s", sig)

		gcTicker.Stop()
		log.Infof("GC ticker closed")

		for name, db := range dbs {
			log.Infof("Closing database '%s'", name)
			if err := db.b.Close(); err != nil {
				log.Error(err)
			}
		}

		os.Exit(0)
	}()
}*/
