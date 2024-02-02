package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"integration-test/routes"
	"integration-test/utils/arangodb"
	"integration-test/utils/constant"
	"integration-test/utils/middleware"
	"os"
	"strconv"
)

func MainHttpHandler(arangoDB arangodb.ArangoDB) {
	gin.SetMode(os.Getenv(constant.EnvGinMode))

	g := gin.Default()

	g.Use(
		middleware.CORSMiddleware(),
		middleware.JSONMiddleware(),
		middleware.RequestId(),
	)

	routes.InitializeHttpRoute(g, arangoDB)
	useSSL, err := strconv.ParseBool(os.Getenv(constant.EnvUseSSL))
	addr := fmt.Sprintf(":%s", os.Getenv(constant.EnvMainPort))

	if err != nil || useSSL {
		g.RunTLS(addr, os.Getenv(constant.EnvPublicSSHPath), os.Getenv(constant.EnvPrivateSSHPath))
	} else {
		// Listen and serve on 0.0.0.0:8080 (For Windows "localhost:8080")
		fmt.Println("GO")
		g.Run(addr)
	}
}
