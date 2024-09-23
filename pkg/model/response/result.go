package result

import (
	"github.com/gin-gonic/gin"
)

type ResultContent struct {
	message string
	data    interface{}
}

type Result struct {
	context *gin.Context
}

func NewResult(context *gin.Context) *Result {
	return &Result{context: context}
}
func (r *Result) Common(code int, message string, data interface{}) {
	res := ResultContent{message, data}
	r.context.JSON(code, gin.H{"message": res.message, "data": res.data})
}

func (r *Result) Success(message string, data interface{}) {
	r.Common(200, message, data)
}

func (r *Result) SuccessData(data interface{}) {
	r.Common(200, "success", data)
}

func (r *Result) SuccessMessage(message string) {
	r.Common(200, message, nil)
}

func (r *Result) Fail(message string, data interface{}) {
	r.Common(500, message, data)
}

func (r *Result) FailData(data interface{}) {
	r.Common(500, "fail", data)
}
func (r *Result) FailMessage(message string) {
	r.Common(500, message, nil)
}
