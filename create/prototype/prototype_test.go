package prototype

import "testing"

func TestPrototype(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}
	file4 := &File{name: "File4"}
	folder1 := &Folder{
		children: []INode{},
		name:      "Folder1",
	}
	folder1.Append(file1)
	folder1.Append(file2)
	folder1.Append(file3)
	folder1.Append(&Folder{
		children: []INode{file4},
		name:      "Folder2",
	})
	folder1.print(" ")
	folder_clone:=folder1.clone()
	folder_clone.print(" ")
}