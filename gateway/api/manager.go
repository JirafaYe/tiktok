package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"strings"
	"time"
)

type Manager struct {
	handler *gin.Engine
}

func New() *Manager {
	return &Manager{
		handler: gin.Default(),
	}
}

func (m *Manager) Run() error {
	err := m.load()
	if err != nil {
		return err
	}
	return m.handler.Run(C.IP + ":" + C.Port)
}

func (m *Manager) load() (err error) {
	err = m.loadPlugin()
	if err != nil {
		return
	}
	return m.loadRoute()
}

func (m *Manager) loadPlugin() error {
	m.loadCORS()
	return nil
}

func (m *Manager) loadCORS() error {
	// iris cors
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowedMethods:   []string{"POST", "GET", "OPTIONS", "DELETE"},
	//	MaxAge:           3600,
	//	AllowedHeaders:   []string{"*"},
	//	AllowCredentials: true,
	//})
	//m.handler.UseRouter(crs)

	// gin cors
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})
	m.handler.Use(crs)
	return nil
}

func (m *Manager) loadRoute() error {
	t := reflect.TypeOf(m)
	for i := 0; i < t.NumMethod(); i++ {
		f := t.Method(i)
		if strings.HasPrefix(f.Name, "Route") &&
			f.Type.NumOut() == 0 &&
			f.Type.NumIn() == 1 {
			log.Println("[GATEWAY] LOAD ROUTE:", f.Name)
			f.Func.Call([]reflect.Value{
				reflect.ValueOf(m),
			})
		}
	}
	return nil
}
