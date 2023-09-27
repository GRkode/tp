/*
 * Film API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	router := gin.Default()
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the CreateFilmAPI part of the API
	CreateFilmAPI CreateFilmAPI
	// Routes for the GetFilmAPI part of the API
	GetFilmAPI GetFilmAPI
	// Routes for the GetFilmsAPI part of the API
	GetFilmsAPI GetFilmsAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
	
		{
			"AddFilm",
			http.MethodPost,
			"/films",
			handleFunctions.CreateFilmAPI.AddFilm,
		},
		{
			"FilmsFilmIDGet",
			http.MethodGet,
			"/films/:filmID",
			handleFunctions.GetFilmAPI.FilmsFilmIDGet,
		},
		{
			"GetFilms",
			http.MethodGet,
			"/films",
			handleFunctions.GetFilmsAPI.GetFilms,
		},
	}
}