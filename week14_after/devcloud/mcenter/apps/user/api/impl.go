package api

import (
	"github.com/infraboard/mcube/app"
)

func init() {
	app.RegistryRESTfulApp(primaryHandler)
	app.RegistryRESTfulApp(subHandler)
}
