package DataStr

import "log"

// LCRS Tree(Left-Child-Right-Sibling Tree)

// Node ...
type Node struct {
	Value  interface{}
	Parent *Node
	Left   *Node //greatest children
	Right  *Node //brother
}

// New ...
func New(v interface{}) *Node {
	return &Node{Value: v}
}

// Insert ...
func Insert(parent *Node, v interface{}) {
	if parent.Left == nil { //往下插入,因此先往左側找尋
		parent.Left = &Node{Value: v, Parent: parent}
	} else { //若左側有東西,表示需要往右邊找(兄弟),直到最右邊後插入
		child := parent.Left
		for child.Right != nil {
			child = child.Right
		}
		child.Right = &Node{Value: v, Parent: child}
	}
}

// FindNode 往下找尋Node 會包含自己的兄弟
// find := partner.FindNode(n, func(node *partner.Node) *partner.Node { //到Tree內去找partner
// 	account := node.Value.(RelationalAccountLevel)
// 	if account.Account == subAccounts[i].ParentAccount {
// 		return node
// 	}
// 	return nil
// })
func FindNode(root *Node, callback func(*Node) *Node) *Node {
	find := callback(root)
	if find != nil {
		return find
	}
	if root.Left != nil {
		find := FindNode(root.Left, callback)
		if find != nil {
			return find
		}
	}
	if root.Right != nil {
		find := FindNode(root.Right, callback)
		if find != nil {
			return find
		}
	}
	return nil
}

// FindParentNode 往上找尋Node
func FindParentNode(root *Node, callback func(*Node) *Node) *Node {
	find := callback(root)
	if find != nil {
		return find
	}
	if root.Parent != nil {
		find := FindParentNode(root.Parent, callback)
		if find != nil {
			return find
		}
	}
	return nil
}

// PrintNode ...
func PrintNode(root *Node) {
	FindNode(root, func(node *Node) *Node {
		log.Printf("%v\n", node.Value)
		return nil
	})
}
