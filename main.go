package main

import (
	"fmt"
	"net/http"
	"os"

	"go-team/app"
	"go-team/controllers"

	"github.com/gorilla/mux"
)

func appRoute() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/teams", controllers.CreateTeam).Methods("POST")
	router.HandleFunc("/api/teams", controllers.GetTeams).Methods("GET")
	router.HandleFunc("/api/teams/{id}", controllers.FindTeam).Methods("GET")
	router.HandleFunc("/api/teams/{id}", controllers.UpdateTeam).Methods("PUT")
	router.HandleFunc("/api/teams/{id}", controllers.DeleteTeam).Methods("DELETE")

	router.HandleFunc("/api/players", controllers.CreatePlayer).Methods("POST")
	router.HandleFunc("/api/players", controllers.GetPlayers).Methods("GET")
	router.HandleFunc("/api/players/{id}", controllers.FindPlayer).Methods("GET")
	router.HandleFunc("/api/players/{id}", controllers.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/api/players/{id}", controllers.DeletePlayer).Methods("DELETE")

	router.Use(app.JwtAuthentication)

	return router
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("app running at: http://localhost:" + port)

	err := http.ListenAndServe(":"+port, appRoute())
	if err != nil {
		fmt.Print(err)
	}
}
