package router

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	v1 "tt/controller/api/v1"
	"tt/docs"
	"tt/model"
)

const prefix = "/api/v1/"

var Echo *echo.Echo

// @title Swagger Users API
// @version 1.0
// @description This is a rest-api server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func init() {
	Echo = echo.New()
	Echo.Validator = &model.CustomValidator{Validator: validator.New()}
	//Swagger
	docs.SwaggerInfo.BasePath = prefix

	Echo.GET(prefix+"swagger/*", echoSwagger.WrapHandler)

	usersEndpoints := Echo.Group(prefix + "users")
	usersEndpoints.POST("", v1.CreateUser)
	usersEndpoints.PUT("", v1.UpdateUser)
	usersEndpoints.DELETE("", v1.DeleteUser)
	usersEndpoints.GET("", v1.SearchUsersOrdered)
}
