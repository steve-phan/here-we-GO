package prototype

import "fmt"

// A folder structure: have childrens and name
type Folder struct {
	Children []INode
	Name     string
}

// Folder implements INode interface
// Folder need to implements Print and Clone
func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, i := range f.Children {
		i.Print(s + s)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{
		Name: f.Name + "_Clone",
	}
	var tempChildren []INode
	// Interace over the Children
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)

	}
	cloneFolder.Children = tempChildren
	/*
		Actually return a cloneFolder.
		But the type of this cloneFolder is a INode
		Simple Straight forward thinking
	*/
	return cloneFolder

}
