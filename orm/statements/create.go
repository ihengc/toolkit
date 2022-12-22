package statements

import "fmt"

/*
* @author: Heng ChenChi
* @date: 2022/12/22 0022 17:59
* @version: 1.0
* @description:
**/

type Create struct {
	Table string
}

func (c Create) String() string {
	return fmt.Sprintf("create table `%s`", c.Table)
}
