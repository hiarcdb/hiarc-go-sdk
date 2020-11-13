# \RetentionPolicyApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRetentionPolicy**](RetentionPolicyApi.md#CreateRetentionPolicy) | **Post** /retentionpolicies | Create a Retention Policy
[**FindRetentionPolicies**](RetentionPolicyApi.md#FindRetentionPolicies) | **Post** /retentionpolicies/find | Find a Retention Policy
[**GetAllRetentionPolicies**](RetentionPolicyApi.md#GetAllRetentionPolicies) | **Get** /retentionpolicies | Get all Retention Policies
[**GetRetentionPolicy**](RetentionPolicyApi.md#GetRetentionPolicy) | **Get** /retentionpolicies/{key} | Get a Retention Policy&#39;s Info
[**UpdateRetentionPolicy**](RetentionPolicyApi.md#UpdateRetentionPolicy) | **Put** /retentionpolicies/{key} | Update a Retention Policy



## CreateRetentionPolicy

> RetentionPolicy CreateRetentionPolicy(ctx, createRetentionPolicyRequest)

Create a Retention Policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createRetentionPolicyRequest** | [**CreateRetentionPolicyRequest**](CreateRetentionPolicyRequest.md)| Retention Policy information | 

### Return type

[**RetentionPolicy**](RetentionPolicy.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindRetentionPolicies

> []RetentionPolicy FindRetentionPolicies(ctx, findRetentionPoliciesRequest)

Find a Retention Policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**findRetentionPoliciesRequest** | [**FindRetentionPoliciesRequest**](FindRetentionPoliciesRequest.md)| Retention Policy query | 

### Return type

[**[]RetentionPolicy**](RetentionPolicy.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRetentionPolicies

> []RetentionPolicy GetAllRetentionPolicies(ctx, )

Get all Retention Policies

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RetentionPolicy**](RetentionPolicy.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetentionPolicy

> RetentionPolicy GetRetentionPolicy(ctx, key)

Get a Retention Policy's Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Retention Policy to get info | 

### Return type

[**RetentionPolicy**](RetentionPolicy.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRetentionPolicy

> RetentionPolicy UpdateRetentionPolicy(ctx, key, updateRetentionPolicyRequest)

Update a Retention Policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Retention Policy to update | 
**updateRetentionPolicyRequest** | [**UpdateRetentionPolicyRequest**](UpdateRetentionPolicyRequest.md)| RetentionPolicy information | 

### Return type

[**RetentionPolicy**](RetentionPolicy.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

