package controllers

import (
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (server *Server) initializeRoutes() {
	server.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(os.Getenv("SWAGGER_BASE")+"/docs/swagger.yaml"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	server.Router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	server.Router.HandleFunc("/sort", server.SortArray).Methods("POST")
	server.Router.HandleFunc("/user", server.CreateUser).Methods("POST")
	server.Router.HandleFunc("/user", server.GetUsers).Methods("GET")
}
