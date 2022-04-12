package trie

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
}

func (s *Trie) Insert(word string) {
	cur := s.root

	for _, w := range []rune(word) {
		c := string(w)
		if cur.next == nil {
			cur.next = make(map[string]*Node)
		}
		if cur.next[c] == nil {
			cur.next[c] = &Node{next: make(map[string]*Node)}
		}
		cur = cur.next[c]
	}
	cur.isWord = true
}

func (s *Trie) Search(word string) bool {
	cur := s.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}
