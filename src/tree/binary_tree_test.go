package tree

import (
	"testing"
	"fmt"
)

func TestBinaryTree(t *testing.T) {
	btn := NewBinTreeNode("A")
	bt := NewBinTree(btn)

	btn.SetLChild(NewBinTreeNode("B"))
	btn.SetRChild(NewBinTreeNode("C"))
	btn.GetLChild().SetRChild(NewBinTreeNode("D"))
	btn.GetRChild().SetLChild(NewBinTreeNode("E"))
	btn.GetRChild().SetRChild(NewBinTreeNode("F"))
	btn.GetLChild().GetRChild().SetLChild(NewBinTreeNode("G"))
	btn.GetRChild().GetLChild().SetRChild(NewBinTreeNode("H"))
	btn.GetRChild().GetRChild().SetRChild(NewBinTreeNode("I"))

	fmt.Println(btn.String())

	nodeB := btn.GetLChild()
	nodeF := btn.GetRChild().GetLChild().GetRChild()

	fmt.Println("结点B是叶子结点吗:", nodeB.IsLeaf())
	fmt.Println("结点F是叶子结点吗:", nodeF.IsLeaf())
	fmt.Println("这棵树的结点总数:", bt.GetSize())

	l := bt.PreOrder()
	fmt.Printf("先序遍历:")
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*BinTreeNode)
		fmt.Printf("%v ", (*obj).GetData())
	}
	fmt.Println()

	l = bt.InOrder()
	fmt.Printf("中序遍历:")
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*BinTreeNode)
		fmt.Printf("%v ", (*obj).GetData())
	}
	fmt.Println()

	l = bt.PostOrder()
	fmt.Printf("后序遍历:")
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*BinTreeNode)
		fmt.Printf("%v ", (*obj).GetData())
	}
	fmt.Println()

	result := bt.Find("E")
	fmt.Printf("结点E的父节点数据：%v\t结点E的右孩子节点数据: %v\n", result.GetParent().GetData(), result.GetRChild().GetData())
}
