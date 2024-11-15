package handler

import (
	"15min-city/dto"
	"15min-city/pkg/errs"
	"15min-city/pkg/helpers"
	"15min-city/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type corridorRouteHandler struct {
	corridorRouteService service.CorridorRouteService
}

type CorridorRouteHandler interface {
	CreateCorridorRoute(c *gin.Context)
	GetCorridorRouteByID(c *gin.Context)
	GetCorridorRouteByName(c *gin.Context)
	UpdateCorridorRoute(c *gin.Context)
	DeleteCorridorRoute(c *gin.Context)
	GetAllCorridorRoutes(c *gin.Context)
}

func NewCorridorRouteHandler(corridorRouteService service.CorridorRouteService) CorridorRouteHandler {
	return &corridorRouteHandler{
		corridorRouteService: corridorRouteService,
	}
}

func (r *corridorRouteHandler) CreateCorridorRoute(c *gin.Context) {
	var corridorRoutePayload dto.CreateCorridorRouteRequest

	if err := c.ShouldBindJSON(&corridorRoutePayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())
		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(corridorRoutePayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := r.corridorRouteService.CreateCorridorRoute(c.Request.Context(), corridorRoutePayload)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (r *corridorRouteHandler) GetCorridorRouteByID(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id from string to integer
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	response, err := r.corridorRouteService.GetCorridorRouteByID(c.Request.Context(), uint(id))
	if err != nil {
		if customErr, ok := err.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), customErr)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	c.JSON(response.Status, response)
}

func (r *corridorRouteHandler) GetCorridorRouteByName(c *gin.Context) {
	name := c.Param("name")

	response, err := r.corridorRouteService.GetCorridorRouteByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

func (r *corridorRouteHandler) UpdateCorridorRoute(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id from string to integer
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var corridorRoutePayload dto.UpdateCorridorRouteRequest

	if err := c.ShouldBindJSON(&corridorRoutePayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())
		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(corridorRoutePayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := r.corridorRouteService.UpdateCorridorRoute(c.Request.Context(), uint(id), corridorRoutePayload)
	if err != nil {
		if customErr, ok := err.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), customErr)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	c.JSON(response.Status, response)
}

func (r *corridorRouteHandler) DeleteCorridorRoute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	response, serviceErr := r.corridorRouteService.DeleteCorridorRoute(c.Request.Context(), uint(id))
	if serviceErr != nil {
		if customErr, ok := serviceErr.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), gin.H{"error": customErr.Message()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(response.Status, response)
}

func (r *corridorRouteHandler) GetAllCorridorRoutes(c *gin.Context) {
	response, err := r.corridorRouteService.GetAllCorridorRoutes(c.Request.Context())
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}
