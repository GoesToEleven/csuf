package trie

import (
	"math/rand"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := New()

	cases := []struct{ key, value string }{
		{"car", "white"},
		{"cat", "black"},
	}

	for _, c := range cases {
		if !trie.Set([]byte(c.key), c.value) {
			t.Error("expected trie to not already contain", c.key)
		}
	}

	for _, c := range cases {
		if trie.Set([]byte(c.key), c.value) {
			t.Error("expected trie to already contain", c.key)
		}
	}

	if trie.Len() != 2 {
		t.Error("expected trie to contain 2 items")
	}

	for _, c := range cases {
		v, ok := trie.Get([]byte(c.key)).(string)
		if !ok || v != c.value {
			t.Errorf("expected trie to have %v, got %v", c.value, v)
		}
	}

	for _, c := range cases {
		if !trie.Delete([]byte(c.key)) {
			t.Error("expected trie to delete", c.key)
		}
	}

	for _, c := range cases {
		if trie.Get([]byte(c.key)) != nil {
			t.Error("expected trie to no longer contain", c.key)
		}
	}

	if trie.Len() != 0 {
		t.Error("expected trie to contain 0 items")
	}

	for _, c := range cases {
		trie.Set([]byte(c.key), c.value)
	}

	iterator := trie.Iterator()
	for _, c := range cases {
		if !iterator.Next() {
			t.Error("expected another item")
		}
		if string(iterator.Key()) != c.key {
			t.Errorf("expected %v, got %v", c.key, iterator.Key())
		}
		if str, ok := iterator.Value().(string); !ok || str != c.value {
			t.Errorf("expected %v, got %v", c.value, iterator.Value())
		}
	}
	if iterator.Next() {
		t.Error("expected only", len(cases), "entries")
	}
}

var keys [][]byte

func init() {
	rand.Seed(12345)
	keys = make([][]byte, 10000)
	for i := 0; i < len(keys); i++ {
		sz := 5 + rand.Intn(20)
		keys[i] = make([]byte, sz)
		for j := 0; j < sz; j++ {
			keys[i][j] = byte(rand.Intn(256))
		}
	}
}

type Map map[string]interface{}

func (m Map) Set(key []byte, value interface{}) {
	m[string(key)] = value
}
func (m Map) Get(key []byte) interface{} {
	return m[string(key)]
}

func BenchmarkMapSet(b *testing.B) {
	m := make(Map)
	for i := 0; i < b.N; i++ {
		m.Set(keys[i%len(keys)], i)
	}
}

func BenchmarkTrieSet(b *testing.B) {
	t := New()
	for i := 0; i < b.N; i++ {
		t.Set(keys[i%len(keys)], i)
	}
}

func BenchmarkMapGet(b *testing.B) {
	m := make(Map)
	for i := 0; i < b.N; i++ {
		m.Set(keys[i%len(keys)], i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(keys[i%len(keys)])
	}
}

func BenchmarkTrieGet(b *testing.B) {
	t := New()
	for i := 0; i < b.N; i++ {
		t.Set(keys[i%len(keys)], i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Get(keys[i%len(keys)])
	}
}
