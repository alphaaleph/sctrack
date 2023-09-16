package handlers

import (
	"fmt"
	"github.com/alphaaleph/sctrack"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
)

// StartYourEngines will set the web service
func StartYourEngines() {

	// start a new server router
	router := newRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	port := os.Getenv("SERVER_PORT")
	http.Handle("/", router)
	sctrack.Log.Info(fmt.Sprintf("Server started at port :%s", port))

	// start web service
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		sctrack.Log.Error("Failed to start the server.", slog.String("Error", err.Error()))
		panic(err)
	}
}

// newRouter will manage the REST calls
func newRouter() (router *mux.Router) {

	sctrack.Log.Debug("Registering routes")
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range getRoutes() {

		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)
	}
	return
}
