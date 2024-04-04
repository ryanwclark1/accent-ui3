/*
accent-auth

Accent's authentication service

API version: 0.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SessionsAPI interface {

	/*
		DeleteSession Delete a session

		**Required ACL**: `auth.sessions.{session_uuid}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sessionUuid The UUID of the session
		@return SessionsAPIDeleteSessionRequest
	*/
	DeleteSession(ctx context.Context, sessionUuid string) SessionsAPIDeleteSessionRequest

	// DeleteSessionExecute executes the request
	DeleteSessionExecute(r SessionsAPIDeleteSessionRequest) (*http.Response, error)

	/*
		GetSessions List sessions

		**Required ACL:** `auth.sessions.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return SessionsAPIGetSessionsRequest
	*/
	GetSessions(ctx context.Context) SessionsAPIGetSessionsRequest

	// GetSessionsExecute executes the request
	//  @return GetSessionsResult
	GetSessionsExecute(r SessionsAPIGetSessionsRequest) (*GetSessionsResult, *http.Response, error)

	/*
		GetUserSessions Retrieves the list of sessions associated to a user

		**Required ACL**: `auth.users.{user_uuid}.sessions.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userUuid The UUID of the user
		@return SessionsAPIGetUserSessionsRequest
	*/
	GetUserSessions(ctx context.Context, userUuid string) SessionsAPIGetUserSessionsRequest

	// GetUserSessionsExecute executes the request
	//  @return GetSessionsResult
	GetUserSessionsExecute(r SessionsAPIGetUserSessionsRequest) (*GetSessionsResult, *http.Response, error)

	/*
		UserDeleteSession Delete a session

		**Required ACL**: `auth.users.{user_uuid}.sessions.{session_uuid}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userUuid The UUID of the user
		@param sessionUuid The UUID of the session
		@return SessionsAPIUserDeleteSessionRequest
	*/
	UserDeleteSession(ctx context.Context, userUuid string, sessionUuid string) SessionsAPIUserDeleteSessionRequest

	// UserDeleteSessionExecute executes the request
	UserDeleteSessionExecute(r SessionsAPIUserDeleteSessionRequest) (*http.Response, error)
}

// SessionsAPIService SessionsAPI service
type SessionsAPIService service

type SessionsAPIDeleteSessionRequest struct {
	ctx         context.Context
	ApiService  SessionsAPI
	sessionUuid string
}

func (r SessionsAPIDeleteSessionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSessionExecute(r)
}

/*
DeleteSession Delete a session

**Required ACL**: `auth.sessions.{session_uuid}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sessionUuid The UUID of the session
	@return SessionsAPIDeleteSessionRequest
*/
func (a *SessionsAPIService) DeleteSession(ctx context.Context, sessionUuid string) SessionsAPIDeleteSessionRequest {
	return SessionsAPIDeleteSessionRequest{
		ApiService:  a,
		ctx:         ctx,
		sessionUuid: sessionUuid,
	}
}

// Execute executes the request
func (a *SessionsAPIService) DeleteSessionExecute(r SessionsAPIDeleteSessionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsAPIService.DeleteSession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{session_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"session_uuid"+"}", url.PathEscape(parameterValueToString(r.sessionUuid, "sessionUuid")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SessionsAPIGetSessionsRequest struct {
	ctx          context.Context
	ApiService   SessionsAPI
	accentTenant *string
	recurse      *bool
	limit        *int32
	offset       *int32
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r SessionsAPIGetSessionsRequest) AccentTenant(accentTenant string) SessionsAPIGetSessionsRequest {
	r.accentTenant = &accentTenant
	return r
}

// Should the query include sub-tenants
func (r SessionsAPIGetSessionsRequest) Recurse(recurse bool) SessionsAPIGetSessionsRequest {
	r.recurse = &recurse
	return r
}

// The limit defines the number of individual objects that are returned
func (r SessionsAPIGetSessionsRequest) Limit(limit int32) SessionsAPIGetSessionsRequest {
	r.limit = &limit
	return r
}

// The offset defines the offsets the start by the number specified
func (r SessionsAPIGetSessionsRequest) Offset(offset int32) SessionsAPIGetSessionsRequest {
	r.offset = &offset
	return r
}

func (r SessionsAPIGetSessionsRequest) Execute() (*GetSessionsResult, *http.Response, error) {
	return r.ApiService.GetSessionsExecute(r)
}

/*
GetSessions List sessions

**Required ACL:** `auth.sessions.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SessionsAPIGetSessionsRequest
*/
func (a *SessionsAPIService) GetSessions(ctx context.Context) SessionsAPIGetSessionsRequest {
	return SessionsAPIGetSessionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSessionsResult
func (a *SessionsAPIService) GetSessionsExecute(r SessionsAPIGetSessionsRequest) (*GetSessionsResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSessionsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsAPIService.GetSessions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recurse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recurse", r.recurse, "")
	} else {
		var defaultValue bool = false
		r.recurse = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accentTenant != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accent-Tenant", r.accentTenant, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SessionsAPIGetUserSessionsRequest struct {
	ctx          context.Context
	ApiService   SessionsAPI
	userUuid     string
	accentTenant *string
	limit        *int32
	offset       *int32
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r SessionsAPIGetUserSessionsRequest) AccentTenant(accentTenant string) SessionsAPIGetUserSessionsRequest {
	r.accentTenant = &accentTenant
	return r
}

// The limit defines the number of individual objects that are returned
func (r SessionsAPIGetUserSessionsRequest) Limit(limit int32) SessionsAPIGetUserSessionsRequest {
	r.limit = &limit
	return r
}

// The offset defines the offsets the start by the number specified
func (r SessionsAPIGetUserSessionsRequest) Offset(offset int32) SessionsAPIGetUserSessionsRequest {
	r.offset = &offset
	return r
}

func (r SessionsAPIGetUserSessionsRequest) Execute() (*GetSessionsResult, *http.Response, error) {
	return r.ApiService.GetUserSessionsExecute(r)
}

/*
GetUserSessions Retrieves the list of sessions associated to a user

**Required ACL**: `auth.users.{user_uuid}.sessions.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userUuid The UUID of the user
	@return SessionsAPIGetUserSessionsRequest
*/
func (a *SessionsAPIService) GetUserSessions(ctx context.Context, userUuid string) SessionsAPIGetUserSessionsRequest {
	return SessionsAPIGetUserSessionsRequest{
		ApiService: a,
		ctx:        ctx,
		userUuid:   userUuid,
	}
}

// Execute executes the request
//
//	@return GetSessionsResult
func (a *SessionsAPIService) GetUserSessionsExecute(r SessionsAPIGetUserSessionsRequest) (*GetSessionsResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetSessionsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsAPIService.GetUserSessions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_uuid}/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"user_uuid"+"}", url.PathEscape(parameterValueToString(r.userUuid, "userUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accentTenant != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accent-Tenant", r.accentTenant, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SessionsAPIUserDeleteSessionRequest struct {
	ctx         context.Context
	ApiService  SessionsAPI
	userUuid    string
	sessionUuid string
}

func (r SessionsAPIUserDeleteSessionRequest) Execute() (*http.Response, error) {
	return r.ApiService.UserDeleteSessionExecute(r)
}

/*
UserDeleteSession Delete a session

**Required ACL**: `auth.users.{user_uuid}.sessions.{session_uuid}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userUuid The UUID of the user
	@param sessionUuid The UUID of the session
	@return SessionsAPIUserDeleteSessionRequest
*/
func (a *SessionsAPIService) UserDeleteSession(ctx context.Context, userUuid string, sessionUuid string) SessionsAPIUserDeleteSessionRequest {
	return SessionsAPIUserDeleteSessionRequest{
		ApiService:  a,
		ctx:         ctx,
		userUuid:    userUuid,
		sessionUuid: sessionUuid,
	}
}

// Execute executes the request
func (a *SessionsAPIService) UserDeleteSessionExecute(r SessionsAPIUserDeleteSessionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionsAPIService.UserDeleteSession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_uuid}/sessions/{session_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_uuid"+"}", url.PathEscape(parameterValueToString(r.userUuid, "userUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"session_uuid"+"}", url.PathEscape(parameterValueToString(r.sessionUuid, "sessionUuid")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}