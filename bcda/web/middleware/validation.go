package middleware

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	responseutils "github.com/CMSgov/bcda-app/bcda/responseutils"
	responseutilsv2 "github.com/CMSgov/bcda-app/bcda/responseutils/v2"
)

var supportedOutputFormats = map[string]struct{}{
	"ndjson":                  {},
	"application/fhir+ndjson": {},
	"application/ndjson":      {}}

type RequestParameters struct {
	Since         time.Time
	ResourceTypes []string
	Version       string // e.g. v1, v2
}

// requestkey is an unexported context key to avoid collisions
type requestkey int

const rk requestkey = 0

func NewRequestParametersContext(ctx context.Context, rp RequestParameters) context.Context {
	return context.WithValue(ctx, rk, rp)
}

func RequestParametersFromContext(ctx context.Context) (RequestParameters, bool) {
	rp, ok := ctx.Value(rk).(RequestParameters)
	return rp, ok
}

// ValidateRequestURL ensure that request matches certain expectations.
// Any error that it finds will result in a http.StatusBadRequest response.
// If successful, it populates the request context with RequestParameters that can be used downstream.
// These paramters can be retrieved by calling RequestParametersFromContext.
func ValidateRequestURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rp RequestParameters

		// Get API version
		version, err := getVersion(r.URL.Path)
		if err != nil {
			// If we cannot discern the version, we cannot discern what fhir response to build so we default.
			// This should never come up as we require a version in the route for all that use this middleware.
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rp.Version = version

		rw, err := getRespWriter(version)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//validate "_outputFormat" parameter
		params, ok := r.URL.Query()["_outputFormat"]
		if ok {
			if _, found := supportedOutputFormats[params[0]]; !found {
				errMsg := fmt.Sprintf("_outputFormat parameter must be one of %v", getKeys(supportedOutputFormats))
				rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, errMsg)
				return
			}
		}

		// we do not support "_elements" parameter
		_, ok = r.URL.Query()["_elements"]
		if ok {
			rw.Exception(w, http.StatusBadRequest, responseutils.RequestErr, "Invalid parameter: this server does not support the _elements parameter.")
			return
		}

		// Check and see if the user has a duplicated the query parameter symbol (?)
		// e.g. /api/v1/Patient/$export?_type=ExplanationOfBenefit&?_since=2020-09-13T08:00:00.000-05:00
		for key := range r.URL.Query() {
			if strings.HasPrefix(key, "?") {
				rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, "Invalid parameter: query parameters cannot start with ?")
				return
			}
		}

		// validate optional "_since" parameter
		params, ok = r.URL.Query()["_since"]
		if ok {
			sinceDate, err := time.Parse(time.RFC3339Nano, params[0])
			if err != nil {
				rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, "Invalid date format supplied in _since parameter.  Date must be in FHIR Instant format.")
				return
			} else if sinceDate.After(time.Now()) {
				rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, "Invalid date format supplied in _since parameter. Date must be a date that has already passed")
				return
			}
			rp.Since = sinceDate
		}

		// validate no duplicate resource types
		params, ok = r.URL.Query()["_type"]
		if ok {
			resourceMap := make(map[string]struct{})
			resourceTypes := strings.Split(params[0], ",")
			for _, resource := range resourceTypes {
				if _, ok := resourceMap[resource]; !ok {
					resourceMap[resource] = struct{}{}
				} else {
					errMsg := fmt.Sprintf("Repeated resource type %s", resource)
					rw.Exception(w, http.StatusBadRequest, responseutils.RequestErr, errMsg)
					return
				}
			}
			rp.ResourceTypes = resourceTypes
		}

		ctx := NewRequestParametersContext(r.Context(), rp)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateRequestHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header

		acceptHeader := h.Get("Accept")
		preferHeader := h.Get("Prefer")

		// Get API version
		version, err := getVersion(r.URL.Path)
		if err != nil {
			// If we cannot discern the version, we cannot discern what fhir response to build so we default.
			// This should never come up as we require a version in the route for all that use this middleware.
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rw, err := getRespWriter(version)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if acceptHeader == "" {
			rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, "Accept header is required")
			return
		} else if acceptHeader != "application/fhir+json" {
			rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, "application/fhir+json is the only supported response format")
			return
		}

		if preferHeader == "" {
			rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, "Prefer header is required")
			return
		} else if preferHeader != "respond-async" {
			rw.Exception(w, http.StatusBadRequest, responseutils.FormatErr, "Only asynchronous responses are supported")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getKeys(kv map[string]struct{}) []string {
	keys := make([]string, 0, len(kv))
	for k := range kv {
		keys = append(keys, k)
	}
	return keys
}

var versionExp = regexp.MustCompile(`\/api\/(v\d+)\/`)

func getVersion(path string) (string, error) {
	parts := versionExp.FindStringSubmatch(path)
	if len(parts) != 2 {
		return "", fmt.Errorf("cannot retrieve version")
	}
	return parts[1], nil
}

type fhirResponseWriter interface {
	Exception(http.ResponseWriter, int, string, string)
	NotFound(http.ResponseWriter, int, string, string)
}

func getRespWriter(version string) (fhirResponseWriter, error) {
	switch version {
	case "v1":
		return responseutils.NewResponseWriter(), nil
	case "v2":
		return responseutilsv2.NewResponseWriter(), nil
	default:
		return nil, fmt.Errorf("unexpected API version: %s", version)
	}
}
