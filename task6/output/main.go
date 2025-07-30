package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {

    data.InitDB()
	data.InitializeUserDB()
	router.RunRouter()
	
}