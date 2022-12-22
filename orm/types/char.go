package types

import "fmt"

/*
* @author: Heng ChenChi
* @date: 2022/12/22 0022 17:53
* @version: 1.0
* @description:
**/

type Char struct {
	Len int
}

func (c Char) String() string {
	return fmt.Sprintf("char(%d)", c.Len)
}
