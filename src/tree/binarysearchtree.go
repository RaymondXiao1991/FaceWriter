package tree

// 二叉查找树的特点：对于每一个节点X,左子树中所有项的值小于X中的值，而它的右子树种所有项的值大于X中的值。

type BinarySearchTree struct {
	root   *binSearchTreeNode // 根节点
	height int                // 树高
	size   int                // 节点数
}

// 节点
type binSearchTreeNode struct {
	key int // 键
	BinTreeNode
}

// NewBinSearchTreeNode 新建节点
func NewBinSearchTreeNode(k int, d DataType) *binSearchTreeNode {
	bstn := &binSearchTreeNode{key: k}
	bstn.SetData(d)
	bstn.SetSize()

	return bstn
}

// 初始化并添加根节点的值
func NewBinarySearchTree(root DataType) *BinarySearchTree {
	return &BinarySearchTree{root: root}
}

/*type Tree interface {
	Put(k, v int)  //新增或修改
	Get(k int) int //查询
	Delete(k int)  //删除
	Size() int     //树的大小
	Min() int      //最小键
	DeleteMin()    //删除最小键
}*/

// 添加元素
func (bst *BinarySearchTree) Put(k int, v DataType) {
	bst.root, _ = put(bst.root, k, v)
}

//在以nd为根节点的树下增加或修改一个节点
//如果创建了新节点，第二个参数返回true，
//如果只是修改，第二个参数返回false
func put(btn *binSearchTreeNode, k, v int) (*binSearchTreeNode, bool) {
	if btn == nil {
		return NewBinSearchTreeNode(k, v), true
	}
	hasNew := false //检测是否创建了新节点
	if k < btn.k {
		btn.left, hasNew = put(nd.left, k, v)
	} else if k > nd.k {
		btn.right, hasNew = put(nd.right, k, v)
	} else {
		btn.v = v //仅修改，不会增加增加节点，就不更新节点大小
	}
	if hasNew { //如果创建了新节点就更新树的大小
		updateSize(btn)
	}
	return btn, hasNew
}

// 是否包含某一个元素

// 移除元素

// 查找最大值

// 查找最大的节点

// 查找最小值

// 查找最小的节点

// 获取树种所有的元素值（并按从小到大排序）
