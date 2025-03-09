package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	Nodes [26]*TrieNode
	Value int
	IsEnd bool
}

func InitTrie() *TrieNode {
	return &TrieNode{
		Nodes: [26]*TrieNode{},
		Value: 0,
		IsEnd: false,
	}
}

func NewTrie(value int, isEnd bool) *TrieNode {
	return &TrieNode{
		Nodes: [26]*TrieNode{},
		Value: value,
		IsEnd: isEnd,
	}
}

func (root *TrieNode) Add(key string, val int) {
	// _insert(root, strings.ToLower(key), val, 0)
	var nRoot = InitTrie()
	_put(root, nRoot, strings.ToLower(key), val, 0, 1)
	root.Nodes = nRoot.Nodes
}

func _insert(root *TrieNode, key string, val, index int) {
	if index >= len(key) {
		return
	}

	var charPos = int(key[index]) - 97
	if root.Nodes[charPos] == nil {
		root.Nodes[charPos] = NewTrie(0, false)
	}
	if index == (len(key) - 1) {
		root.Nodes[charPos].Value = val
		root.Nodes[charPos].IsEnd = true
	}

	/*
		below print statement denotes that every time an insert operation is done.. it makes changes
		to original position
	*/
	// fmt.Printf("char_pos: %v, Address: %v\n", charPos, &root.Nodes[charPos].Value)

	_insert(root.Nodes[charPos], key, val, index+1)
}

func _put(root, nRoot *TrieNode, key string, val, index, isRootToConsider int) {
	if index >= len(key) {
		return
	}

	var charPos = int(key[index]) - 97

	nRoot.Nodes[charPos] = NewTrie(0, false)
	if isRootToConsider == 1 {
		// ignore all the nodes which doesnt lie in root to node path
		for i := range 26 {
			if i == charPos {
				continue
			}
			nRoot.Nodes[i] = root.Nodes[i]
		}

		if root.Nodes[charPos] != nil {
			_put(root.Nodes[charPos], nRoot.Nodes[charPos], key, val, index+1, 1)
		} else {
			_put(root, nRoot.Nodes[charPos], key, val, index+1, 0)
		}
	} else {
		_put(root, nRoot.Nodes[charPos], key, val, index+1, 0)
	}

	if index == (len(key) - 1) {
		nRoot.Nodes[charPos].Value = val
		nRoot.Nodes[charPos].IsEnd = true
	}
}

func (root *TrieNode) Get(key string) int {
	return _fetch(root, strings.ToLower(key), 0)
}

func _fetch(root *TrieNode, key string, index int) int {
	var charPos = int(key[index]) - 97
	if index == (len(key) - 1) {
		if root.Nodes[charPos].IsEnd {
			return root.Nodes[charPos].Value
		}

		return -1
	}

	if root.Nodes[charPos] == nil {
		return -1
	}

	return _fetch(root.Nodes[charPos], key, index+1)
}

func main() {
	var root = InitTrie()

	root.Add("apple", 20)
	// fmt.Println(root)

	root.Add("apply", 23)
	// // fmt.Println()

	var val = root.Get("apple")
	fmt.Printf("key: %v, val: %v\n", "apple", val)

	var val2 = root.Get("apply")
	fmt.Printf("key: %v, val: %v\n", "apply", val2)

	var val3 = root.Get("app")
	fmt.Printf("key: %v, val: %v\n", "app", val3)

	root.Add("app", 40)
	var val4 = root.Get("app")
	fmt.Println()
	fmt.Printf("key: %v, val: %v\n", "app", val4)
}
