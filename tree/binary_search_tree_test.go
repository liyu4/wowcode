package  tree 

import (
	"testing"
)

func TestBstAdd(t  *testing.T) {
	table :=[]int{10,4,12}
	bst := NewBst()

	for _, tt := range table {
		t.Logf("start add:[%v] element to tree", tt)
		bst.Add(tt)
	}
}