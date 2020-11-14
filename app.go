package app

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type App struct {
	LambdaApi *ginadapter.GinLambda
	Routing   *gin.Engine
}

func NewApp(r *gin.Engine) *App {
	return &App{Routing: r}
}

func (a *App) StartServer() {
	StartApi(a.Routing)
}

func (a *App) StartLambda() {
	a.LambdaApi = StartLambdaApi(a.Routing)
	lambda.Start(a.Handler)
}

func (a *App) Start() {
	if os.Getenv("API_TYPE") == "LAMBDA" {
		a.StartLambda()
	} else {
		a.StartServer()
	}
}

func (a *App) Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return a.LambdaApi.ProxyWithContext(ctx, req)
}
