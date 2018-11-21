package engine

type Page struct {
	count int
	data  []map[string][]byte
}

type Listener interface {
	onEvent(event *Event)
}

const (
	BeforeExecute = iota
	AfterExecute
)

type Event struct {
	E      int
	Source Source
}

func IsBeforeExecute(event *Event) bool {
	return event.E == BeforeExecute
}
func IsAfterExecute(event *Event) bool {
	return event.E == AfterExecute
}

type Source interface {
	GetName() string
}

//es查询，sql查询，nosql查询，按需求定制
type Engine interface {
	Execute(param map[string]string) ([]map[string][]byte, error)
	AddListener(listener Listener)
	RemoveListener(listener Listener)
}


//查询功能性能瓶颈一般来源于查询源
type Template interface {
	//模板动态转化，可以在参数校验时同时执行模板转化，以增加解释执行的效率（在接口定下以后，一般不会出现数据校验失败的情况）
	Resolve(block string, param map[string]string) (string, error)
}
