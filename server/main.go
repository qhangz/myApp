package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
    "github.com/myApp/db"
    "github.com/myApp/router"
)

func main() {
	fmt.Println("Hello World")

	db.InitDB()
	
    r := router.InitRouter()
    r.Run(":8080")
}
