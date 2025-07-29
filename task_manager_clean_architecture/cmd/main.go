package main

import "task_manager_clean_architecture/infrastructure/router"

func main() {
    r := router.SetUpRouter()
    r.Run(":8080")
}
