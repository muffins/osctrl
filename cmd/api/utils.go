package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmpsec/osctrl/pkg/settings"
)

// Helper to send metrics if it is enabled
func incMetric(name string) {
	if _metrics != nil && settingsmgr.ServiceMetrics(settings.ServiceAPI) {
		_metrics.Inc(name)
	}
}

// Helper to refresh the environments map until cache/Redis support is implemented
func refreshEnvironments() {
	log.Printf("Refreshing environments...\n")
	var err error
	envsmap, err = envs.GetMap()
	if err != nil {
		log.Printf("error refreshing environments %v\n", err)
	}
}

// Helper to refresh the settings until cache/Redis support is implemented
func refreshSettings() {
	log.Printf("Refreshing settings...\n")
	var err error
	settingsmap, err = settingsmgr.GetMap(settings.ServiceAPI)
	if err != nil {
		log.Printf("error refreshing settings %v\n", err)
	}
}

// Usage for service binary
func apiUsage() {
	fmt.Printf("NAME:\n   %s - %s\n\n", serviceName, serviceDescription)
	fmt.Printf("USAGE: %s [global options] [arguments...]\n\n", serviceName)
	fmt.Printf("VERSION:\n   %s\n\n", serviceVersion)
	fmt.Printf("DESCRIPTION:\n   %s\n\n", appDescription)
	fmt.Printf("GLOBAL OPTIONS:\n")
	flag.PrintDefaults()
	fmt.Printf("\n")
}

// Display binary version
func apiVersion() {
	fmt.Printf("%s v%s\n", serviceName, serviceVersion)
	os.Exit(0)
}

// Helper to compose paths for API
func _apiPath(target string) string {
	return apiPrefixPath + apiVersionPath + target
}

// Helper to verify if a platform is valid
func checkValidPlatform(platform string) bool {
	platforms, err := nodesmgr.GetAllPlatforms()
	if err != nil {
		return false
	}
	for _, p := range platforms {
		if p == platform {
			return true
		}
	}
	return false
}

// Helper to remove duplicates from []string
func removeStringDuplicates(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	i := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[i] = v
		i++
	}
	return s[:i]
}

// Helper to send HTTP response
func apiHTTPResponse(w http.ResponseWriter, cType string, code int, data interface{}) {
	if cType != "" {
		w.Header().Set(contentType, cType)
	}
	content, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error serializing response: %v", err)
		content = []byte("error serializing response")
	}
	w.WriteHeader(code)
	_, _ = w.Write(content)
}

// Helper to handle API error responses
func apiErrorResponse(w http.ResponseWriter, msg string, code int, err error) {
	log.Printf("%s: %v", msg, err)
	apiHTTPResponse(w, JSONApplicationUTF8, code, ApiErrorResponse{Error: msg})
}

// Helper to generate a random query name
func generateQueryName() string {
	return "query_" + randomForNames()
}

// Helper to generate a random carve name
func generateCarveName() string {
	return "carve_" + randomForNames()
}

// Helper to generate a random MD5 to be used with queries/carves
func randomForNames() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	hasher := md5.New()
	_, _ = hasher.Write([]byte(fmt.Sprintf("%x", b)))
	return hex.EncodeToString(hasher.Sum(nil))
}
