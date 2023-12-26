package v1

import (
	"pluto/config"
	"pluto/log"
	"pluto/manage"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Router struct {
	manager *manage.Manager
	config  *config.Config
	logger  *log.PlutoLog
	bundle  *i18n.Bundle
}

func NewRouter(manager *manage.Manager, config *config.Config, logger *log.PlutoLog, bundle *i18n.Bundle) *Router {
	return &Router{
		manager: manager,
		config:  config,
		logger:  logger,
		bundle:  bundle,
	}
}
