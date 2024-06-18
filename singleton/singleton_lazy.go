package singleton

import (
	"sync"
)

/*
	单例模式(第一次访问时才创建实例)
*/

var (
	lazySingleton *LazySingleton
	once          = &sync.Once{}
)

type LazySingleton struct {
	data string
}

func GetLazyInstance() *LazySingleton {
	once.Do(
		func() {
			lazySingleton = &LazySingleton{"data"}
		})
	return lazySingleton
}
