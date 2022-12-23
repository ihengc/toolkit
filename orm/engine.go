package orm

import (
	"database/sql"
	"errors"
	"reflect"
	"toolkit/orm/dialect"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/23 0023 15:38
* @version: 1.0
* @description:
**/

type Engine struct {
	dbPool *sql.DB
	stmts  map[string]*Statement
	config *Config
}

type Config struct {
}

type OptionInterface interface {
	apply(option OptionInterface)
}

func Open(dialect dialect.DialectInterface, options ...OptionInterface) (*Engine, error) {
	dbPool, err := sql.Open(dialect.GetDriverName(), dialect.GetDataSourceName())
	if err != nil {
		return nil, err
	}
	engine := &Engine{}
	engine.dbPool = dbPool
	return engine, nil
}

func (e *Engine) All(model interface{}) ([]interface{}, error) {
	var modelName string
	modelType := reflect.TypeOf(model)
	switch modelType.Kind() {
	case reflect.Pointer:
	case reflect.Map:
	case reflect.Struct:
	}
	if modelName == "" {
		return nil, errors.New("unknown model kind")
	} else {
		stmt, ok := e.stmts[modelName]
		if !ok {
			stmt = NewStatement(model)
			e.stmts[modelName] = stmt
		}
		return stmt.All()
	}
}
