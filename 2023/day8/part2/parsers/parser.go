package parsers

import (
	avl "github.com/emirpasic/gods/trees/avltree"
)

func Parse(lines []string) (string, *avl.Tree) {
	instructions := lines[0]
	tree := avl.NewWithStringComparator()

	// 0-2    7-9 12-14
	// NTQ = (RMN, HLB)
	for _, line := range lines[2:] {
		key := line[0:3]
		value := []string{line[7:10], line[12:15]}

		tree.Put(key, value)
	}

	return instructions, tree
}
