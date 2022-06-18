package prototype

import "fmt"

type INode interface {
	print(string)
	clone()INode
}

type File struct {
	name string
}

func (f *File)print(indentation string){
	fmt.Println(indentation+f.name)
}

func (f *File)clone()INode{
	return &File{name:f.name+"_clone"}
}
type Folder struct {
	children []INode
	name string
}

func (f *Folder)print(indentation string){
	fmt.Println(indentation+f.name)
	for _, v := range f.children {
		v.print(indentation+indentation)
	}
}

func (f *Folder)clone()INode{
	cloneNode :=&Folder{name:f.name+"_clone"}
	temp:=make([]INode,0)
	for _, v := range f.children {
		t :=v.clone()
		temp = append(temp,t )
	}
	cloneNode.children = temp
	return cloneNode
}

func (f *Folder)Append(node INode){
	f.children = append(f.children, node)
}