package TREE

type ThreadTree struct {
	Root *ThreadNode
}

func NewEmptyThreadTree() *ThreadTree {
	return &ThreadTree{nil}
}

func NewThreadTree(prelist []interface{}) *ThreadTree {
	tt := NewEmptyThreadTree()
	tt.Root = create(prelist)
	inorderThread(tt.Root,nil)
	return tt
}

func create(preList []interface{}) *ThreadNode {
	p := new(ThreadNode)
	if len(preList) != 0 {
		if preList[0] != nil {
			p.Data = preList[0]
			p.Left = create(preList[1:])
			p.Right = create(preList[1:])
		}
	}
	return p
}

func inorderThread(p *ThreadNode,front *ThreadNode) {
	if p != nil {
		inorderThread(p.Left,front)
		if p.Left == nil {
			p.Ltag = true
			p.Left = front
		}
		if p.Right == nil {
			p.Rtag = true
		}
		if  front != nil && p.Rtag{
			front.Right = p
		}
		front = p
		inorderThread(p.Right,front)
	}
}