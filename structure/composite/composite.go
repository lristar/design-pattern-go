package composite


import (
	"fmt"
	"strings"
)
const indent = " "

type File struct {
	name string
}

func (f *File)Search(name string){
	if strings.Contains(f.name,name){
		fmt.Println("找到相似的文件--",f.name)
	}
}

type Fold struct {
	children []Composite
	name string
}

func (f *Fold)Search(name string){
	if strings.Contains(f.name,name){
		fmt.Println("找到相似的文件夹--",f.name)
	}
	for _, v := range f.children {
		v.Search(name)
	}
}

func (f *Fold)Append(c Composite){
	f.children = append(f.children, c)
}

type Composite interface {
	Search(name string)
}