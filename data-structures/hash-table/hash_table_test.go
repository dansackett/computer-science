package hash

import (
	"fmt"
	"testing"
)

func TestCanInitialize(t *testing.T) {
	ht := NewHashTable()

	if ht.NumEntries != 0 {
		t.Errorf("Initial hash table should have 0 entries, has %d", ht.NumEntries)
	}

	if ht.NumBuckets != 128 {
		t.Errorf("Initial hash table should have 128 buckets, has %d", ht.NumBuckets)
	}
}

func TestCanResizeItself(t *testing.T) {
	ht := NewHashTable()

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key_%d", i)
		ht.Put(key, i)
	}

	if ht.NumBuckets != 256 {
		t.Errorf("Resized hash table should have 256 buckets, has %d", ht.NumBuckets)
	}
}

func TestHashFunction(t *testing.T) {
	ht := NewHashTable()

	hash := ht.Hash("testing")

	if hash != 18433708127455886847 {
		t.Errorf("The hash function doesn't seem correct. Found %d", hash)
	}

	if hash%ht.NumBuckets != 127 {
		t.Errorf("The table index should be 127. Found %d", hash%ht.NumBuckets)
	}
}

func TestCanPutValue(t *testing.T) {
	ht := NewHashTable()
	val := "My Value"

	ht.Put("key", val)

	if ht.NumEntries != 1 {
		t.Error("Value not inserted correctly")
	}
}

func TestCanGetValue(t *testing.T) {
	ht := NewHashTable()
	val := "My Value"

	ht.Put("key", val)

	if ht.Get("key") != val {
		t.Errorf("Tried to get value '%s' but returned %s", val, ht.Get("key"))
	}
}

func TestCanUpdateValue(t *testing.T) {
	ht := NewHashTable()
	initialVal := "My Value"
	newVal := "NEW Value"

	ht.Put("key", initialVal)
	ht.Put("key", newVal)

	if ht.Get("key") != newVal {
		t.Errorf("Tried to get value '%s' but returned %s", newVal, ht.Get("key"))
	}

	if ht.NumEntries != 1 {
		t.Error("Update should not increase the entry count")
	}
}

func TestCanGetValueAfterResize(t *testing.T) {
	ht := NewHashTable()

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key_%d", i)
		ht.Put(key, i)
	}

	if ht.Get("key_50") != 50 {
		t.Errorf("key_50 should return %d after resize, returned %d", 50, ht.Get("key_50"))
	}
}

func TestCanDeleteValue(t *testing.T) {
	ht := NewHashTable()

	ht.Put("key", "testing")
	ht.Del("key")

	if ht.Get("key") != nil {
		t.Errorf("Deleted value should be empty, found %d", ht.Get("key"))
	}

	if ht.NumEntries != 0 {
		t.Error("Delete should decrease the entry count")
	}
}
