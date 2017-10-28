package main

import (
	"math/rand"
	"time"
)

const MAX_LEVEL = 16

type skip_list struct {
	level  int
	header *node
}

type node struct {
	key     int
	value   interface{}
	forward []*node
}

func newSkipList() *skip_list {
	header := newNodeWithLevel(MAX_LEVEL)
	header.key = 0

	for i := 0; i < MAX_LEVEL; i++ {
		header.forward[i] = nil
	}

	list := &skip_list{
		level:  0,
		header: header,
	}
	return list
}

func (list *skip_list) get(key int) *node {
	x := list.header
	for l := list.level; l >= 0; l-- {
		for x.forward[l] != nil && x.forward[l].key <= key {
			if x.forward[l].key == key {
				return x.forward[l]
			}
			x = x.forward[l]
		}
	}

	return nil
}

func (list *skip_list) insert(key int, value interface{}) *node {
	updates := make([]*node, MAX_LEVEL)

	x := list.header
	for l := list.level; l >= 0; l-- {
		for x.forward[l] != nil && x.forward[l].key < key {
			x = x.forward[l]
		}
		updates[l] = x
	}

	var level int
	level = createRandLevel()
	if level > list.level {
		for i := list.level; i < level; i++ {
			updates[i] = list.header
		}
		list.level = level
	}

	newNode := newNodeWithLevel(level)
	newNode.key = key
	newNode.value = value

	for i := 0; i < level; i++ {
		newNode.forward[i] = updates[i].forward[i]
		updates[i].forward[i] = newNode
	}

	return newNode
}

func createRandLevel() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(MAX_LEVEL-1) + 1
}

func newNodeWithLevel(level int) *node {
	n := &node{}
	n.forward = make([]*node, level)

	return n
}
