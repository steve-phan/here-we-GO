package main

import (
	"fmt"

	"bookable24.de/main/normal-pattern/creational-pattern/prototype/prototype"
)

func main() {
	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}

	folder1 := &prototype.Folder{
		Children: []prototype.INode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.INode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrint the origin Folder2")
	folder2.Print("   ")

	cloneFolder := folder2.Clone()
	fmt.Println(("\nPrint the clone of Folder2"))
	cloneFolder.Print("   ")

}
