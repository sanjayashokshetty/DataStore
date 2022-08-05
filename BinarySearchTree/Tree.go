package BinarySearchTree

type Tree struct {
	root *Node
}

func (t *Tree) InsertBST(key, value string) {
	if t.root == nil {
		t.root = createNode(key, value, nil, nil)
	} else {
		insertNode(t.root, createNode(key, value, nil, nil))
	}
}

func insertNode(node *Node, newNode *Node) {
	if node.key > newNode.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (t *Tree) SearchBST(key string) *Node {
	if t.root == nil {
		return nil
	}
	hash := keyhash(key)
	return searchNode(hash, t.root)
}

func searchNode(key int, currNode *Node) *Node {
	if currNode == nil {
		return nil
	}
	if currNode.key == key {
		return currNode
	}
	if currNode.key > key {
		return searchNode(key, currNode.left)
	} else {
		return searchNode(key, currNode.right)
	}
}
