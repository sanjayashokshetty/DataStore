package main

import (
	"DataStore/BinarySearchTree"
	"fmt"
)

func main() {
	bstTree := BinarySearchTree.Tree{}

	bstTree.InsertBST("abcd", "I am abcd")
	bstTree.InsertBST("efgg", "I am efgh")
	bstTree.InsertBST("hijk", "I am hijk")

	fmt.Println("Searching for abcd, hijk, xyz")

	fmt.Println(bstTree.SearchBST("abcd"))
	fmt.Println(bstTree.SearchBST("hijk"))
	fmt.Println(bstTree.SearchBST("xyz"))
}
