package routes

import (
	"net/http"
)

var usersRoutes = []Route{
	Route{
		Uri:     "/users",
		Methods: http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Uri:     "/user/{id}",
		Methods: http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		Uri:     "/users",
		Methods: http.MethodPost,
		Handler: controllers.CreateUsers,
	},
	Route{
		Uri:     "/users/{id}",
		Methods: http.MethodPutconst,
		Handler: controllers.UpdateUsers,
	},
	Route{
		Uri:     "/users/{id}",
		Methods: http.MethodDelete,
		Handler: controllers.DeleteUsers,
	},
}
