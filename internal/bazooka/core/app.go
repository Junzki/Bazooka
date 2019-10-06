package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/exp/errors"
)

type BazookaApp struct {
	Config *BazookaConfig
	DbConn *gorm.DB

	Engine *gin.Engine
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
	}

	app = &BazookaApp{
		Config: c,
		DbConn: nil,
		Engine: nil,
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

	if nil == app.Engine {
		app.Engine = gin.Default()
	}

	return app, nil
}
