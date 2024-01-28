package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

	"pluto/localization"

	"github.com/k0kubun/pp"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"pluto/manage"
	"pluto/utils/admin"

	"pluto/server"

	plog "pluto/log"
	"pluto/route"

	"pluto/config"

	"go.uber.org/fx"

	"pluto/database"

	perror "pluto/datatype/pluto_error"
	"pluto/utils/rsa"
	"pluto/utils/view"

	_ "github.com/go-sql-driver/mysql"
)

// VERSION is the pluto build version
var VERSION = ""

func register(router *route.Router, db *sql.DB, config *config.Config, bundle *i18n.Bundle) error {

	if err := rsa.Init(config); err != nil {
		log.Fatalln(err.Error())
		return err
	}

	if err := view.InitView(config); err != nil {
		log.Fatalln(err.Error())
		return err
	}

	if err := admin.Init(db, config, bundle); err != nil {
		if err.PlutoCode == perror.ServerError.PlutoCode {
			log.Fatalln(err.LogError.Error())
			return err.LogError
		}
		log.Print(err)
	}

	// register routes
	router.Register()

	return nil
}

func main() {
	_ = pp.Println // prevent warning

	app := fx.New(
		fx.Provide(
			func() []string {
				return os.Args
			},
			func() string {
				return VERSION
			},
			config.NewConfig,
			database.NewDatabase,
			plog.NewLogger,
			server.NewMux,
			route.NewRouter,
			localization.NewBundle,
			manage.NewManager,
		),
		fx.Invoke(register),
		fx.NopLogger,
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
