package router

import (
	"github.com/gorilla/mux"
	"github.com/wishwaprabodha/go-webscraper/internal/analyzer/adapters/controllers"
)

func StartRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/analyze", controllers.AnalyzeWebPage).Methods("GET")
	return router
}
