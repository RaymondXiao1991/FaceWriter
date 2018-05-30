package tree

import "queue"

// 若设二叉树的深度为h，除第h层外，其它各层(1～h-1)的结点数都达到最大个数，第h层所有的结点都连续集中在最左边，这就是完全二叉树（Complete Binary Tree）
/*首先若一个节点只有右孩子，肯定不是完全二叉树；
其次若只有左孩子或没有孩子，那么对于一个高度为h的完全二叉树，当前节点的高度肯定是h-1，也就是高度h的所有节点都没有孩子，否则不是完全二叉树，
因此设置flag标记当前节点是不是到了h-1高度*/

// IsCompleteTree 判断是否是完全二叉树
func (bt *BinTree) IsCompleteTree() bool {
	flag := false
	q := queue.NewQueue()
	q.Push(bt.root)

	for !q.Empty() {
		p := q.Pop()
		if flag {
			if p.HasLChild() || p.HasRChild() {
				return false
			} else {
				if p.HasLChild() && p.HasRChild() {
					q.Push(p.GetLChild)
					q.Push(p.GetRChild)
				} else if p.HasRChild() {
					q.Push(p.GetLChild)
					flag = true
				} else {
					flag = true
				}
			}
		}
	}
	return true
}
