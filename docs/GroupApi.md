# \GroupApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToGroup**](GroupApi.md#AddUserToGroup) | **Put** /groups/{key}/users/{userKey} | Add a User to a Group
[**CreateGroup**](GroupApi.md#CreateGroup) | **Post** /groups | Create a Group
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /groups/{key} | Delete a Group
[**FindGroup**](GroupApi.md#FindGroup) | **Post** /groups/find | Find a Group
[**GetAllGroups**](GroupApi.md#GetAllGroups) | **Get** /groups | Get all Groups
[**GetGroup**](GroupApi.md#GetGroup) | **Get** /groups/{key} | Get a Group&#39;s Info
[**GetGroupsForCurrentUser**](GroupApi.md#GetGroupsForCurrentUser) | **Get** /users/current/groups | Get the Groups for the current User
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Put** /groups/{key} | Update a Group



## AddUserToGroup

> map[string]interface{} AddUserToGroup(ctx, key, userKey)

Add a User to a Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Group | 
**userKey** | **string**| Key of User to add to Group | 

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


## CreateGroup

> Group CreateGroup(ctx, createGroupRequest)

Create a Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createGroupRequest** | [**CreateGroupRequest**](CreateGroupRequest.md)| Group information | 

### Return type

[**Group**](Group.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> map[string]interface{} DeleteGroup(ctx, key)

Delete a Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Group to delete | 

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


## FindGroup

> []Group FindGroup(ctx, findGroupsRequest)

Find a Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**findGroupsRequest** | [**FindGroupsRequest**](FindGroupsRequest.md)| Group query | 

### Return type

[**[]Group**](Group.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllGroups

> []Group GetAllGroups(ctx, )

Get all Groups

### Required Parameters

This endpoint does not need any parameter.

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


## GetGroup

> Group GetGroup(ctx, key)

Get a Group's Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Group to get info | 

### Return type

[**Group**](Group.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsForCurrentUser

> []Group GetGroupsForCurrentUser(ctx, optional)

Get the Groups for the current User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGroupsForCurrentUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGroupsForCurrentUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHiarcUserKey** | **optional.String**|  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, key, updateGroupRequest)

Update a Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Group to update | 
**updateGroupRequest** | [**UpdateGroupRequest**](UpdateGroupRequest.md)| Group information | 

### Return type

[**Group**](Group.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

