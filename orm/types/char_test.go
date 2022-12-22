package types

import "testing"

/*
* @author: Heng ChenChi
* @date: 2022/12/22 0022 17:54
* @version: 1.0
* @description:
**/

func TestChar_String(t *testing.T) {
	char := Char{Len: 10}
	t.Log(char.String())
}
