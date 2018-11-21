package input

import (
	"net/http"
	"suction-pump/engine"
)

type AntMatcher interface {
	//如果路径不匹配则返回error，否则返回匹配成功的参数
	PattenMatch(path string, patten string) (map[string]string, error)
}

type HttpMatcher interface {
	//通过请求，匹配到一个输入实体,跟参数信息
	Centrifuge(request *http.Request) (Input, map[string]string)
}

//模板匹配器
type EngineMatcher interface {
	Match(input *Input) engine.Engine
}
