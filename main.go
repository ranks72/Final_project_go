package main

import (
	"final_project_go/database"
	"final_project_go/handler"
)

func main() {
	database.StartDb()
	handler.StartServer()
}
