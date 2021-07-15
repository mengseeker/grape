package server

import (
	"grape/grapeapi/server/api/cluster"
	"grape/grapeapi/server/api/group"
	h "grape/grapeapi/server/api/helpers"
	"grape/grapeapi/server/api/login"
	"grape/grapeapi/server/api/namespace"
	"grape/grapeapi/server/api/service"
	"grape/grapeapi/server/ui"
	"grape/pkg/logger"
	"net"

	"github.com/gin-gonic/gin"
)

var (
	log = logger.NewLogger("server")
)

type ServerConfig struct {
	GinMode string
	ApiAddr string
	MCPAddr string
}

func Serve(config ServerConfig) {
	gin.SetMode(config.GinMode)
	go serveApi(config.ApiAddr)
	serveMCP(config.MCPAddr)
}

func serveApi(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("apiserver listen failed: %q", err.Error())
	}
	log.Infof("apiserver listening at %v", lis.Addr())
	err = getRouter().RunListener(lis)
	if err != nil {
		log.Fatalf("apiserver failed: %q", err.Error())
	}
}

func serveMCP(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("mcpserver listen failed: %q", err.Error())
	}
	log.Infof("mcpserver listening at %v", lis.Addr())
	err = newMCPServer().RunListener(lis)
	if err != nil {
		log.Fatalf("mcpserver failed: %q", err.Error())
	}
}

func getRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(h.Recovery))
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "health")
	})
	service.Mount(r)
	cluster.Mount(r)
	login.Mount(r)
	namespace.Mount(r)
	group.Mount(r)
	ui.Mount(r)
	return r
}
