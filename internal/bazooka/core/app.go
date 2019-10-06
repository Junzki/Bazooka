package core

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/errors"

	"bazooka/internal/pkg/assets"
)

type BazookaApp struct {
	Config *BazookaConfig
	DbConn *gorm.DB

	Engine *gin.Engine
	BaseDir string
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

	log.Debugf("Working directory: %s", app.BaseDir)
	return app, nil
}
