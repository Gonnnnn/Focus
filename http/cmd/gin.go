package main

import (
	"errors"
	"focus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginWrapper struct {
	focus focus.Focus
}

type CreateRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartTimestamp int64  `json:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
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
	var createRequest CreateRequest
	err := c.BindJSON(&createRequest)
	if err != nil {
		c.String(http.StatusBadRequest, errors.New("invalid request body").Error())
		return
	}
	activity, err := g.focus.CreateActivity(createRequest.Title, createRequest.Description, createRequest.StartTimestamp, createRequest.EndTimestamp)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (g *ginWrapper) Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
