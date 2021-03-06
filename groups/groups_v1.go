package groups_v1

import (
	"net/http"
)

type ServiceGroup struct {
	GroupName           string
	TemplateName        string
	Capacity            int
	DataCenter          []string
	HealthCheckInterval int //default will be 300
	InstanceTags        map[string]interface{}
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
