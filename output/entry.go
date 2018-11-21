package output

//输出规范
type Output struct {
	entries []Entry
}

const (
	String = iota
	Float
	Int
)
//一级缓存内存 二级缓存 redis 
type Entry struct {
	name      string
	entryType string
	//返回值是否可缓存
	cacheAble bool
}
