package core

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"bazooka/internal/bazooka/routes"
	"bazooka/internal/pkg/assets"
	routeConfig "bazooka/internal/pkg/config"
)

type BazookaApp struct {
	Config *BazookaConfig
	DbConn *gorm.DB

	Engine *gin.Engine
	BaseDir string
}

func (a *BazookaApp) Svc() *http.Server {
	addr := fmt.Sprintf("%s:%d", a.Config.Listen, a.Config.Port)

	svc := &http.Server{
		Addr:              addr,
		Handler:           a.Engine,
		TLSConfig:         nil,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	return svc
}

func (a *BazookaApp) RegisterRouteSet(s *routeConfig.RouteSet) {
	prefix := s.Prefix
	if "" == prefix {
		prefix = "/"
	}

	group := a.Engine.Group(prefix)
	s.Bind(group)
}


var app *BazookaApp


func GetApp() (*BazookaApp, error) {
	if nil == app {
		return nil, errors.New("app not initialized, call InitApp first.")
	}

	return app, nil
}


func InitApp(c *BazookaConfig, extra...interface{}) (*BazookaApp, error) {
	if nil == c {
		return nil, errors.New("config is nil")
	}

	if ! c.Debug {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	cwd, _ := os.Getwd()
	app = &BazookaApp{
		Config: c,
		DbConn: nil,
		Engine: gin.Default(),
		BaseDir: cwd,
	}

	for _, i := range extra{
		if nil == i {
			continue  // Ignore empty items.
		}

		switch v := i.(type) {
		case *gorm.DB:
			app.DbConn = v
		case *gin.Engine:
			app.Engine = v
		case string:
			// Maybe BaseDir
			if assets.CheckPath(v) {
				app.BaseDir = v
				continue
			}
		default:
			continue
		}
	}

	if nil == app.DbConn {
		db, err := GetDbConn()
		if nil != err {
			return nil, err
		}
		app.DbConn = db
	}

	loadRouteSets(app)

	log.Debugf("Working directory: %s", app.BaseDir)
	return app, nil
}

func loadRouteSets(a *BazookaApp) {
	a.RegisterRouteSet(&routes.PublicRouteSet)
}
