package main

import (
	"fmt"
)

type node struct {
	value string
	next  *node
	pre   *node
}

func NewNode() *node {
	return &node{
		value: "",
		next:  nil,
		pre:   nil,
	}
}

type hmap [7]node

func NewHmap() *hmap {
	return &hmap{}
}
func main() {
	h := NewHmap()
	fmt.Print("%v\n", h)
}
