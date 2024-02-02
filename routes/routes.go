package routes

import (
	"github.com/gin-gonic/gin"
	"integration-test/app/pkg/user"
	"integration-test/utils/arangodb"
)

func InitializeHttpRoute(g *gin.Engine, arangoDB arangodb.ArangoDB) {
	user.InitializeUser(arangoDB)
	userController := user.NewHttpUserController()

	group := g.Group("api/v1")
	{
		group.GET("/user", userController.GetUser)
		group.GET("/user/:id", userController.FindUserById)
		group.POST("/user", userController.Create)
		group.PUT("/user/:id", userController.UpdateUser)
		group.DELETE("/user/:id", userController.DeleteUser)
	}
}
