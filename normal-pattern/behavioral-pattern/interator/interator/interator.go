package interator

type Iterator interface {
	HasNext() bool
	GetNext() *User
}
