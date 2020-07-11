package cache

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/fatih/color"
	"os/user"
	"path/filepath"
	"strconv"
	"time"
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

func (c *Cache) Close() {
	c.handle.Close()
}

func (c *Cache) ShowLatest(n int) {
	latest, _ := c.GetLatest(n)

	for ts, value := range latest {
		tm := time.Unix(ts, 0)

		timeColor := color.New(color.FgMagenta)
		timeColor.Printf("[%s] ", tm)

		linkColor := color.New(color.FgCyan)
		linkColor.Printf("%s\n", value)
	}
}

func (c *Cache) GetLatest(n int) (map[int64]string, error) {
	result := make(map[int64]string)

	err := c.handle.View(func(tx *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		opts.Reverse = false
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

			ts, err := strconv.ParseInt(string(key), 10, 64)

			if err != nil {
				continue
			}

			result[ts] = string(value)
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
