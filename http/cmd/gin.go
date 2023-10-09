package main

import (
	"focus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginWrapper struct {
	focus focus.Focus
}

type ListResponse struct {
	Activities []*focus.Activity `json:"activities"`
}

type CreateResponse struct {
	Activity *focus.Activity `json:"activity"`
}

func (g *ginWrapper) List(c *gin.Context) {
	activities, err := g.focus.Activities([]string{})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, ListResponse{Activities: activities})
}

func (g *ginWrapper) Create(c *gin.Context) {
	activity, err := g.focus.CreateActivity("title", "description", 0)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (g *ginWrapper) Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
