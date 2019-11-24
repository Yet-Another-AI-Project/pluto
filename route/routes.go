package route

import (
	"net/http"

	"github.com/leeif/pluto/config"
	"github.com/leeif/pluto/middleware"

	"github.com/gorilla/mux"
	"github.com/leeif/pluto/log"
	"github.com/leeif/pluto/manage"

	"github.com/urfave/negroni"
)

type middle func(handlers ...negroni.HandlerFunc) http.Handler

type PlutoRoute struct {
	path        string
	description string
	method      string
	middle      middle
	handler     negroni.HandlerFunc
}

type Router struct {
	manager *manage.Manager
	config  *config.Config
	logger  *log.PlutoLog
	mw      *middleware.Middleware
	mux     *mux.Router
}

func (r *Router) registerAPIRoutes() {
	routes := []PlutoRoute{
		{
			path:        "/user/register",
			description: "register user with email",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.register,
		},
		{
			path:        "/user/register/verify/mail",
			description: "send registration verification mail",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.verifyMail,
		},
		{
			path:        "/user/login",
			description: "login with mail",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.login,
		},
		{
			path:        "/user/login/google/mobile",
			description: "login with google account for mobile app",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.googleLoginMobile,
		},
		{
			path:        "/user/login/apple/mobile",
			description: "login with apple account for mobile app",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.appleLoginMobile,
		},
		{
			path:        "/user/login/wechat/mobile",
			description: "login with wechat account for mobile app",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.wechatLoginMobile,
		},
		{
			path:        "/user/password/reset/mail",
			description: "send password reset mail",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.passwordResetMail,
		},
		{
			path:        "/user/info/me",
			description: "get user info",
			method:      "POST",
			middle:      r.mw.TokenVerifyMiddleware,
			handler:     r.userInfo,
		},
		{
			path:        "/user/info/me/update",
			description: "update user info",
			method:      "POST",
			middle:      r.mw.TokenVerifyMiddleware,
			handler:     r.updateUserInfo,
		},
		{
			path:        "/auth/refresh",
			description: "refresh access token",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.refreshToken,
		},
		{
			path:        "/auth/publickey",
			description: "get the rsa public key",
			method:      "GET",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.publicKey,
		},
		{
			path:        "/healthcheck",
			description: "health check api",
			method:      "GET",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.healthCheck,
		},
	}
	r.registerRoutes(routes, "/api")
}

func (r *Router) registerWebRoutes() {
	routes := []PlutoRoute{
		{
			path:        "/mail/verify/{token}",
			description: "verify the mail registration",
			method:      "GET",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.registrationVerifyPage,
		},
		{
			path:        "/password/reset/{token}",
			description: "reset password page",
			method:      "GET",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.resetPasswordPage,
		},
		{
			path:        "/password/reset/{token}",
			description: "reset password",
			method:      "POST",
			middle:      r.mw.NoVerifyMiddleware,
			handler:     r.resetPassword,
		},
	}
	r.registerRoutes(routes, "/")
}

func (router *Router) registerRoutes(routes []PlutoRoute, prefix string) {
	sub := router.mux.PathPrefix(prefix).Subrouter()
	for _, r := range routes {
		sub.Handle(r.path, r.middle(r.handler)).Methods(r.method)
	}
}

func (r *Router) Register() {
	r.registerAPIRoutes()
	r.registerWebRoutes()
}

func NewRouter(mux *mux.Router, manager *manage.Manager, config *config.Config, logger *log.PlutoLog) *Router {
	return &Router{
		manager: manager,
		config:  config,
		logger:  logger,
		mw:      middleware.NewMiddle(logger),
		mux:     mux,
	}
}
