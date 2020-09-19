package tree

import(
	"fmt"
)

type treeListNode struct {
	value  int
	parent *treeListNode
	right  *treeListNode
	left  *treeListNode
}

type Bst struct {
	root *treeListNode
}

func NewBst() *Bst {
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
		bst.root.value = value
		return
	}
	root := bst.root
	find := new(treeListNode)
	for root != nil {
		find = root
		if root.right == nil || root.left == nil {
			root = nil
			continue
		}
		if value > root.value {
			root = root.right
		} else {
			root = root.left
		}
	}

	if value > find.value {
		find.right = &treeListNode{value, find, nil, nil}
	} else {
		find.left = &treeListNode{value, find, nil, nil}
	}
}


//--------------------------
//			  10
//         4     12
//       2   7  11 14
//
//--------------------------
// 前序遍历
/*
性质，前序遍历是说从先遍历根节点，再遍历左边的节点到再遍历右边的节点
1： 先拿到根节点10
2： 再拿到左边的节点4
3： 在拿到左边的节点2
4： 2是叶子节点，下面左右都是空，则4的左边遍历完毕
5： 开始遍历4的右边，得到7
6： 遍历7的左边，为空
7： 遍历7的右边，为空，则10的左边遍历完毕
则最终的结果是： 10 4 2 7 12 11 14

测试代码：
root value:10
value:4
value:2
value:7
value:12
value:11
value:14

*/




func (bst *Bst) preOrderTree(node *treeListNode) {
	if node == nil {
		return
	}
	fmt.Printf("pre order value:%v\n",node.value)
	bst.preOrderTree(node.left)
	bst.preOrderTree(node.right)
}


func (bst *Bst) inOrderTree(node *treeListNode) {
	if node  == nil {
		return
	}

	bst.inOrderTree(node.left)
	fmt.Printf("in order tree, input value:[%d]\n",node.value)
	bst.inOrderTree(node.right)
}


func (bst *Bst) postOrderTree(node *treeListNode) {
	if node  == nil {
		return
	}
	bst.postOrderTree(node.left)
	bst.postOrderTree(node.right)
	fmt.Printf("after order tree, input value:[%d]\n",node.value)
}

