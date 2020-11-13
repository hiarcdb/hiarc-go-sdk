# \ClassificationApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClassification**](ClassificationApi.md#CreateClassification) | **Post** /classifications | Create a Classification
[**DeleteClassification**](ClassificationApi.md#DeleteClassification) | **Delete** /classifications/{key} | Delete a Classification
[**FindClassification**](ClassificationApi.md#FindClassification) | **Post** /classifications/find | Find a Classification
[**GetAllClassifications**](ClassificationApi.md#GetAllClassifications) | **Get** /classifications | Get all Classifications
[**GetClassification**](ClassificationApi.md#GetClassification) | **Get** /classifications/{key} | Get a Classification&#39;s Info
[**UpdateClassification**](ClassificationApi.md#UpdateClassification) | **Put** /classifications/{key} | Update a Classification



## CreateClassification

> Classification CreateClassification(ctx, createClassificationRequest, optional)

Create a Classification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createClassificationRequest** | [**CreateClassificationRequest**](CreateClassificationRequest.md)| Classification information | 
 **optional** | ***CreateClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateClassificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**Classification**](Classification.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClassification

> map[string]interface{} DeleteClassification(ctx, key, optional)

Delete a Classification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Classification to delete | 
 **optional** | ***DeleteClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteClassificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**map[string]interface{}**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindClassification

> []Classification FindClassification(ctx, findClassificationsRequest, optional)

Find a Classification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**findClassificationsRequest** | [**FindClassificationsRequest**](FindClassificationsRequest.md)| Classification query | 
 **optional** | ***FindClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindClassificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]Classification**](Classification.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllClassifications

> []Classification GetAllClassifications(ctx, optional)

Get all Classifications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllClassificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllClassificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]Classification**](Classification.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClassification

> Classification GetClassification(ctx, key, optional)

Get a Classification's Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Classification to get info | 
 **optional** | ***GetClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClassificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**Classification**](Classification.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClassification

> Classification UpdateClassification(ctx, key, updateClassificationRequest, optional)

Update a Classification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Classification to get info | 
**updateClassificationRequest** | [**UpdateClassificationRequest**](UpdateClassificationRequest.md)| Classification information | 
 **optional** | ***UpdateClassificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateClassificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**Classification**](Classification.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

