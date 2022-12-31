package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/multimoml/shared/etcd/docs"
)

type configServer struct {
	UnimplementedConfigServer
}

func (s *configServer) GetConfig(_ context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	value, err := os.ReadFile("configs/" + req.Key)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ConfigResponse{Value: string(value)}, nil
}

func GetConfig(key string) (string, error) {
	value, err := os.ReadFile("configs/" + key)
	if err != nil {
		return "", err
	}

	return string(value), nil
}

func GetConfigs() (string, error) {
	// Read all files in directory
	files, err := os.ReadDir("configs")
	if err != nil {
		return "", err
	}

	// Return list of files, one per line
	var configs string
	for _, file := range files {
		configs += file.Name() + "\n"
	}

	return configs, nil
}

// @title Configuration store
// @version 1.0.0
// @host localhost:8080
func main() {
	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			panic(err)
		}

		s := grpc.NewServer()
		RegisterConfigServer(s, &configServer{})

		log.Fatal(s.Serve(lis))
	}()

	router := gin.Default()

	router.GET("/:id", Config)
	router.GET("/", Configs)
	router.GET("/live", Liveness)
	router.GET("/ready", Readiness)

	// Redirect /openapi to /openapi/index.html
	router.GET("/openapi", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/openapi/index.html")
	})
	router.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(router.Run(":8080"))
}

// Liveness returns a 200 if the microservice is live
// @Summary Get liveness status of the microservice
// @Description Get liveness status of the microservice
// @Success 200 {string} string
// @Router /live [get]
func Liveness(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Readiness returns a 200 if the microservice is ready
// @Summary Get readiness status of the microservice
// @Description Get readiness status of the microservice
// @Success 200 {string} string
// @Router /ready [get]
func Readiness(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Config returns a config
// @Summary Get a config
// @Description Get a config
// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /{id} [get]
func Config(c *gin.Context) {
	key := c.Param("id")

	// If key is not in the list of configs, return a 404
	configs, err := GetConfigs()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if !strings.Contains(configs, key) {
		c.String(http.StatusNotFound, "Config not found")
		return
	}

	value, err := GetConfig(key)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, value)
}

// Configs returns a list of configs
// @Summary Get all configs
// @Description Get all configs
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router / [get]
func Configs(c *gin.Context) {
	configs, err := GetConfigs()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, configs)
}
