package main

import (
	"net/http"

	"github.com/jmpsec/osctrl/pkg/settings"
	"github.com/jmpsec/osctrl/pkg/utils"
)

const (
	metricAPIReq    = "api-req"
	metricAPIErr    = "api-err"
	metricAPIOK     = "api-ok"
	metricHealthReq = "health-req"
	metricHealthOK  = "health-ok"
)

// JSONApplication for Content-Type headers
const JSONApplication string = "application/json"

// ContentType for header key
const contentType string = "Content-Type"

// JSONApplicationUTF8 for Content-Type headers, UTF charset
const JSONApplicationUTF8 string = JSONApplication + "; charset=UTF-8"

var errorContent = []byte("❌")
var okContent = []byte("✅")

// Handle health requests
func healthHTTPHandler(w http.ResponseWriter, r *http.Request) {
	incMetric(metricHealthReq)
	utils.DebugHTTPDump(r, settingsmgr.DebugHTTP(settings.ServiceAPI), true)
	// Send response
	apiHTTPResponse(w, JSONApplicationUTF8, http.StatusOK, okContent)
	incMetric(metricHealthOK)
}

// Handle root requests
func rootHTTPHandler(w http.ResponseWriter, r *http.Request) {
	incMetric(metricHealthReq)
	utils.DebugHTTPDump(r, settingsmgr.DebugHTTP(settings.ServiceAPI), true)
	// Send response
	apiHTTPResponse(w, JSONApplicationUTF8, http.StatusOK, okContent)
	incMetric(metricHealthOK)
}

// Handle error requests
func errorHTTPHandler(w http.ResponseWriter, r *http.Request) {
	incMetric(metricAPIReq)
	utils.DebugHTTPDump(r, settingsmgr.DebugHTTP(settings.ServiceAPI), true)
	// Send response
	apiHTTPResponse(w, JSONApplicationUTF8, http.StatusInternalServerError, errorContent)
	incMetric(metricAPIErr)
}

// Handle forbidden error requests
func forbiddenHTTPHandler(w http.ResponseWriter, r *http.Request) {
	incMetric(metricAPIReq)
	utils.DebugHTTPDump(r, settingsmgr.DebugHTTP(settings.ServiceAdmin), true)
	// Send response
	apiHTTPResponse(w, JSONApplicationUTF8, http.StatusForbidden, errorContent)
	incMetric(metricAPIErr)
}
