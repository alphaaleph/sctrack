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

// getRoutes return a list of all available REST calls
func getRoutes() Routes {
	return Routes{
		Route{"GET", "/", getHomePage},

		// carrier calls
		Route{"DELETE", "/api/carrier/{id}", DeleteCarrierByID},
		Route{"GET", "/api/carrier/all", GetCarriers},
		Route{"GET", "/api/carrier/{id}", GetCarrierDataByID},
		Route{"POST", "/api/carrier", AddCarrier},

		// todos calls
		Route{"DELETE", "/api/todos/carrier/{carrier_id}", DeleteTodosByCarrierID},
		Route{"DELETE", "/api/todos/{uuid}", DeleteTodoByUUID},
		Route{"GET", "/api/todos/all", GetTodos},
		Route{"GET", "/api/todos/carrier/{carrier_id}", GetTodosByCarrierID},
		Route{"GET", "/api/todos/{uuid}", GetTodoByUUID},
		Route{"POST", "/api/todos", AddTodo},

		// journal calls
		Route{"DELETE", "/api/journal/{uuid}", DeleteJournalByUUID},
		Route{"DELETE", "/api/journal/{uuid}/{index}", DeleteJournalByUUIDIndex},
		Route{"GET", "/api/journal/all", GetJournals},
		Route{"GET", "/api/journal/{uuid}", GetJournalByUUID},

		// action calls
		Route{"GET", "/api/action/all", GetActions},
	}
}
