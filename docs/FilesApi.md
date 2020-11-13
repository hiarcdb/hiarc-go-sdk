# \FilesApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterAllowedFiles**](FilesApi.md#FilterAllowedFiles) | **Post** /files/allowed | Filter a list of File keys that a User can access



## FilterAllowedFiles

> []string FilterAllowedFiles(ctx, allowedFilesRequest, optional)

Filter a list of File keys that a User can access

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**allowedFilesRequest** | [**AllowedFilesRequest**](AllowedFilesRequest.md)| File key list | 
 **optional** | ***FilterAllowedFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FilterAllowedFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**[]string**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

