package cache

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"os/user"
	"path/filepath"
)

type Cache struct {
	handle *badger.DB
}

func NewCache(name string) (*Cache, error) {
	opts := badger.DefaultOptions(getCacheFolder(name))
	opts.Logger = nil
	db, err := badger.Open(opts)

	if err != nil {
		return nil, fmt.Errorf("db open: %v", err)
	}

	return &Cache{
		handle: db,
	}, nil
}

func (c *Cache) GetLatest(n int) (map[string]string, error) {
	result := make(map[string]string)

	err := c.handle.View(func(tx *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := tx.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			// Break out when we got the items we wanted
			if n == 0 {
				break
			}

			item := it.Item()
			key := item.Key()
			value, err := item.ValueCopy(nil)

			if err != nil {
				return err
			}

			result[string(key)] = string(value)
			n--
		}

		return nil
	})

	return result, err
}

func (c *Cache) Write(key, value string) error {
	return c.handle.Update(func(tx *badger.Txn) error {
		return tx.Set([]byte(key), []byte(value))
	})
}

func getCacheFolder(name string) string {
	currUser, _ := user.Current()
	folder := fmt.Sprintf("/.config/skyput/%s/", name)
	cacheDir := filepath.Join(currUser.HomeDir, folder)
	return cacheDir
}