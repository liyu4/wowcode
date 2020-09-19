package  tree 

import (
	"testing"
)

var bst = NewBst()
func TestBstAdd(t  *testing.T) {
	table :=[]int{10,4,12,2,7,11,14}
	for _, tt := range table {
		t.Logf("start add:[%v] element to tree", tt)
		bst.Add(tt)
	}
}

func TestPreOrder(t *testing.T) {
	bst.preOrderTree(bst.root)
}

func TestInOrder(t *testing.T) {
	bst.inOrderTree(bst.root)
}

func TestPostOrder(t *testing.T) {
	bst.postOrderTree(bst.root)
}

