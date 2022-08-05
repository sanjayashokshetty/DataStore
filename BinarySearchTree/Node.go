package BinarySearchTree

type Node struct {
	key   int
	value string
	left  *Node
	right *Node
}

func createNode(key, value string, left, right *Node) *Node {
	return &Node{keyhash(key), value, left, right}
}
