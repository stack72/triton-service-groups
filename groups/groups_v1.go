//  Copyright (c) 2018, Joyent, Inc. All rights reserved.
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at http://mozilla.org/MPL/2.0/.

package groups_v1

import (
	"net/http"

	"github.com/joyent/triton-service-groups/session"
)

type ServiceGroup struct {
	GroupName           string
	TemplateId          string
	AccountName         string
	Capacity            int
	DataCenter          []string
	HealthCheckInterval int //default will be 300
	InstanceTags        map[string]interface{}
}

func Get(session *session.TsgSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func Create(session *session.TsgSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func Update(session *session.TsgSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func Delete(session *session.TsgSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func List(session *session.TsgSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
