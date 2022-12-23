package orm

import "reflect"

/*
* @author: Heng ChenChi
* @date: 2022/12/23 0023 15:37
* @version: 1.0
* @description:
**/

type Statement struct {
	name      string
	model     interface{}
	columns   []string
	engine    *Engine
	selectSQL string
}

func (s *Statement) init(model interface{}) {

}

func NewStatement(model interface{}) *Statement {
	stmt := &Statement{}
	stmt.init(model)
	return stmt
}

func (s *Statement) All() ([]interface{}, error) {
	rows, err := s.engine.dbPool.Query(s.selectSQL)
	if err != nil {
		return nil, err
	}
	results := make([]interface{}, 8)
	for rows.Next() {
		p := reflect.New(reflect.TypeOf(s.model)).Pointer()
		err := rows.Scan(p)
		if err != nil {
			return nil, err
		}
		results = append(results, p)
	}
	return results, nil
}
