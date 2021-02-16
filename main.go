package main

import (
	"jikkosoft/api"
	"jikkosoft/docs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title Rest API
// @version 1.0.0
// @description Rest API for Test project
// @termsOfService http://swagger.io/terms/
// @contact.name Restful API's
// @contact.email kiruu1238@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	err := godotenv.Load()
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_BASE")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	if err != nil {
		log.Fatalf("Error getting Environment Variables %v", err)
	}
	api.Run()
}
