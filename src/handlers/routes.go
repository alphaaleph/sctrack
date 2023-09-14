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

		Route{"DELETE", "/api/carrier/{id}", DeleteCarrierByID},
		Route{"GET", "/api/carrier/all", GetCarriers},
		Route{"GET", "/api/carrier/{id}", GetCarrierDataByID},
		Route{"POST", "/api/carrier", AddCarrier},

		Route{"DELETE", "/api/journal/{uuid}", DeleteJournalByUUID},
		Route{"GET", "/api/journal/all", GetJournal},
		Route{"GET", "/api/journal/{uuid}", GetJournalEntry},
		Route{"POST", "/api/journal", AddJournal},

		Route{"DELETE", "/api/todos/{uuid}", DeleteTodoByUUID},
		Route{"GET", "/api/todos/all", GetTodos},
		Route{"GET", "/api/todos/{uuid}", GetTodosEntry},
		Route{"POST", "/api/todos", AddTodos},
	}
}
