package tcp

import (
	"context"
	"sync"
)

/*
* @author: Heng ChenChi
* @date: 2022/12/13 0013 11:56
* @version: 1.0
* @description:
**/

type Server struct {
	Addr    string
	handler Handler

	mu sync.Mutex

	// onShutdown 在调用 Shutdown 时被调用
	onShutdown []func()
}

// Shutdown 关闭tcp服务
func (srv *Server) Shutdown(ctx context.Context) error {
	return nil
}

// RegisterOnShutdown 注册一个函数，函数在 Shutdown 执行时被调用
func (srv *Server) RegisterOnShutdown(f func()) {
	srv.mu.Lock()
	srv.onShutdown = append(srv.onShutdown, f)
	srv.mu.Unlock()
}

func (srv *Server) ListenAndServe() error {
	return nil
}

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{
		Addr:    addr,
		handler: handler,
	}
	return server.ListenAndServe()
}
