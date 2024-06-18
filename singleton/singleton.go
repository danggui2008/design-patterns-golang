package singleton

/*
	单例模式
*/
type Singleton struct {
	data string
}

var singleton *Singleton

func init() {
	singleton = &Singleton{data: "data"}
}

func GetInstance() *Singleton {
	return singleton
}
