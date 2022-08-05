package BinarySearchTree

func keyhash(s string) int {
	var hash int
	for pos, letter := range s {
		hash = hash + pos*int(letter)
	}
	return hash
}
