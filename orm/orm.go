package orm

import (
	"database/sql"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/21 0021 10:31
* @version: 1.0
* @description:
**/

type DBConn struct {
	connPool *sql.DB
}
