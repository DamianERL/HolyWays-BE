package main

import (
	"fmt"
	"os"
	"net/http"
	"holyways/database"
	"holyways/pkg/mysql"
	"holyways/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	//initialDB
	mysql.DataBaseinit()
	//run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	var port = os.Getenv("PORT")
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe("localhost:" + port,r)

}
