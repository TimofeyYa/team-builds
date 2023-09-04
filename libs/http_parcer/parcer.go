package httpparcer

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorHTTP struct {
	Msg  string
	Code uint16
}

func (e *ErrorHTTP) Error() string {
	return fmt.Sprintf("%s: code %d", e.Msg, e.Code)
}

func Parce[reqData any, resData any](f func(c context.Context, body reqData) (*resData, *ErrorHTTP)) func(*gin.Context) {
	return func(c *gin.Context) {
		var body reqData
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
			return
		}

		result, err := f(c, body)
		if err != nil {
			c.AbortWithStatusJSON(int(err.Code), gin.H{"status": false, "message": err.Msg})
			return
		}

		c.JSON(200, result)
	}
}
