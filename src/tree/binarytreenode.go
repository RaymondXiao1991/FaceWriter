package tree

type DataType interface{}

// BinTreeNode
type BinTreeNode struct {
	data   DataType     // 数据域
	parent *BinTreeNode // 父节点
	lChild *BinTreeNode // 左孩子
	rChild *BinTreeNode // 右孩子
	height int          // 以该结点为根的子树的高度
	size   int          // 该结点子孙数(包括结点本身)
}

// NewBinTreeNode 新建节点
func NewBinTreeNode(d DataType) *BinTreeNode {
	return &BinTreeNode{data: d, size: 1}
}

// GetData 获得节点数据
func (btn *BinTreeNode) GetData() DataType {
	if btn == nil {
		return nil
	}
	return btn.data
}

// SetData 设置节点数据
func (btn *BinTreeNode) SetData(d DataType) {
	btn.data = d
}

// HasParent 判断是否有父亲
func (btn *BinTreeNode) HasParent() bool {
	return btn.parent != nil
}

// GetParent 获得父亲节点
func (btn *BinTreeNode) GetParent() *BinTreeNode {
	if btn == nil {
		return nil
	}
	return btn.parent
}

// SetParent 设置父亲节点
func (btn *BinTreeNode) SetParent(p *BinTreeNode) {
	btn.parent = p
}

// CutOffParent 断开与父亲的关系
func (btn *BinTreeNode) CutOffParent() {
	if !btn.HasParent() {
		return
	}
	if btn.IsChild() {
		if btn.IsLChildOf(btn) {
			btn.parent.lChild = nil
		} else if btn.IsRChildOf(btn) {
			btn.parent.rChild = nil
		}
	}
	btn.parent = nil
	btn.SetHeight()
	btn.SetSize()
}

// HasLChild 判断是否有左孩子
func (btn *BinTreeNode) HasLChild() bool {
	return btn.lChild != nil
}

// GetLChild 获得左孩子节点
func (btn *BinTreeNode) GetLChild() *BinTreeNode {
	if btn == nil {
		return nil
	}
	return btn.lChild
}

// SetLChild 设置当前节点的左孩子,返回原左孩子
func (btn *BinTreeNode) SetLChild(lc *BinTreeNode) *BinTreeNode {
	if btn == nil {
		return lc
	}
	btn.lChild = lc
	return lc
}

// HasRChild 判断是否有右孩子
func (btn *BinTreeNode) HasRChild() bool {
	return btn.rChild != nil
}

// GetRChild 获得右孩子节点
func (btn *BinTreeNode) GetRChild() *BinTreeNode {
	if btn == nil {
		return nil
	}
	return btn.rChild
}

// SetRChild 设置当前节点的右孩子,返回原右孩子
func (btn *BinTreeNode) SetRChild(rc *BinTreeNode) *BinTreeNode {
	if btn == nil {
		return rc
	}
	btn.rChild = rc
	return rc
}

// 判断是否为叶子节点
func (btn *BinTreeNode) IsChild() bool {
	return btn.parent != nil
}

// IsLChildOf 判断是否为某结点的左孩子
func (btn *BinTreeNode) IsLChildOf(n *BinTreeNode) bool {
	return n.HasParent() && btn.lChild == n
}

// IsRChildOf 判断是否为某结点的右孩子
func (btn *BinTreeNode) IsRChildOf(n *BinTreeNode) bool {
	return n.HasParent() && btn.rChild == n
}

// GetHeight 取结点的高度,即以该结点为根的树的高度
func (btn *BinTreeNode) GetHeight() int {
	return btn.height
}

// SetHeight 更新当前结点及其祖先的高度
func (btn *BinTreeNode) SetHeight() {
	btn.height = 1
	if btn.HasLChild() {
		btn.height += btn.GetLChild().GetHeight()
	}
	if btn.HasRChild() {
		btn.height += btn.GetRChild().GetHeight()
	}
	if btn.HasParent() {
		btn.parent.SetHeight()
	}
}

// GetSize 取以该结点为根的树的结点数
func (btn *BinTreeNode) GetSize() int {
	return btn.size
}

// SetSize 更新当前结点及其祖先的子孙数
func (btn *BinTreeNode) SetSize() {
	btn.size = 1
	if btn.HasLChild() {
		btn.size += btn.GetLChild().GetSize()
	}
	if btn.HasRChild() {
		btn.size += btn.GetRChild().GetSize()
	}
	if btn.HasParent() {
		btn.parent.SetSize()
	}
}
