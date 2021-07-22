/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type AdminApi interface {

	/*
			 * Prometheus Get snapshot metrics from the service. If you're using k8s, you can then add annotations to your deployment like so:
			 * ```
		metadata:
		annotations:
		prometheus.io/port: "4434"
		prometheus.io/path: "/metrics/prometheus"
		```
			 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 * @return AdminApiApiPrometheusRequest
	*/
	Prometheus(ctx context.Context) AdminApiApiPrometheusRequest

	/*
	 * PrometheusExecute executes the request
	 */
	PrometheusExecute(r AdminApiApiPrometheusRequest) (*http.Response, error)
}

// AdminApiService AdminApi service
type AdminApiService service

type AdminApiApiPrometheusRequest struct {
	ctx        context.Context
	ApiService AdminApi
}

func (r AdminApiApiPrometheusRequest) Execute() (*http.Response, error) {
	return r.ApiService.PrometheusExecute(r)
}

/*
 * Prometheus Get snapshot metrics from the service. If you're using k8s, you can then add annotations to your deployment like so:
 * ```
metadata:
annotations:
prometheus.io/port: "4434"
prometheus.io/path: "/metrics/prometheus"
```
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AdminApiApiPrometheusRequest
*/
func (a *AdminApiService) Prometheus(ctx context.Context) AdminApiApiPrometheusRequest {
	return AdminApiApiPrometheusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *AdminApiService) PrometheusExecute(r AdminApiApiPrometheusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.Prometheus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics/prometheus"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
