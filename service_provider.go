package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/goravel/framework/contracts/binding"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/log"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/contracts/view"
)

const BindingRoute = "goravel.gin.route"

var (
	App                 foundation.Application
	ConfigFacade        config.Config
	LogFacade           log.Log
	ValidationFacade    validation.Validation
	ViewFacade          view.View
	ResponseViewFactory func(*gin.Context) http.ResponseView
)

type ServiceProvider struct{}

func (r *ServiceProvider) Relationship() binding.Relationship {
	return binding.Relationship{
		Bindings: []string{
			BindingRoute,
		},
		Dependencies: []string{
			binding.Config,
			binding.Log,
			binding.Validation,
		},
		ProvideFor: []string{
			binding.Route,
		},
	}
}

func (r *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.BindWith(BindingRoute, func(app foundation.Application, parameters map[string]any) (any, error) {
		return NewRoute(app.MakeConfig(), app.MakeView(), parameters)
	})
}

func (r *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	LogFacade = app.MakeLog()
	ValidationFacade = app.MakeValidation()
}
