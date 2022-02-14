package tree_test

import (
	"algos/tree"
	"testing"
)

func TestBfs(t *testing.T) {

	var root = tree.Node{
		Val:   1,
		Left:  &tree.Node{2, &tree.Node{4, nil, nil}, &tree.Node{5, nil, nil}},
		Right: &tree.Node{3, &tree.Node{6, nil, nil}, &tree.Node{7, nil, nil}},
	}

	tree.Bfs(&root)
}
