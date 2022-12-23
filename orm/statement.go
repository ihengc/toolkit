package orm

/*
* @author: Heng ChenChi
* @date: 2022/12/23 0023 15:37
* @version: 1.0
* @description:
**/

// Statement 表示单个表的各种SQL语句
type Statement struct {
	engine *Engine
	// SelectColumns 查询表字段
	SelectColumns []string
	// Distinct 查询去重
	Distinct bool
	// Clauses 语句
	Clauses map[string]string
}

// BuildCondition 构建条件语句
func (s *Statement) BuildCondition() {

}
