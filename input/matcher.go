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
	Centrifuge(request *http.Request) (*Input, map[string]string)
}

type DefaultHttpMatcher struct {
	antMatcher AntMatcher
	inputs     []*Input
}

func (defaultHttpMatcher *DefaultHttpMatcher) Centrifuge(request *http.Request) (*Input, map[string]string) {
	uri := request.URL.RequestURI()
	for _, input := range defaultHttpMatcher.inputs {
		m, err := defaultHttpMatcher.antMatcher.PattenMatch(uri, input.path)
		if err != nil {
			return input, m
		}
	}
	return nil, nil
}

//模板匹配器
type EngineMatcher interface {
	Match(input *Input) engine.Engine
}
