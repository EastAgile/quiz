package main

import (
	"./queue"
	"./trie"
	"bufio"
	"fmt"
	"os"
)

func LongestCompound(wordList []string, t *trie.Trie) string {
	var prefixes []string

	q := queue.NewQueue()

	for index := range wordList {
		word := wordList[index]
		prefixes = t.PrefixesOf(word)

		for pindex := range prefixes {
			pword := prefixes[pindex]
			suffix := word[len(pword):len(word)]
			q.Push(queue.NewNode(word, suffix))
		}
	}

	longestString := ""
	longestStringLength := 0

	for {
		qNode := q.Pop()
		if qNode == nil {
			break
		} else {
			currentLength := len(qNode.Word)

			if longestStringLength >= currentLength {
				// Do not process if currentLength < current longest length
				continue
			}
			if t.Has(qNode.Suffix) && (longestStringLength < currentLength) {
				longestString = qNode.Word
				longestStringLength = len(qNode.Word)

			} else {
				prefixes = t.PrefixesOf(qNode.Suffix)
				for pindex := range prefixes {
					pword := prefixes[pindex]
					suffix := qNode.Suffix[len(pword):len(qNode.Suffix)]
					q.Push(queue.NewNode(qNode.Word, suffix))
				}
			}
		}
	}
	return longestString
}

func main() {
	wordList := make([]string, 0)
	t := trie.NewTrie()

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Missing file argument!")
		return
	}
	file, err := os.Open(args[0])

	if err != nil {
		fmt.Println("Cannot open file!")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		wordList = append(wordList, word)
		t.Add(word)
	}

	fmt.Printf("Longest compound word is: %s\n", LongestCompound(wordList, t))
}
