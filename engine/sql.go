package engine

import "github.com/jmoiron/sqlx"

type SqlOptimization interface {
	optimize(sql string) string
}

//sql分页包装器
type PageSqlWrapper interface {
	Wrapper(sql string) string
	//可以做查询优化器
	CountWrapper(sql string) string
}

type NamedSqlEngine struct {
	template     Template
	db           *sqlx.DB
	text         string
	optimization SqlOptimization
}

func (namedSqlEngine *NamedSqlEngine) ExecutePage(param map[string]string, pageSize, pageNo int) (*Page, error) {
	return nil, nil
}
func (namedSqlEngine *NamedSqlEngine) Execute(param map[string]string) ([]map[string][]byte, error) {
	sql, err := namedSqlEngine.template.Resolve(namedSqlEngine.text, param)
	if err != nil {
		return nil, err
	}
	result, err := namedSqlEngine.db.NamedQuery(sql, param)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	return nil, nil
}

type SqlEngine NamedSqlEngine

func (sqlEngine *SqlEngine) ExecutePage(param map[string]string, pageSize, pageNo int) (*Page, error) {
	return nil, nil
}
func (sqlEngine *SqlEngine) Execute(param map[string]string) ([]map[string][]byte, error) {
	sql, err := sqlEngine.template.Resolve(sqlEngine.text, param)
	if err != nil {
		return nil, err
	}
	result, err := sqlEngine.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	return nil, nil
}
