package hgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

type Http struct {
	*gin.Engine
	listen net.Listener
	Config Config
}

func (h *Http) Init() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", h.Config.Host, h.Config.Port))
	if err != nil {
		panic(err)
	}

	h.listen = listen

}

func (h *Http) Run() {
	go func() {
		h.Engine.RunListener(h.listen)
	}()
}

func New(cfg Config) *Http {
	engine := gin.New()
	return &Http{
		Engine: engine,
		Config: cfg,
	}
}
