package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/themadman159/starter-api/routers"
	"github.com/themadman159/starter-api/utils"
)

func main() {
	db := utils.NewDatabase()
	_, sqlDB, err := db.ConnectMysql()
	defer sqlDB.Close()
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	routers.Route(app)

	app.Listen("server start at port : 8080")
}
