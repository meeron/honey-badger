package db

import (
	"context"
	"testing"

	"github.com/meeron/honey-badger/config"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	db := getDb()

	t.Run("should return Hit=False if key is not found", func(t *testing.T) {
		_, hit, err := db.Get("test-key1")
		if err != nil {
			panic(err)
		}

		assert.False(t, hit)
	})

	t.Run("should return data if key is found", func(t *testing.T) {
		const key = "test-key2"
		db.Set(key, []byte{1, 3, 4}, 0)

		data, hit, err := db.Get(key)
		if err != nil {
			panic(err)
		}

		assert.True(t, hit)
		assert.NotEmpty(t, data)
	})
}

func TestAdd(t *testing.T) {
	db := getDb()

	t.Run("should set key", func(t *testing.T) {
		const key = "set-test-key"
		var data = []byte{1, 2, 3}

		if err := db.Set(key, data, 0); err != nil {
			panic(err)
		}

		dataRes, _, _ := db.Get(key)

		assert.EqualValues(t, data, dataRes)
	})
}

func TestDeleteByKey(t *testing.T) {
	db := getDb()

	t.Run("should delete value by key", func(t *testing.T) {
		const key = "test-key2"
		db.Set(key, []byte{1, 2, 3}, 0)

		if err := db.DeleteByKey(key); err != nil {
			panic(err)
		}

		_, hit, err := db.Get(key)

		assert.False(t, hit)
		assert.Nil(t, err)
	})
}

func TestGetByPrefix(t *testing.T) {
	db := getDb()

	t.Run("should return data by prefix as map", func(t *testing.T) {
		var (
			data1 = []byte{1, 2, 3}
			data2 = []byte{4, 5, 6}
		)
		var (
			key1 = "prefix-test-1"
			key2 = "prefix-test-2"
		)

		db.Set(key1, data1, 0)
		db.Set(key2, data2, 0)

		res, err := db.GetByPrefix(context.Background(), "prefix-")
		if err != nil {
			panic(err)
		}

		assert.EqualValues(t, data1, res[key1])
		assert.EqualValues(t, data2, res[key2])
	})
}

func TestDeleteByPrefix(t *testing.T) {
	db := getDb()

	t.Run("should delete data by prefix", func(t *testing.T) {
		var (
			data1 = []byte{1, 2, 3}
			data2 = []byte{4, 5, 6}
		)
		var (
			key1 = "deleteprefix-test-1"
			key2 = "deleteprefix-test-2"
		)

		db.Set(key1, data1, 0)
		db.Set(key2, data2, 0)

		if err := db.DeleteByPrefix("deleteprefix"); err != nil {
			panic(err)
		}

		res, _ := db.GetByPrefix(context.Background(), "deleteprefix")

		assert.Empty(t, res)
	})
}

func TestSetBatch(t *testing.T) {
	db := getDb()

	t.Run("should set batch entries", func(t *testing.T) {
		data := map[string][]byte{
			"batch-1": make([]byte, 1),
			"batch-2": make([]byte, 1),
			"batch-3": make([]byte, 1),
			"batch-4": make([]byte, 1),
			"batch-5": make([]byte, 1),
		}

		err := db.SetBatch(data)
		dbData, _ := db.GetByPrefix(context.TODO(), "batch-")

		assert.Nil(t, err)
		assert.Equal(t, len(data), len(dbData))
	})
}

func getDb() *Database {
	db, err := CreateCtx(config.BadgerConfig{}).GetDb("test")
	if err != nil {
		panic(err)

	}

	return db
}
