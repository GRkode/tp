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
	"github.com/gin-gonic/gin"
)

type GetFilmAPI struct {
	// Get /films/:filmID
	// Get details of a specific film
	FilmsFilmIDGet gin.HandlerFunc
}
