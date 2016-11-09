package main

import (
	"github.com/EastAgile/quiz/trie"
	"testing"
)

var tests = []struct {
	list []string
	want string
}{
	{list: []string{"a", "b", "ab", "abc", "abcd"}, want: "ab"},
	{list: []string{"ab", "cd"}, want: ""},
}

func TestLongestCompound(test *testing.T) {
	for index := range tests {
		item := tests[index]
		t := trie.NewTrie()
		for i := range item.list {
			t.Add(item.list[i])
		}
		want := item.want
		got := longestCompound(item.list, t)
		if got != want {
			test.Errorf("Expect %s, Got %s", want, got)
		}
	}
}
