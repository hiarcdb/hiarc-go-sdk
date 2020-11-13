# \LegalHoldApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLegalHold**](LegalHoldApi.md#CreateLegalHold) | **Post** /legalholds | Create a Legal Hold
[**GetLegalHold**](LegalHoldApi.md#GetLegalHold) | **Get** /legalholds/{key} | Get a Legal Hold&#39;s Info



## CreateLegalHold

> LegalHold CreateLegalHold(ctx, createLegalHoldRequest)

Create a Legal Hold

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createLegalHoldRequest** | [**CreateLegalHoldRequest**](CreateLegalHoldRequest.md)| Legal Hold information | 

### Return type

[**LegalHold**](LegalHold.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegalHold

> LegalHold GetLegalHold(ctx, key)

Get a Legal Hold's Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Legal Hold to get info | 

### Return type

[**LegalHold**](LegalHold.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

