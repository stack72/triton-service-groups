package router

import (
	"net/http"

	"path"

	"github.com/joyent/triton-service-groups/groups"
	"github.com/joyent/triton-service-groups/templates"
)

var (
	urlPrefix = "/v1/tsg"
)

var templateRoutes = Routes{
	Route{
		"GetTemplate",
		http.MethodGet,
		path.Join(urlPrefix, "templates", "{name}"),
		templates_v1.Get,
	},
	Route{
		"CreateTemplate",
		http.MethodPost,
		path.Join(urlPrefix, "templates"),
		templates_v1.Create,
	},
	Route{
		"UpdateTemplate",
		http.MethodPut,
		path.Join(urlPrefix, "templates", "{name}"),
		templates_v1.Update,
	},
	Route{
		"DeleteTemplate",
		http.MethodDelete,
		path.Join(urlPrefix, "templates", "{name}"),
		templates_v1.Delete,
	},
	Route{
		"ListTemplates",
		http.MethodGet,
		path.Join(urlPrefix, "templates"),
		templates_v1.List,
	},
}

var groupRoutes = Routes{
	Route{
		"GetGroup",
		http.MethodGet,
		path.Join(urlPrefix, "{name}"),
		groups_v1.Get,
	},
	Route{
		"CreateGroup",
		http.MethodPost,
		path.Join(urlPrefix),
		groups_v1.Create,
	},
	Route{
		"UpdateGroup",
		http.MethodPut,
		path.Join(urlPrefix, "{name}"),
		groups_v1.Update,
	},
	Route{
		"DeleteGroup",
		http.MethodDelete,
		path.Join(urlPrefix, "{name}"),
		groups_v1.Delete,
	},
	Route{
		"ListGroups",
		http.MethodGet,
		path.Join(urlPrefix),
		groups_v1.List,
	},
}
