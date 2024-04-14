package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Celesca/isekai-shop-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

// ไม่มีการ mock server จึงใช้เป็น struct ได้เลย
type echoServer struct {
	app  *echo.Echo
	db   *gorm.DB
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoServer {
	echoApp := echo.New()
	// Logger - เราแสดง Level ที่มันต่ำสุดและไล่ลงไปเรื่อยๆ ดังนั้นมันโชว์ได้ตั้งแต่ DEBUG ถึง fatalLevel
	echoApp.Logger.SetLevel(log.DEBUG)

	// Singleton ให้กับ Server
	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})

	return server

}

func (s *echoServer) Start() {

	s.app.GET("/v1/health", s.healthCheck)

	s.httpListening()

}

// Listening to HTTP Protocol for the HTTPServe
func (s *echoServer) httpListening() {
	// localhost เป็นอะไรก็จะเป็น IP ตัวนั้นอัตโนมัติ
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
