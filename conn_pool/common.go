package conn_pool

import "sync"

/*
* @author: Heng ChenChi
* @date: 2022/12/13 0013 14:03
* @version: 1.0
* @description:
**/

// withLock 在持有lock中运行给定的方法
func withLock(lock sync.Locker, fn func()) {
	lock.Lock()
	defer lock.Unlock()
	fn()
}
