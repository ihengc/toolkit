package dialect

/*
* @author: Heng ChenChi
* @date: 2022/12/23 0023 15:39
* @version: 1.0
* @description:
**/

type DialectInterface interface {
	GetDriverName() string
	GetDataSourceName() string
}
