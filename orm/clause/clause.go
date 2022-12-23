package clause

import "strings"

/*
* @author: Heng ChenChi
* @date: 2022/12/23 0023 16:30
* @version: 1.0
* @description:
**/

type Builder interface {
	WriteByte(c byte)
	WriteString(s string)
	WriteQuoted(s Column)
}

type Clause struct {
	SQL strings.Builder
}

func (c *Clause) WriteByte(b byte) {
	c.SQL.WriteByte(b)
}

func (c *Clause) WriteString(s string) {
	c.SQL.WriteString(s)
}

func (c *Clause) WriteQuoted(s Column) {
	c.SQL.WriteByte('`')
	c.SQL.WriteString(s.Name)
	c.SQL.WriteByte('`')
}
