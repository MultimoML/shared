package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type configServer struct {
	UnimplementedConfigServer
}

func (s *configServer) GetConfig(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	value, err := os.ReadFile("configs/" + req.Key)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ConfigResponse{Value: string(value)}, nil
}

func GetConfig(key string) (string, error) {
	value, err := os.ReadFile(key)
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

	router.GET("/config/:id", func(c *gin.Context) {
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
	})

	router.GET("/configs", func(c *gin.Context) {
		configs, err := GetConfigs()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(http.StatusOK, configs)
	})

	log.Fatal(router.Run(":8080"))
}
