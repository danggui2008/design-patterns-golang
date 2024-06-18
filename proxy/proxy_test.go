package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := NewProxyImage("image1.png")
	proxy.Display()
}
