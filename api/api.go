package api

import (
	"fmt"
	"jikkosoft/api/controllers"
	"os"
)

var server = controllers.Server{}

func Run() {
	port := fmt.Sprintf(":%v", os.Getenv("HTTP_SERVER_PORT"))
	server.Run(port)
}
