package handler

import (
	"15min-city/dto"
	"15min-city/pkg/errs"
	"15min-city/pkg/helpers"
	"15min-city/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type datasetHandler struct {
	datasetService service.DatasetService
}

type DatasetHandler interface {
	CreateDataset(c *gin.Context)
	GetDatasetByID(c *gin.Context)
	GetDatasetByName(c *gin.Context)
	GetDatasetByCategory(c *gin.Context)
	UpdateDataset(c *gin.Context)
	DeleteDataset(c *gin.Context)
	GetAllDatasets(c *gin.Context)
}

func NewDatasetHandler(datasetService service.DatasetService) DatasetHandler {
	return &datasetHandler{
		datasetService: datasetService,
	}
}

func (d *datasetHandler) CreateDataset(c *gin.Context) {
	var datasetPayload dto.CreateDatasetRequest

	if err := c.ShouldBindJSON(&datasetPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())
		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(datasetPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := d.datasetService.CreateDataset(c.Request.Context(), datasetPayload)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (d *datasetHandler) GetDatasetByID(c *gin.Context) {
	idStr := c.Param("datasetID")

	// Convert the id from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Handle conversion error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dataset ID"})
		return
	}

	response, err := d.datasetService.GetDatasetByID(c.Request.Context(), id)
	if err != nil {
		// Assert the error type
		if customErr, ok := err.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), customErr)
		} else {
			// Handle other errors
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	c.JSON(response.Status, response)
}

func (d *datasetHandler) GetDatasetByName(c *gin.Context) {
	name := c.Param("name")

	response, err := d.datasetService.GetDatasetByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (d *datasetHandler) GetDatasetByCategory(c *gin.Context) {
	category := c.Param("category")

	datasets, err := d.datasetService.GetDatasetByCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datasets": datasets})
}

func (d *datasetHandler) UpdateDataset(c *gin.Context) {
	idStr := c.Param("datasetID")

	// Convert the id from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Handle conversion error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dataset ID"})
		return
	}

	var datasetPayload dto.UpdateDatasetRequest

	if err := c.ShouldBindJSON(&datasetPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())
		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(datasetPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := d.datasetService.UpdateDataset(c.Request.Context(), id, datasetPayload)
	if err != nil {
		// Assert the error type
		if customErr, ok := err.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), customErr)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	c.JSON(response.Status, response)
}

func (d *datasetHandler) DeleteDataset(c *gin.Context) {
	idStr := c.Param("datasetID")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	fmt.Println("Request to delete dataset with ID:", id)

	response, serviceErr := d.datasetService.DeleteDataset(c.Request.Context(), int(id))
	if serviceErr != nil {
		if customErr, ok := serviceErr.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), gin.H{"error": customErr.Message()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	fmt.Println("Dataset deleted successfully:", id)
	c.JSON(response.Status, response)
}

func (d *datasetHandler) GetAllDatasets(c *gin.Context) {
	datasets, err := d.datasetService.GetAllDatasets(c.Request.Context())
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, datasets)
}
