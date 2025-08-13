package main

import "questbotAI/internal/service"

func main() {
	db := service.ConnectDB()
	service.CreateTable(db)
	service.StartBot()
}
