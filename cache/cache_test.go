package cache

import (
	"fmt"
	"testing"
)

const nrItems = 10

func TestCache_Write(t *testing.T) {
	cache, _ := NewCache("test")

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("key_%06d", i)
		err := cache.Write(key, "some_value")

		if err != nil {
			t.Error(err)
		}
	}

	latest, err := cache.GetLatest(nrItems)

	if err != nil {
		t.Error(err)
	}

	if len(latest) != nrItems {
		t.Errorf("Too many items recieved, wanted %d, got %d",
			nrItems, len(latest))
	}
}