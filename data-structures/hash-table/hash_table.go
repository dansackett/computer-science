package hash

const (
	// MaxLoadFactor is the ratio of entries to buckets which will force a resize
	MaxLoadFactor float64 = .75
	// DefaultTableSize is set higher to avoid many resizes
	DefaultTableSize uint64 = 128
	// SizeIncrease is the resize multiple when determining the new size of an
	// updated hash table
	SizeIncrease uint64 = 2
	// FNVOffset is the constant used in the FNV hash function as seen on
	// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
	FNVOffset uint64 = 14695981039346656037
	// FNVPrime is the constant prime factor used in the FNV hash function
	FNVPrime uint64 = 1099511628211
)

// HashTable is the primary data structure for storing our hash table. We keep
// implicit counts of buckets and entries with buckets being singly linked
// lists for separate chaining.
type HashTable struct {
	Buckets                []*List
	NumEntries, NumBuckets uint64
}

// NewHashTable creates a hash table based on the default sizing constraints.
func NewHashTable() *HashTable {
	return &HashTable{
		Buckets:    make([]*List, DefaultTableSize),
		NumBuckets: DefaultTableSize,
		NumEntries: 0,
	}
}

// resize rehashes a hash table which has grown beyond its load factor
// bringing it back to a more efficient state.
func (t *HashTable) resize() {
	newBucketSize := t.NumBuckets * SizeIncrease
	ht := &HashTable{
		Buckets:    make([]*List, newBucketSize),
		NumBuckets: newBucketSize,
		NumEntries: 0,
	}
	for _, list := range t.Buckets {
		if list != nil {
			for e := range list.Iter() {
				ht.Put(e.key, e.val)
			}
		}
	}
	*t = *ht
}

// LoadFactor returns the current load of the table.
func (t *HashTable) LoadFactor() float64 {
	return float64(t.NumEntries) / float64(t.NumBuckets)
}

// Hash handles hashing the key based on the FNV hash function. More can be
// read about it here:
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
func (t *HashTable) Hash(s string) uint64 {
	var hash uint64
	bts := []byte(s)
	hash = FNVOffset
	for _, b := range bts {
		hash *= FNVPrime
		hash ^= uint64(b)
	}
	return hash
}

// Put attempts to insert a new key/value pair or update an existing key with
// the new value.
func (t *HashTable) Put(key string, val interface{}) {
	if t.LoadFactor() > MaxLoadFactor {
		t.resize()
	}
	index := t.Hash(key) % uint64(t.NumBuckets)
	if t.Buckets[index] == nil {
		l := NewList()
		l.Append(key, val)
		t.Buckets[index] = l
		t.NumEntries++
		return
	}
	for n := range t.Buckets[index].Iter() {
		if n.key == key {
			n.val = val
			return
		}
	}
	t.Buckets[index].Append(key, val)
	t.NumEntries++
}

// Get attempts to find the value based on a key.
func (t *HashTable) Get(key string) interface{} {
	index := t.Hash(key) % uint64(t.NumBuckets)
	if t.Buckets[index] == nil {
		return nil
	}
	for n := range t.Buckets[index].Iter() {
		if n.key == key {
			return n.val
		}
	}
	return nil
}

// Del attempts to delete a key/value pair from the hash table.
func (t *HashTable) Del(key string) {
	index := t.Hash(key) % uint64(t.NumBuckets)
	if t.Buckets[index] == nil {
		return
	}
	t.Buckets[index].Remove(key)
	t.NumEntries--
}
