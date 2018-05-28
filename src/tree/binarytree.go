package tree

import "container/list"

// BinTree 二叉树
type BinTree struct {
	root   *BinTreeNode // 根节点
	height int          // 树高
	size   int          // 节点数
}

// NewBinTree 新建二叉树
func NewBinTree(root *BinTreeNode) *BinTree {
	return &BinTree{root: root}
}

// GetSize 获得二叉树总节点数
func (btn *BinTree) GetSize() int {
	return bt.size
}

// IsEmpty判断二叉树是否为空
func (bt *BinTree) IsEmpty() bool {
	return bt == nil
}

// GetData 获得二叉树根节点
func (bt *BinTree) GetData() DataType {
	return bt.root.GetData()
}

// GetHeight 获得二叉树高度
func (bt *BinTree) GetHeight() int {
	return bt.height
}

// Find 查询第一个与数据e相等的节点
func (bt *BinTree) Find(e DataType) *BinTreeNode {
	if bt == nil {
		return nil
	}
	return isEqual(bt.root, e)
}

// 判断两棵二叉树是相等
func isEqual(bt *BinTreeNode, e DataType) *BinTreeNode {
	if bt.GetData() == e {
		return bt
	}
	if bt.HasLChild() {
		if p := isEqual(bt.GetLChild(), e); p != nil {
			return p
		}
	}
	if bt.HasRChild() {
		if p := isEqual(bt.GetRChild(), e); p != nil {
			return p
		}
	}
	return nil
}

// 先序遍历, 并保存在链表里

func preOrder(btn *BinTreeNode, l *list.List) {

}

// 中序遍历, 并保存在链表里

// 后序遍历, 并保存在链表里
