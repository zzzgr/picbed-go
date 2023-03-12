package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"picbed-go/logger"
	"picbed-go/server/router"
	"picbed-go/setting"
	"strconv"
	"syscall"
)

func main() {
	// 配置
	err := setting.SetUp()
	if err != nil {
		fmt.Println("加载配置出错: " + err.Error())
		return
	}

	// 日志
	err = logger.SetUp()
	if err != nil {
		zap.L().Error("日志配置出错: " + err.Error())
		return
	}

	// 路由
	r := router.Setup()

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.App.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	zap.L().Info("服务已启动, 端口: " + strconv.Itoa(setting.App.Port))

	// 等待终端信号来优雅关闭服务器
	quit := make(chan os.Signal, 1) // 创建一个接受信号的通道

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞此处，当接受到上述两种信号时，才继续往下执行
	zap.L().Info("服务已停止")
}
