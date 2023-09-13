package handlers

import (
	"fmt"
	"github.com/alphaaleph/sctrack"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"strings"
)

// getHomePage presents a welcome screen for the root URL of the server.
func getHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WELCOME TO THE DASHBOARD INJEST SERVER at %s", getHostname(w))
}

// getHostname will return the server name.
func getHostname(w http.ResponseWriter) string {
	sctrack.Log.Info("In get hostname")
	//get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		sctrack.Log.Error("Failed to get host name", slog.String("Error", err.Error()))
	}
	return strings.ToUpper(hostname)
}
