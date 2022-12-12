package conn_pool

import "time"

/*
* @author: Heng ChenChi
* @date: 2022/12/12 0012 17:08
* @version: 1.0
* @description:
**/

type Stats struct {
	// MaxOpenConnections 最大打开的数据库连接数
	MaxOpenConnections int

	// OpenConnections 成功建立连接的数量，包括正在被使用的和空闲的连接
	OpenConnections int
	// InUse 在使用中的连接的数量
	InUse int
	// Idle 空闲连接的数量
	Idle int

	// WaitCount 等待创建的连接数
	WaitCount int64
	// WaitDuration 创建一个新连接时的等待时长
	WaitDuration time.Duration
	// MaxIdleClosed 达到最大空闲数被关闭的连接数量
	MaxIdleClosed int64
	// MaxIdleTimeClosed 达到最大空闲时长被关闭的连接数量
	MaxIdleTimeClosed int64
	// MaxLifetimeClosed 达到最大存活时长被关闭的连接数量
	MaxLifetimeClosed int64
}
