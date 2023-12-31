package main

import (
	"errors"
	"focus"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ginWrapper struct {
	focus focus.Focus
}

type CreateRequest struct {
	Title          string `json:"title" validate:"required"`
	Description    string `json:"description" validate:"required"`
	StartTimestampMilli int64  `json:"startTimestampMilli" validate:"required,min=1"`
	EndTimestampMilli   int64  `json:"endTimestampMilli" validate:"required,min=1"`
}

type DeleteRequest struct {
	Id string `json:"id" validate:"required"`
}

type CompleteRequest struct {
	Id string `json:"id" validate:"required"`
}

type ListResponse struct {
	Activities []*focus.Activity `json:"activities"`
}

type CreateResponse struct {
	Activity *focus.Activity `json:"activity"`
}

var validate = validator.New()

func (g *ginWrapper) List(c *gin.Context) {
	ids := c.QueryArray("ids")
	activities, err := g.focus.Activities(ids)
	if err != nil {
		log.Printf("failed to get activities: %+v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "main.html", gin.H{"Activities": activities})
}

func (g *ginWrapper) Create(c *gin.Context) {
	var createRequest CreateRequest
	err := c.BindJSON(&createRequest)
	if err != nil {
		c.String(http.StatusBadRequest, errors.New("invalid request body").Error())
		return
	}

	err = validate.Struct(createRequest)
	if err != nil {
		c.String(http.StatusBadRequest, errors.New("invalid request body").Error())
		return
	}

	activity, err := g.focus.CreateActivity(createRequest.Title, createRequest.Description, createRequest.StartTimestampMilli, createRequest.EndTimestampMilli)
	if err != nil {
		log.Printf("failed to create activity: %+v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (g *ginWrapper) Delete(c *gin.Context) {
	var deleteRequest DeleteRequest
	if err := c.BindJSON(&deleteRequest); err != nil {
		c.String(http.StatusBadRequest, errors.New("invalid request body").Error())
		return
	}

	if err := validate.Struct(deleteRequest); err != nil {
		c.String(http.StatusBadRequest, errors.New("invalid request body").Error())
		return
	}

	if err := g.focus.DeleteActivity(deleteRequest.Id); err != nil {
		log.Printf("failed to delete activity: %+v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.Status(http.StatusOK)
}

func (g *ginWrapper) Complete(c *gin.Context) {
	var completeRequest CompleteRequest
	if err := c.BindJSON(&completeRequest); err != nil {
		c.String(http.StatusBadRequest, errors.New("invalid request body").Error())
		return
	}

	if err := validate.Struct(completeRequest); err != nil {
		c.String(http.StatusBadRequest, errors.New("invalid request body").Error())
		return
	}

	activity, err := g.focus.CompleteActivity(completeRequest.Id)
	if err != nil {
		log.Printf("failed to complete activity: %+v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (g *ginWrapper) Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
