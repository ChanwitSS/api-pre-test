package main

import (
	"net/http"
	"os"

	// "post/docs"
	"post/internal/models"
	"post/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Robinhood Interview API
// @version 0.1
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// func SetupSwagger() {
// 	var host string
// 	docs.SwaggerInfo.Host = host
// }

func init() {
	godotenv.Load(".env")
	models.Setup()
	// redis.Setup()
	// SetupSwagger()
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	routersInit := routers.InitRouter()
	// readTimeout := setting.ServerSetting.ReadTimeout
	// writeTimeout := setting.ServerSetting.WriteTimeout

	server := &http.Server{
		Addr:    ":3000",
		Handler: routersInit,
		// ReadTimeout: readTimeout,
		// WriteTimeout:   writeTimeout,
		// MaxHeaderBytes: maxHeaderBytes,
	}

	// log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
