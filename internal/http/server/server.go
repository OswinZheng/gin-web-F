package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/OswinZheng/gin-web-F/configs"
	"github.com/OswinZheng/gin-web-F/internal/http/middleware"
	"github.com/OswinZheng/gin-web-F/internal/http/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	// 初始化 gin
	gin.SetMode(configs.Get().Server.RunMode)
	engine := gin.New()
	engine.Use(middleware.RequestIdMiddleware())
	engine.Use(middleware.RateLimit())
	engine.Use(cors.New(cors.Config{
		AllowMethods:    []string{"PUT", "POST", "PATCH", "GET", "DELETE"},
		AllowAllOrigins: true,
	}))
	engine.Use(middleware.CatchError())
	router.InitGraphqlRoute(engine)
	router.InitRouter(engine)

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", configs.Get().Server.Port),
		Handler:        engine,
		ReadTimeout:    time.Duration(configs.Get().Server.ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(configs.Get().Server.WriteTimeOut) * time.Second,
		MaxHeaderBytes: configs.Get().Server.MaxHeaderBytes,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
