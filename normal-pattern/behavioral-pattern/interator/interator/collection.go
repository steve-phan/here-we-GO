package interator

type Collection interface {
	CreateInterator() Iterator
}
