package public

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/errors"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func ResponseError(c *gin.Context, err BusinessException) {
	resp := &Response{Success: false, Code: err.Code(), Message: err.Error(), Content: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Success: true, Code: http.StatusOK, Message: "", Content: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

// ResponseAny : do response
func ResponseAny(c *gin.Context, err interface{}, any interface{}) {
	var exception *errors.Error
	if err == nil {
		exception = errors.Parse("")
		exception.Code = http.StatusOK
	} else {
		exception = err.(*errors.Error)
	}
	resp := &Response{Success: true, Code: exception.GetCode(), Message: exception.GetDetail(), Content: any}
	if exception.Code != http.StatusOK {
		resp.Success = false
		defer c.AbortWithError(http.StatusOK, exception)
	}
	c.JSON(http.StatusOK, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
