package main

import (
	"context"
	"log"

	"github.com/KOBATATU/todo/ent"
	"github.com/KOBATATU/todo/ent/migrate"
	controller "github.com/KOBATATU/todo/internal/controller/user"
	db "github.com/KOBATATU/todo/internal/repository/db/user"
	service "github.com/KOBATATU/todo/internal/service/user"
	"github.com/KOBATATU/todo/internal/validation"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	validation.Init()

		// ent.Clientの作成
	client, err := ent.Open("postgres", "host=localhost port=5432 user=user password=postgres dbname=todo sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
    // マイグレーションの実行
    err = client.Schema.Create(
        ctx, 
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true), 
    )
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	userService := service.NewUserService( db.NewUserRepository(client))

	controller.SetupUserRoutes(e, userService)
	e.Logger.Fatal(e.Start(":1323"))
}
