package middleware

import (
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		userData, err := helpers.GetUserData(bearerToken)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		c.Set("userData", userData)
		c.Next()
	}
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, isOK := c.MustGet("userData").(*entity.User)

		if !isOK {
			unauthenticatedErr := errs.NewUnauthenticatedError("You are not authenticated")

			c.AbortWithStatusJSON(unauthenticatedErr.Status(), unauthenticatedErr)
			return
		}

		role := userData.Role

		if role != "admin" {
			unauthorizedErr := errs.NewUnauthorizedError("You are not authorized")

			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		c.Next()
	}
}
