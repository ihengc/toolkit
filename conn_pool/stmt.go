package conn_pool

import (
	"database/sql/driver"
	"sync"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/12 0012 16:58
* @version: 1.0
* @description:
**/

type stmt struct {
	sync.Locker
	raw driver.Stmt
	// closed stmt是否被关闭
	closed bool
	// closeErr 上次 Close 调用的返回值
	closeErr error
}

// Close 确保 driver.Stmt 只被关闭一次
func (s *stmt) Close() error {
	s.Lock()
	defer s.Unlock()
	if s.closed {
		return s.closeErr
	}
	s.closed = true
	s.closeErr = s.raw.Close()
	return s.closeErr
}
