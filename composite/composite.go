package composite

import (
	"fmt"
)

/*
	组合模式
*/

// FileSystemComponent接口【组件（Component）】
type FileSystemComponent interface {
	DisplayInfo()
}

// File【叶子（Leaf）】
type File struct {
	Name string
}

// File实现组件接口
func (f *File) DisplayInfo() {
	fmt.Printf("File %s\n", f.Name)
}

// Directory【复合（Composite）】
type Directory struct {
	name       string
	components []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) AddComponent(component FileSystemComponent) {
	d.components = append(d.components, component)
}

// Directory实现组件接口
func (d *Directory) DisplayInfo() {
	fmt.Printf("Directory %s\n", d.name)
	for _, c := range d.components {
		c.DisplayInfo()
	}
}
