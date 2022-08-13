package prototype

// A node have two methods
type INode interface {
	Print(string)
	Clone() INode
}
