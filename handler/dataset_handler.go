package handler

import (
	"15min-city/dto"
	"15min-city/pkg/errs"
	"15min-city/pkg/helpers"
	"15min-city/service"
	"encoding/csv"
	"io"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type datasetHandler struct {
	datasetService service.DatasetService
}

type DatasetHandler interface {
	CreateDataset(c *gin.Context)
	GetDatasetByID(c *gin.Context)
	GetDatasetByName(c *gin.Context)
	GetDatasetByKecamatan(c *gin.Context)
	GetDatasetByKelurahan(c *gin.Context)
	GetDatasetByCategory(c *gin.Context)
	UpdateDataset(c *gin.Context)
	DeleteDataset(c *gin.Context)
	GetAllDatasets(c *gin.Context)
	UploadDatasets(c *gin.Context)
	GetDatasetsByDistance(c *gin.Context)
}

func NewDatasetHandler(datasetService service.DatasetService) DatasetHandler {
	return &datasetHandler{
		datasetService: datasetService,
	}
}

// Fungsi untuk upload file (CSV/Excel)
func (d *datasetHandler) UploadDatasets(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File diperlukan"})
		return
	}

	tempFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tidak dapat membuka file"})
		return
	}
	defer tempFile.Close()

	ext := filepath.Ext(file.Filename)
	if ext == ".csv" {
		d.processCSV(c, tempFile)
	} else if ext == ".xlsx" {
		d.processExcel(c, tempFile)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung. Silakan unggah file CSV atau Excel."})
	}
}

// Fungsi untuk memproses CSV
func (d *datasetHandler) processCSV(c *gin.Context, file io.Reader) {
	reader := csv.NewReader(file)
	var createdDatasets []dto.CreateDatasetResponse

	// Lewati header CSV
	if _, err := reader.Read(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read CSV header"})
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading CSV file"})
			return
		}

		// Konversi record menjadi struct
		datasetPayload := dto.CreateDatasetRequest{
			Name:      record[0],
			Latitude:  parseToFloat(record[1]),
			Longitude: parseToFloat(record[2]),
			Category:  record[3],
		}

		response, err := d.datasetService.CreateDataset(c.Request.Context(), datasetPayload)
		if err != nil {
			if customErr, ok := err.(errs.ErrMessage); ok {
				c.JSON(customErr.Status(), customErr)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
			}
			return
		}
		createdDatasets = append(createdDatasets, *response)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Seluruh dataset berhasil ditambahkan",
		"data":    createdDatasets,
	})
}

func (d *datasetHandler) processExcel(c *gin.Context, file io.Reader) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel file"})
		return
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel sheet"})
		return
	}

	var createdDatasets []dto.CreateDatasetResponse

	// Lewati header Excel
	for _, row := range rows[1:] {
		datasetPayload := dto.CreateDatasetRequest{
			Name:      row[0],
			Latitude:  parseToFloat(row[1]),
			Longitude: parseToFloat(row[2]),
			Category:  row[3],
		}

		response, err := d.datasetService.CreateDataset(c.Request.Context(), datasetPayload)
		if err != nil {
			c.JSON(err.Status(), err)
			return
		}
		createdDatasets = append(createdDatasets, *response)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Seluruh dataset berhasil ditambahkan",
		"data":    createdDatasets,
	})
}

// Helper untuk mengonversi string ke float64
func parseToFloat(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return f
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

	// Tambahkan pesan sukses ke dalam respons
	c.JSON(http.StatusOK, gin.H{
		"message": "Dataset berhasil ditambahkan",
		"data":    response,
	})
}

func (d *datasetHandler) GetDatasetByID(c *gin.Context) {
    idStr := c.Param("datasetID")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid dataset ID",
        })
        return
    }

    response, err := d.datasetService.GetDatasetByID(c.Request.Context(), id)
    if err != nil {
        if customErr, ok := err.(errs.ErrMessage); ok {
            c.JSON(customErr.Status(), gin.H{
                "status":  customErr.Status(),
                "message": customErr.Message(),
            })
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{
                "status":  http.StatusInternalServerError,
                "message": "An unexpected error occurred",
            })
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Dataset found successfully",
        "data":    response,
    })
}

func (d *datasetHandler) GetDatasetByName(c *gin.Context) {
	name := c.Param("name")
	datasets, err := d.datasetService.GetDatasetByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Dataset berhasil didapatkan",
		"data":    datasets,
	})
}

func (d *datasetHandler) GetDatasetByKecamatan(c *gin.Context) {
	kecamatan := c.Param("kecamatan")
	datasets, err := d.datasetService.GetDatasetByKecamatan(c.Request.Context(), kecamatan)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seluruh dataset berhasil didapatkan",
		"data":    datasets,
	})
}

func (d *datasetHandler) GetDatasetByKelurahan(c *gin.Context) {
	kelurahan := c.Param("kelurahan")
	datasets, err := d.datasetService.GetDatasetByKelurahan(c.Request.Context(), kelurahan)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seluruh dataset berhasil didapatkan",
		"data":    datasets,
	})
}

func (d *datasetHandler) GetDatasetByCategory(c *gin.Context) {
	category := c.Param("category")
	datasets, err := d.datasetService.GetDatasetByCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seluruh dataset berhasil didapatkan",
		"data":    datasets,
	})
}

func (d *datasetHandler) UpdateDataset(c *gin.Context) {
	idStr := c.Param("datasetID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
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
		if customErr, ok := err.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), customErr)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Dataset berhasil diperbaharui",
		"data":    response,
	})
}

func (d *datasetHandler) DeleteDataset(c *gin.Context) {
	idStr := c.Param("datasetID")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	response, serviceErr := d.datasetService.DeleteDataset(c.Request.Context(), int(id))
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), gin.H{"error": serviceErr.Message()})
		return
	}

	c.JSON(response.Status, gin.H{
		"message": "Dataset berhasil dihapus",
		"data":    response,
	})
}

func (d *datasetHandler) GetAllDatasets(c *gin.Context) {
	datasets, err := d.datasetService.GetAllDatasets(c.Request.Context())
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seluruh dataset berhasil didapatkan",
		"data":    datasets,
	})
}

func (d *datasetHandler) GetDatasetsByDistance(c *gin.Context) {
    latitude, err := strconv.ParseFloat(c.DefaultQuery("latitude", "0"), 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid latitude parameter",
        })
        return
    }

    longitude, err := strconv.ParseFloat(c.DefaultQuery("longitude", "0"), 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid longitude parameter",
        })
        return
    }

    distance, err := strconv.ParseFloat(c.DefaultQuery("distance", "0"), 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid distance parameter",
        })
        return
    }

    req := dto.DatasetDistanceRequest{
        Latitude:  latitude,
        Longitude: longitude,
        Distance:  distance,
    }

    datasets, err := d.datasetService.GetDatasetsByDistance(c.Request.Context(), req)
    if err != nil {
		if customErr, ok := err.(errs.ErrMessage); ok {
			c.JSON(customErr.Status(), gin.H{
				"status":  customErr.Status(),
				"message": customErr.Message(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "An unexpected error occurred",
			})
		}
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Datasets within distance range retrieved successfully",
        "data":    datasets,
    })
}
