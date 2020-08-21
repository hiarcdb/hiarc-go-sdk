/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"os"
)

// Linger please
var (
	_ _context.Context
)

// FileApiService FileApi service
type FileApiService service

type apiCreateFileRequest struct {
	ctx _context.Context
	apiService *FileApiService
	xHiarcUserKey *string
	request *EntityBase
	file **os.File
}


func (r apiCreateFileRequest) XHiarcUserKey(xHiarcUserKey string) apiCreateFileRequest {
	r.xHiarcUserKey = &xHiarcUserKey
	return r
}

func (r apiCreateFileRequest) Request(request EntityBase) apiCreateFileRequest {
	r.request = &request
	return r
}

func (r apiCreateFileRequest) File(file *os.File) apiCreateFileRequest {
	r.file = &file
	return r
}

/*
CreateFile Create a File
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiCreateFileRequest
*/
func (a *FileApiService) CreateFile(ctx _context.Context) apiCreateFileRequest {
	return apiCreateFileRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return File
*/
func (r apiCreateFileRequest) Execute() (File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  File
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FileApiService.CreateFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
			
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.xHiarcUserKey != nil {
		localVarHeaderParams["X-Hiarc-User-Key"] = parameterToString(*r.xHiarcUserKey, "")
	}
	if r.request != nil {
		paramJson, err := parameterToJson(*r.request)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("request", paramJson)
	}
	localVarFormFileName = "file"
	var localVarFile *os.File
	if r.file != nil {
		localVarFile = *r.file
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["AdminApiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["X-Hiarc-Api-Key"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiDownloadFileRequest struct {
	ctx _context.Context
	apiService *FileApiService
	key string
	xHiarcUserKey *string
}


func (r apiDownloadFileRequest) XHiarcUserKey(xHiarcUserKey string) apiDownloadFileRequest {
	r.xHiarcUserKey = &xHiarcUserKey
	return r
}

/*
DownloadFile Download a File
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param key Key of file to download
@return apiDownloadFileRequest
*/
func (a *FileApiService) DownloadFile(ctx _context.Context, key string) apiDownloadFileRequest {
	return apiDownloadFileRequest{
		apiService: a,
		ctx: ctx,
		key: key,
	}
}

/*
Execute executes the request
 @return *os.File
*/
func (r apiDownloadFileRequest) Execute() (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FileApiService.DownloadFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/{key}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", _neturl.QueryEscape(parameterToString(r.key, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xHiarcUserKey != nil {
		localVarHeaderParams["X-Hiarc-User-Key"] = parameterToString(*r.xHiarcUserKey, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["AdminApiKeyAuth"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["X-Hiarc-Api-Key"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}