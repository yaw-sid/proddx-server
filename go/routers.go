/*
 * Proddx API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.Handler(route.Method, route.Pattern, handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CompaniesGet",
		strings.ToUpper("Get"),
		"/companies",
		CompaniesGet,
	},

	Route{
		"CompaniesIdDelete",
		strings.ToUpper("Delete"),
		"/companies/{id}",
		CompaniesIdDelete,
	},

	Route{
		"CompaniesIdGet",
		strings.ToUpper("Get"),
		"/companies/{id}",
		CompaniesIdGet,
	},

	Route{
		"CompaniesIdPut",
		strings.ToUpper("Put"),
		"/companies/{id}",
		CompaniesIdPut,
	},

	Route{
		"CompaniesPost",
		strings.ToUpper("Post"),
		"/companies",
		CompaniesPost,
	},

	Route{
		"ProductsGet",
		strings.ToUpper("Get"),
		"/products",
		ProductsGet,
	},

	Route{
		"ProductsIdDelete",
		strings.ToUpper("Delete"),
		"/products/{id}",
		ProductsIdDelete,
	},

	Route{
		"ProductsIdGet",
		strings.ToUpper("Get"),
		"/products/{id}",
		ProductsIdGet,
	},

	Route{
		"ProductsIdPut",
		strings.ToUpper("Put"),
		"/products/{id}",
		ProductsIdPut,
	},

	Route{
		"ProductsPost",
		strings.ToUpper("Post"),
		"/products",
		ProductsPost,
	},

	Route{
		"ReviewsGet",
		strings.ToUpper("Get"),
		"/reviews",
		ReviewsGet,
	},

	Route{
		"ReviewsIdDelete",
		strings.ToUpper("Delete"),
		"/reviews/{id}",
		ReviewsIdDelete,
	},

	Route{
		"ReviewsIdGet",
		strings.ToUpper("Get"),
		"/reviews/{id}",
		ReviewsIdGet,
	},

	Route{
		"ReviewsIdPut",
		strings.ToUpper("Put"),
		"/reviews/{id}",
		ReviewsIdPut,
	},

	Route{
		"ReviewsPost",
		strings.ToUpper("Post"),
		"/reviews",
		ReviewsPost,
	},
}
