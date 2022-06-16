package migrate

import (
	"fmt"

	"github.com/OswinZheng/gin-web-F/internal/model/auth"
	"github.com/OswinZheng/gin-web-F/internal/model/book"
	"github.com/OswinZheng/gin-web-F/internal/repository/postgres"
)

func Run() {
	db := postgres.Db
	if !db.Migrator().HasTable("user") {
		fmt.Println("migrate user")
		_ = db.AutoMigrate(&auth.Auth{})
	}

	if !db.Migrator().HasTable("book") {
		fmt.Println("migrate book")
		_ = db.AutoMigrate(&book.Book{})
	}
	fmt.Println("migrate success")
}
