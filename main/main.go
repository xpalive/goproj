package main

import (
	"fmt"
	_ "github.com/emicklei/go-restful/v3"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http"
)

func main() {
	fmt.Print("Hello, World!")
}
