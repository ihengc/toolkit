package clause

/*
* @author: Heng ChenChi
* @date: 2022/12/23 0023 16:28
* @version: 1.0
* @description:
**/

// Select select语句
type Select struct {
	Distinct bool // all | distinct, distinctrow
	Columns  []Column
}

// Build 构建select语句
func (s Select) Build(builder Builder) {
	builder.WriteString("select ")
	if len(s.Columns) > 0 {
		if s.Distinct {
			builder.WriteString("distinct ")
		}
		for index, column := range s.Columns {
			if index > 0 {
				builder.WriteByte(',')
			}
			builder.WriteQuoted(column)
		}
	} else {
		builder.WriteByte('*')
	}
}
