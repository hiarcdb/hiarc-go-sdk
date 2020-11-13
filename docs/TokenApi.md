# \TokenApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserToken**](TokenApi.md#CreateUserToken) | **Post** /tokens/user | Create a Token for a User



## CreateUserToken

> UserCredentials CreateUserToken(ctx, createUserTokenRequest)

Create a Token for a User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createUserTokenRequest** | [**CreateUserTokenRequest**](CreateUserTokenRequest.md)| User information | 

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

