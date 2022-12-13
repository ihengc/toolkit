package tcp

import "net"

/*
* @author: Heng ChenChi
* @date: 2022/12/13 0013 12:05
* @version: 1.0
* @description:
**/

// conn 表示一个tcp连接
type conn struct {
	raw net.Conn
}
