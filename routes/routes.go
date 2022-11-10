package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router){
	UserRoutes(r)
	FundRoutes(r)
	AuthRoutes(r)
	// CartRoutes(r)
	// OrderRoutes(r)
	TransactionRoutes(r)

}

