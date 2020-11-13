# \GroupsApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupsForUser**](GroupsApi.md#GetGroupsForUser) | **Get** /users/{key}/groups | Get Groups for a User



## GetGroupsForUser

> []Group GetGroupsForUser(ctx, key, optional)

Get Groups for a User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of user | 
 **optional** | ***GetGroupsForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGroupsForUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**|  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

