package orm

import "testing"

/*
* @author: Heng ChenChi
* @date: 2022/12/21 0021 11:15
* @version: 1.0
* @description:
**/

type testModel struct {
	ID   int64 `tk:"primaryKey"`
	Name string
	Age  int
}

func TestDBConn_Create(t *testing.T) {
	model := testModel{
		Name: "testModel",
		Age:  12,
	}
	db := &DBConn{}
	db.Create(&model)
}
