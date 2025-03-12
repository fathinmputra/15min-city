package handler

import (
	"15min-city/dto"
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/pkg/helpers"
	"15min-city/service"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ResetPassword(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateImage(c *gin.Context)
	GetImageByUser(c *gin.Context)
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) Register(c *gin.Context) {
	var userPayload dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&userPayload); err != nil {
		bindErr := errs.NewUnprocessableEntityError(err.Error())

		c.JSON(bindErr.Status(), bindErr)
		return
	}

	if err := helpers.ValidateStruct(userPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := u.userService.Register(userPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) GetUserByID(c *gin.Context) {
	userData, exists := c.Get("userData")
	if !exists {
		err := errs.NewBadRequestError("User data not found")
		c.JSON(err.Status(), err)
		return
	}
	user, ok := userData.(*entity.User)
	if !ok {
		err := errs.NewInternalServerError("User data is of invalid type")
		c.JSON(err.Status(), err)
		return
	}
	userID := int(user.ID)

	response, err := u.userService.GetUserByID(userID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, gin.H{
		"status":  response.Status,
		"message": "Pengguna ditemukan",
		"data":    response,
	})
}

func (u *userHandler) Login(c *gin.Context) {
	var userCredentials dto.LoginUserRequest

	if bindErr := c.ShouldBindJSON(&userCredentials); bindErr != nil {
		err := errs.NewUnprocessableEntityError("Invalid credentials")

		c.JSON(err.Status(), err)
		return
	}

	if err := helpers.ValidateStruct(userCredentials); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := u.userService.Login(userCredentials)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, gin.H{
		"message": "Login Berhasil",
		"data":    response,
	})
}

func (u *userHandler) Logout(c *gin.Context) {
	c.Set("userData", nil)

	response := dto.LogoutResponse{
		Status:  http.StatusOK,
		Message: "You have successfully logged out",
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) ResetPassword(c *gin.Context) {
	var userPayload dto.ResetPasswordRequest

	if bindErr := c.ShouldBindJSON(&userPayload); bindErr != nil {
		err := errs.NewBadRequestError(bindErr.Error())

		c.JSON(err.Status(), err)
		return
	}

	if err := helpers.ValidateStruct(userPayload); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	response, err := u.userService.ResetPassword(userPayload)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.Status, response)
}

func (u *userHandler) CreateImage(c *gin.Context) {
	userData, exists := c.Get("userData")
	if !exists {
		err := errs.NewBadRequestError("User data not found")
		c.JSON(err.Status(), err)
		return
	}
	user, ok := userData.(*entity.User)
	if !ok {
		err := errs.NewInternalServerError("User data is of invalid type")
		c.JSON(err.Status(), err)
		return
	}
	userID := int(user.ID)

	form, err := c.MultipartForm()
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusBadRequest, "Failed to parse form data")
		return
	}

	files := form.File["image"]

	image := dto.CreateDokumenRequest{
		FormFile: files,
	}

	exampleResponse, err := u.userService.CreateImage(c.Request.Context(), userID, image)

	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusInternalServerError, "Failed to create image profile")
		return
	}

	response := helpers.WebResponse{
		Error:   false,
		Message: "OK",
		Data:    exampleResponse,
	}

	helpers.WriteToResponseBody(c, response)
}

func (u *userHandler) GetImageByUser(c *gin.Context) {
	userData, exists := c.Get("userData")
	if !exists {
		err := errs.NewBadRequestError("User data not found")
		c.JSON(err.Status(), err)
		return
	}
	user, ok := userData.(*entity.User)
	if !ok {
		err := errs.NewInternalServerError("User data is of invalid type")
		c.JSON(err.Status(), err)
		return
	}
	userID := int(user.ID)

	IDf := c.Param("fileID")
	IDfInt, err := strconv.Atoi(IDf)
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusBadRequest, "Invalid Dokumen Engineering ID")
		return
	}

	response := u.userService.GetImageByUser(c.Request.Context(), userID, IDfInt)
	if response == nil {
		helpers.WriteErrorResponse(c, http.StatusNotFound, "Image profile not found")
		return
	}
	filePath := "public/imageprofile/" + response.DokumenPath
	fmt.Println("Ini directory", filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		helpers.WriteErrorResponse(c, http.StatusNotFound, "File not found")
		return
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		helpers.WriteErrorResponse(c, http.StatusInternalServerError, "Failed to read image file")
		return
	}

	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(fileBytes)
}
