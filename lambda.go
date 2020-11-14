package app

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

func StartLambdaApi(r *gin.Engine) *ginadapter.GinLambda {
	return ginadapter.New(r)
}
