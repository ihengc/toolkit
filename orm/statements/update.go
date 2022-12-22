package statements

import "fmt"

/*
* @author: Heng ChenChi
* @date: 2022/12/22 0022 18:01
* @version: 1.0
* @description:
**/

type Update struct {
	Table string
}

func (u Update) String() string {
	return fmt.Sprintf("update `%s` set", u.Table)
}
