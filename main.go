package main

import (
	"github.com/OswinZheng/gin-web-F/configs"
	"github.com/OswinZheng/gin-web-F/internal/http/server"
	"github.com/OswinZheng/gin-web-F/internal/repository/postgres"
	"github.com/OswinZheng/gin-web-F/internal/repository/redis"
	"github.com/OswinZheng/gin-web-F/migrate"
)

//go:generate gqlgen

func init() {
	// 加载配置
	configs.InitConfig()
	// 初始化数据库
	postgres.New()
	// 数据库迁移 migrate
	migrate.Run()
	// 连接redis
	redis.Client()
	// 定时任务
	//go schedule.Start()
}

func main() {
	// 初始化 http 服务
	server.Run()
}
