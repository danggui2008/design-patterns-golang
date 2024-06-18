package proxy

import "fmt"

/*
	代理模式
*/

//图像接口【抽象角色】
type Image interface {
	Display()
}

//图片【真实角色】
type RealImage struct {
	fileName string
}

func (r *RealImage) Display() {
	fmt.Printf("displaying  image %s\n", r.fileName)
}

//代理图片【代理角色】
type ProxyImage struct {
	fileName string
	image    RealImage
}

func NewProxyImage(fileName string) *ProxyImage {
	return &ProxyImage{fileName: fileName, image: RealImage{"realFileName"}}
}

func (p *ProxyImage) Display() {
	p.image.Display()
}
