sssspackage tree

type treeListNode struct {
	value int 
	parent  *treeListNode
	right *treeListNode
	leaft *treeListNode
}


type Bst struct {
	root *treeListNode
}

func NewBst() *Bst{
	return &Bst{}
}

//--------------------------
//			  10
//         4     12
//       2   7  11 14
//
//--------------------------
// 添加元素
/*
性质： 二叉搜索树的基本性质是，根节点大于左边节点，小于右边节点
也就是10大于4并且小于12
节点4和12也遵循这个规律
假设现在插入3则要经过这些步骤 
1: 3和根节点的value10比较，3小于根节点的值，则说明这个元素应该放在根节点的左边，继续
2：左边的元素是4， 3和4比较，3比4小，说明3应该放在4的左边，继续
3：找到4的左边是2，3和2比较，3比2大，说明3应该放在2的右边，继续
4：找到2的右边是空，则将3放置在2的右边，结束
*/


func (bst *Bst) Add(value int) {
	if bst.root == nil { // 找不到父节点，说明此节点为根节点
		bst.root = new(treeListNode)
		bst.root.value =value
		return 
	}
	root := bst.root
	find := new(treeListNode)
	for root != nil {
	if root.right == nil || root.leaft == nil {
			find = root
			root  = nil
			continue
	}
	
	if value > root.value {
				find = root.right
				root = find.right
			
		}else{
				find = root.leaft
				root  = find.leaft			
		}
	}

	if value > find.value {
		find.right = &treeListNode{value, nil,nil, nil}	
	}else{
		find.leaft = &treeListNode{value, nil,nil, nil}	
	}
}