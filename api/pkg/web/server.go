package web

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"some/api/pkg/config"
	"some/api/routers"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Config config.Server
	Server *http.Server
}

func NewServer() *Server {
	return &Server{
		Config: config.Conf.Server,
	}
}

func (w *Server) Start() {
	addr := fmt.Sprintf(":%d", w.Config.HttpPort)
	logrus.Info("启动web服务，监听端口：", w.Config.HttpPort)

	w.Server = &http.Server{
		Addr:              addr,
		Handler:           routers.Router,
		ReadHeaderTimeout: w.Config.ReadTimeout,
		WriteTimeout:      w.Config.WriteTimeout,
	}

	// pprof
	go func() {
		addr := fmt.Sprintf(":%v", 8807)
		logrus.Info(http.ListenAndServe(addr, nil))
	}()

	go func() {
		// service connections
		if err := w.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		} else {
			logrus.Info("Web server close...")
		}
	}()

	// 阻塞进程，直到接收到关闭信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

// 关闭http服务
// 参考资料： https://www.dazhuanlan.com/2019/11/28/5ddf97117da5a/
func (w *Server) Destroy() {
	ctx, cancel := context.WithTimeout(context.Background(), w.Config.CloseTimeout)
	defer cancel()

	fmt.Println("close timeout:", w.Config.CloseTimeout)
	if err := w.Server.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		logrus.Infof("timeout of %d seconds.", w.Config.CloseTimeout) // TODO 这里好像有问题
	}

	logrus.Info("Web server close.")
}
