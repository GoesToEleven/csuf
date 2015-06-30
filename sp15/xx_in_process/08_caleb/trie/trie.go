package trie

type (
	Trie struct {
		root   node
		length int
	}
	Iterator struct {
		*inode
	}
	node struct {
		value    interface{}
		children [256]*node
		length   byte
	}
	inode struct {
		parent *inode
		node   *node
		key    []byte
		pos    int
	}
)

func New() *Trie {
	return new(Trie)
}

// Get returns the value associated with the specified key or nil.
func (t *Trie) Get(key []byte) interface{} {
	n := &t.root
	for _, b := range key {
		n = n.children[b]
		if n == nil {
			return nil
		}
	}
	return n.value
}

// Set inserts or replaces the specified key and value.
func (t *Trie) Set(key []byte, value interface{}) bool {
	n := &t.root
	// get to the leaf node (creating any branches along the way)
	for _, b := range key {
		c := n.children[b]
		if c == nil {
			c = new(node)
			n.children[b] = c
			n.length++
		}
		n = c
	}
	added := n.value == nil
	n.value = value
	if added {
		t.length++
	}
	return added
}

// Delete deletes the specified key from the trie.
func (t *Trie) Delete(key []byte) bool {
	path := make([]*node, len(key))
	n := &t.root
	// build a path of nodes
	for i, b := range key {
		path[i] = n
		n = n.children[b]
		// if the node doesn't exist then
		// the key must not be in the trie
		if n == nil {
			return false
		}
	}

	n.value = nil
	// if this is a leaf node, we need to remove
	// it from its parent (and so on up the line
	// for unvalued nodes that become leaves)
	if n.length == 0 {
		for i := len(key) - 1; i >= 0; i-- {
			path[i].children[key[i]] = nil
			path[i].length--
			if path[i].value != nil || path[i].length > 0 {
				break
			}
		}
	}

	t.length--

	return true
}

// Len returns the number of key value pairs in the trie.
func (t *Trie) Len() int {
	return t.length
}

// Iterator returns an iterator that can be used to traverse
// the trie in sorted order. The iterator starts before the
// first key value pair, so Next() should be called before
// Key() and Value()
func (t *Trie) Iterator() *Iterator {
	return &Iterator{&inode{nil, &t.root, []byte{}, -1}}
}

// Next moves the iterator to the next key value pair in the trie
func (i *Iterator) Next() bool {
	// (1) at beginning of node, see if it has a value
	if i.pos == -1 {
		i.pos++
		if i.node.value != nil {
			return true
		}
	}
	// (2) at the end of a node, go up to the parent
	if i.pos > 255 {
		if i.parent == nil {
			return false
		}
		i.inode = i.inode.parent
		return i.Next()
	}
	// (3) walking the node's children
	for ; i.pos < 256; i.pos++ {
		if i.node.children[i.pos] != nil {
			newNode := &inode{
				parent: i.inode,
				node:   i.node.children[i.pos],
				key:    append(i.key, byte(i.pos)),
				pos:    -1,
			}
			i.pos++
			i.inode = newNode
			break
		}
	}

	return i.Next()
}

// Key returns the current key
func (i *Iterator) Key() []byte {
	return i.key
}

// Value returns the current value
func (i *Iterator) Value() interface{} {
	return i.node.value
}
