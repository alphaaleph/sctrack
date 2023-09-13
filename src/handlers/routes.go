package handlers

import (
	"net/http"
)

// Route structure defines the REST calls functionality
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes holds a list of REST calls
type Routes []Route

func getRoutes() Routes {
	return Routes{
		Route{"GET", "/", getHomePage},

		//Route{"DELETE", "/api/carrier/{id}", sctrack.Db.DeleteCarrierByID},
		//Route{"GET", "/api/carrier/all", handler.GetCarriers},
		/*Route{"GET", "/api/carrier/{id}", handler.GetCarrierDataByID},
		Route{"POST", "/api/carrier", handler.AddCarrier},

		Route{"DELETE", "/api/journal/{uuid}", handler.DeleteJournalByUUID},
		Route{"GET", "/api/journal/all", handler.GetJournal},
		Route{"GET", "/api/journal/{uuid}", handler.GetJournalEntry},
		Route{"POST", "/api/journal", handler.AddJournal},

		Route{"DELETE", "/api/todos/{uuid}", handler.DeleteTodoByUUID},
		Route{"GET", "/api/todos/all", handler.GetTodos},
		Route{"GET", "/api/todos/{uuid}", handler.GetTodosEntry},
		Route{"POST", "/api/todos", handler.AddTodos},*/
	}
}
