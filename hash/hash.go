package hash

import "sort"

const (
	NEW       = "NEW"
	UPDATED   = "UPDATED"
	NOT_FOUND = "NOT_FOUND"
	ABSENT    = "ABSENT"
	ERASED    = "ERASED"
)

type Outdump struct {
	Key   string
	Value int64
}

type entry struct {
	key   string
	value int64
	next  *entry
}

type HashTable struct {
	buckets []*entry
	size    int
}

func (h *HashTable) bucketIndex(key string) int {
	return int(hash(key) % uint64(len(h.buckets)))
}

func hash(key string) uint64 {
	h := uint32(21056) // Только младшие 32 бита
	for i := 0; i < len(key); i++ {
		h = h*33 + uint32(key[i])
	}
	return uint64(h)
}

func New() *HashTable {
	return &HashTable{
		buckets: make([]*entry, 8),
		size:    0,
	}
}

func (h *HashTable) rehash() {
	old := h.buckets
	h.buckets = make([]*entry, len(old)*2)
	for _, head := range old {
		for i := head; i != nil; {
			next := i.next
			j := h.bucketIndex(i.key)
			i.next = h.buckets[j]
			h.buckets[j] = i
			i = next
		}
	}
}

func (h *HashTable) Insert(key string, value int64) string {
	index := h.bucketIndex(key)
	for i := h.buckets[index]; i != nil; i = i.next {
		if i.key == key {
			i.value = value
			return UPDATED
		}
	}
	h.buckets[index] = &entry{key: key, value: value, next: h.buckets[index]}
	h.size++
	if float64(h.size)/float64(len(h.buckets)) > 1 {
		h.rehash()
	}
	return NEW
}

func (h *HashTable) Erase(key string) string {
	index := h.bucketIndex(key)
	i := h.buckets[index]
	prev := i
	if i == nil {
		return ABSENT
	}
	if i.key == key {
		h.buckets[index] = i.next
		h.size--
		return ERASED
	}
	for i != nil {
		if i.key == key {
			prev.next = i.next
			h.size--
			return ERASED
		}
		prev = i
		i = i.next
	}

	return ABSENT
}

func (h *HashTable) Find(key string) (int64, bool) {
	index := h.bucketIndex(key)
	for i := h.buckets[index]; i != nil; i = i.next {
		if i.key == key {
			return i.value, true
		}
	}
	return 0, false
}

func (h *HashTable) Size() int {
	return h.size
}

func (h *HashTable) Dump() []Outdump {
	var dump []Outdump
	for i := 0; i < len(h.buckets); i++ {
		temp := h.buckets[i]
		for temp != nil {
			data := &Outdump{Key: temp.key, Value: temp.value}
			dump = append(dump, *data)
			temp = temp.next
		}
	}
	sort.Slice(dump, func(i, j int) bool {
		return dump[i].Key < dump[j].Key
	})
	return dump
}
