package main

import (
	"fmt"
	"log"

	"github.com/cbuelvasc/backend/config"
	"github.com/cbuelvasc/backend/controller"
	_ "github.com/cbuelvasc/backend/docs"
	"github.com/cbuelvasc/backend/handler"
	"github.com/cbuelvasc/backend/repository"
	"github.com/cbuelvasc/backend/routes"
	"github.com/cbuelvasc/backend/security"
	"github.com/cbuelvasc/backend/util"
	"github.com/labstack/echo/v4"
)

var userController *controller.UserController
var tweetController *controller.TweetController

// @title Twitter REST API
// @description Provides access to the core features of Twitter REST API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/yellow/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	e.HTTPErrorHandler = handler.ErrorHandler
	e.Validator = util.NewValidationUtil()
	config.CORSConfig(e)
	security.WebSecurityConfig(e)

	security.WebSecurityConfig(e)

	routes.GetUserApiRoutes(e, userController)
	routes.GetTweetApiRoutes(e, tweetController)
	routes.GetSwaggerRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}

func init() {
	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}
	userRepository := repository.NewUserRepository(mongoConnection)
	authValidator := security.NewAuthValidator(userRepository)
	userController = controller.NewUserController(userRepository, authValidator)

	tweetRepository := repository.NewTeewtRepository(mongoConnection)
	tweetController = controller.NewTweetController(tweetRepository, userRepository)
}
