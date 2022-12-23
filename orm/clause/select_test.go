package clause

import (
	"strings"
	"testing"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/23 0023 16:37
* @version: 1.0
* @description:
**/

type selectClause struct {
	sql strings.Builder
}

func (s *selectClause) WriteByte(c byte) {
	s.sql.WriteByte(c)
}

func (s *selectClause) WriteString(c string) {
	s.sql.WriteString(c)
}

func (s *selectClause) WriteQuoted(c Column) {
	s.sql.WriteByte('`')
	s.sql.WriteString(c.Name)
	s.sql.WriteByte('`')
}

func TestSelect_Build(t *testing.T) {
	sel := Select{}
	sel.Columns = []Column{{Name: "age"}, {Name: "account"}, {Name: "email"}}
	sel.Distinct = true
	builder := &selectClause{}
	sel.Build(builder)
	t.Log(builder.sql.String())
}
