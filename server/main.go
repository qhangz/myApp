package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
    "github.com/myapp/db"
    "github.com/myapp/router"
)

func main() {
	fmt.Println("Hello World")

	db.InitDB()
	
    r := router.InitRouter()
    r.Run(":8080")
}
