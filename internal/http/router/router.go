package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OswinZheng/gin-web-F/internal/api/auth"
	"github.com/OswinZheng/gin-web-F/internal/api/book"
	"github.com/OswinZheng/gin-web-F/internal/http/middleware"
	"github.com/OswinZheng/gin-web-F/internal/repository/rabbitmq"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.ForwardedByClientIP = true
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello")
	})
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	bookRouter := r.Group("/book", middleware.JWT())
	{
		bookRouter.GET("/:id", book.GetBook)
		bookRouter.PUT("/", book.AddBook)
		bookRouter.POST("/:id", book.UpdateBook)
		bookRouter.DELETE("/:id", book.RemoveBook)
	}

	// test rabbitmq
	r.GET("/publish", func(context *gin.Context) {
		go func() {
			for i := 0; i < 30; i++ {
				body := "foobar" + string(i)
				fmt.Println(i)
				if err := rabbitmq.Publish("test-exchange", "direct", "test-key", body, true); err != nil {
					log.Fatalf("%s", err)
				}
				log.Printf("published %dB OK", len(body))
			}
		}()
	})

	r.GET("/consume", func(context *gin.Context) {
		go func() {
			_, err := rabbitmq.NewConsumer("test-exchange", "direct", "test-queue", "test-key", "simple-consumer", func(msg interface{}) error {
				fmt.Println("msg:", string(msg.([]byte)))
				return nil
			})
			if err != nil {
				log.Fatalf("%s", err)
			}
		}()
		context.JSON(http.StatusOK, "consume")
	})
}
