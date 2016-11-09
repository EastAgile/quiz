package trie

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) Add(word string) {
	currentNode := t.root

	for index := 0; index < len(word); index++ {
		char := word[index]
		child := currentNode.Find(char)
		if child != nil {
			currentNode = child
		} else {
			newNode := NewNode()
			newNode.Value = char
			currentNode.addChild(newNode)
			currentNode = newNode
		}
	}

	currentNode.End = true
}

func (t *Trie) Has(word string) bool {
	if len(word) == 0 {
		return false
	}

	currentNode := t.root
	for index := 0; index < len(word); index++ {
		foundNode := currentNode.Find(word[index])
		if foundNode == nil {
			return false
		}
		currentNode = foundNode
	}
	return currentNode.End
}

func (t *Trie) PrefixesOf(word string) (prefixes []string) {
	currentNode := t.root
	strlen := 0

	for index := 0; index < len(word); index++ {
		foundNode := currentNode.Find(word[index])
		if foundNode == nil {
			return prefixes
		}
		currentNode = foundNode
		strlen++

		if currentNode.End {
			prefixes = append(prefixes, word[0:strlen])
		}
	}
	return prefixes
}
