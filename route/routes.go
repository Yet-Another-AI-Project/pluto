package route

import (
	"net/http"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"pluto/config"
	"pluto/middleware"
	v1 "pluto/route/v1"

	perror "pluto/datatype/pluto_error"
	"pluto/log"
	"pluto/manage"
	routeUtils "pluto/utils/route"

	"github.com/gorilla/mux"
)

type middle func(handlerWrapper middleware.HandlerWrapper, handlers ...func(http.ResponseWriter, *http.Request) *perror.PlutoError) http.Handler

type PlutoRoute struct {
	path        string
	description string
	method      string
	middle      middle
	handler     func(w http.ResponseWriter, r *http.Request) *perror.PlutoError
}

type Router struct {
	logger *log.PlutoLog
	config *config.Config
	v1     *v1.Router
	mux    *mux.Router
}

func (router *Router) registerRoutes(routes []PlutoRoute, prefix string, isWeb bool) {
	sub := router.mux.PathPrefix(prefix).Subrouter()
	for _, r := range routes {
		if isWeb {
			sub.Handle(r.path, r.middle(router.plutoWebHandlerWrapper, r.handler)).Methods(r.method)
		} else {
			sub.Handle(r.path, r.middle(router.plutoHandlerWrapper, r.handler)).Methods(r.method)
		}
	}
}

func (router *Router) RegisterV1() {
	v1Prefix := "/v1"
	router.registerWebV1Routes("")
	router.registerHealthV1Routes(v1Prefix)
	router.registerRBACV1Routes(v1Prefix)
	router.registerUserV1Routes(v1Prefix)
	router.registerTokenV1Routes(v1Prefix)
	router.registerOauthV1Routes(v1Prefix)
}

func (router *Router) Register() {
	router.RegisterV1()
	router.mux.NotFoundHandler = http.HandlerFunc(router.notFoundHandler)
}

func (router *Router) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		routeUtils.ResponseHTMLError("404.html", nil, r, w, http.StatusNotFound, router.config)
		return
	}
	routeUtils.ResponseError(perror.NotFound, w)
}

func NewRouter(mux *mux.Router, manager *manage.Manager, config *config.Config, logger *log.PlutoLog, bundle *i18n.Bundle) *Router {

	v1Router := v1.NewRouter(manager, config, logger, bundle)

	return &Router{
		logger: logger,
		config: config,
		mux:    mux,
		v1:     v1Router,
	}
}
