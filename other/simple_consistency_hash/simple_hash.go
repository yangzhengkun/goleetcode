package simple_consistency_hash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashFunc func(key []byte) uint32

type Map struct {
	hash     HashFunc
	replicas int
	keycode  []int
	hashMap  map[int]string
}

func NewMap(replicas int, fn HashFunc) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keycode = append(m.keycode, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keycode)
}

func (m *Map) Get(key string) string {
	if len(m.keycode) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keycode), func(i int) bool { return m.keycode[i] >= hash })
	if idx == len(m.keycode) {
		idx = 0
	}
	return m.hashMap[m.keycode[idx]]
}
