// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"html/template"
	"log/slog"
	"os"

	"goki.dev/goki/config"
)

// AppJSTmpl is the template used to build the app.js file
var AppJSTmpl = template.Must(template.New("app.js").Parse(appJS))

// AppJSData is the data passed to AppJSTmpl
type AppJSData struct {
	Env                     string
	LoadingLabel            string
	Wasm                    string
	WasmContentLengthHeader string
	WorkerJS                string
	AutoUpdateInterval      int64
}

// MakeAppJS exectues [AppJSTmpl] based on the given configuration information.
func MakeAppJS(c *config.Config) ([]byte, error) {
	if c.Web.Env == nil {
		c.Web.Env = make(map[string]string)
	}
	c.Web.Env["GOAPP_VERSION"] = c.Version
	c.Web.Env["GOAPP_STATIC_RESOURCES_URL"] = staticPath
	c.Web.Env["GOAPP_ROOT_PREFIX"] = c.Build.Package

	for k, v := range c.Web.Env {
		if err := os.Setenv(k, v); err != nil {
			slog.Error("setting app env variable failed", "name", k, "value", "err", err)
		}
	}

	d := AppJSData{
		Env:                     jsonString(c.Web.Env),
		LoadingLabel:            c.Web.LoadingLabel,
		Wasm:                    "/app.wasm",
		WasmContentLengthHeader: c.Web.WasmContentLengthHeader,
		WorkerJS:                "/app-worker.js",
		AutoUpdateInterval:      c.Web.AutoUpdateInterval.Milliseconds(),
	}
	b := &bytes.Buffer{}
	err := AppJSTmpl.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// AppWorkerJSData is the data passed to [config.Config.Web.ServiceWorkerTemplate]
type AppWorkerJSData struct {
	Version          string
	ResourcesToCache string
}

// MakeWorkerJS executes [config.Config.Web.ServiceWorkerTemplate].
func MakeAppWorkerJS(c *config.Config) ([]byte, error) {
	resources := []string{
		"/app.css",
		"/app.js",
		"/app.wasm",
		"/manifest.webmanifest",
		"/wasm_exec.js",
		"/",
	}

	tmpl, err := template.New("app-worker.js").Parse(c.Web.ServiceWorkerTemplate)
	if err != nil {
		return nil, err
	}

	d := AppWorkerJSData{
		Version:          c.Version,
		ResourcesToCache: jsonString(resources),
	}

	b := &bytes.Buffer{}
	err = tmpl.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
