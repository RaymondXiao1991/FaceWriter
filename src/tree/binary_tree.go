package tree

import (
	"container/list"
)

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
func (bt *BinTree) GetSize() int {
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

// PreOrder 先序遍历(“根结点-左孩子-右孩子”), 并保存在链表里
func (bt *BinTree) PreOrder() *list.List {
	l := list.New()
	preOrder(bt.root, l)
	return l
}

func preOrder(btn *BinTreeNode, l *list.List) {
	if btn == nil {
		return
	}
	l.PushBack(btn)
	preOrder(btn.GetLChild(), l)
	preOrder(btn.GetRChild(), l)
}

// InOrder 中序遍历(“左孩子-根结点-右孩子”), 并保存在链表里
func (bt *BinTree) InOrder() *list.List {
	l := list.New()
	inOrder(bt.root, l)
	return l
}

func inOrder(btn *BinTreeNode, l *list.List) {
	if btn == nil {
		return
	}
	inOrder(btn.GetLChild(), l)
	l.PushBack(btn)
	inOrder(btn.GetRChild(), l)
}

// PostOrder 后序遍历(“左孩子-右孩子-根结点”), 并保存在链表里
func (bt *BinTree) PostOrder() *list.List {
	l := list.New()
	postOrder(bt.root, l)
	return l
}

func postOrder(btn *BinTreeNode, l *list.List) {
	if btn == nil {
		return
	}
	postOrder(btn.GetLChild(), l)
	postOrder(btn.GetRChild(), l)
	l.PushBack(btn)
}

// Walk 步进 tree bt 将所有的值从 tree 发送到 channel ch。
func Walk(bt *BinTree, ch chan DataType) {
	rangeTree(bt.root, ch)
	close(ch)
}

// rangeTree 遍历二叉树，把当前节点值传入通道
func rangeTree(btn *BinTreeNode, ch chan DataType) {
	if btn != nil {
		rangeTree(btn.GetLChild(), ch)
		ch <- btn.GetData()
		rangeTree(btn.GetRChild(), ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *BinTree) bool {
	//建立两个通道
	ch1 := make(chan DataType)
	ch2 := make(chan DataType)
	//遍历两个二叉树，把值传入各自的通道
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	//遍历通道进行比较，有不同的值就返回false
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}
