# \CollectionApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChildToCollection**](CollectionApi.md#AddChildToCollection) | **Put** /collections/{key}/children/{childKey} | Add a child item to a Collection
[**AddFileToCollection**](CollectionApi.md#AddFileToCollection) | **Put** /collections/{key}/files | Add a File to a Collection
[**AddGroupToCollection**](CollectionApi.md#AddGroupToCollection) | **Put** /collections/{key}/groups | Add a Group to a Collection
[**AddUserToCollection**](CollectionApi.md#AddUserToCollection) | **Put** /collections/{key}/users | Add a User to a Collection
[**CreateCollection**](CollectionApi.md#CreateCollection) | **Post** /collections | Create a Collection
[**DeleteCollection**](CollectionApi.md#DeleteCollection) | **Delete** /collections/{key} | Delete a Collection
[**FindCollection**](CollectionApi.md#FindCollection) | **Post** /collections/find | Find a Collection
[**GetAllCollections**](CollectionApi.md#GetAllCollections) | **Get** /collections | Get all Collections
[**GetCollection**](CollectionApi.md#GetCollection) | **Get** /collections/{key} | Get a Collection&#39;s Info
[**GetCollectionChildren**](CollectionApi.md#GetCollectionChildren) | **Get** /collections/{key}/children | Get a Collection&#39;s child Collections
[**GetCollectionFiles**](CollectionApi.md#GetCollectionFiles) | **Get** /collections/{key}/files | Get a Collection&#39;s Files
[**GetCollectionItems**](CollectionApi.md#GetCollectionItems) | **Get** /collections/{key}/items | Get a Collection&#39;s child items, including Collections and Files
[**RemoveFileFromCollection**](CollectionApi.md#RemoveFileFromCollection) | **Delete** /collections/{key}/files/{fileKey} | Remove a File from a Collection
[**UpdateCollection**](CollectionApi.md#UpdateCollection) | **Put** /collections/{key} | Update a Collection



## AddChildToCollection

> map[string]interface{} AddChildToCollection(ctx, key, childKey, optional)

Add a child item to a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Collection | 
**childKey** | **string**| Key of item to add as child to Collection | 
 **optional** | ***AddChildToCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddChildToCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**map[string]interface{}**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFileToCollection

> map[string]interface{} AddFileToCollection(ctx, key, addFileToCollectionRequest, optional)

Add a File to a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Collection | 
**addFileToCollectionRequest** | [**AddFileToCollectionRequest**](AddFileToCollectionRequest.md)| Add File request | 
 **optional** | ***AddFileToCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddFileToCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**map[string]interface{}**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupToCollection

> map[string]interface{} AddGroupToCollection(ctx, key, addGroupToCollectionRequest, optional)

Add a Group to a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Collection | 
**addGroupToCollectionRequest** | [**AddGroupToCollectionRequest**](AddGroupToCollectionRequest.md)| Add Group request | 
 **optional** | ***AddGroupToCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddGroupToCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**map[string]interface{}**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToCollection

> map[string]interface{} AddUserToCollection(ctx, key, addUserToCollectionRequest, optional)

Add a User to a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Collection | 
**addUserToCollectionRequest** | [**AddUserToCollectionRequest**](AddUserToCollectionRequest.md)| Add User request | 
 **optional** | ***AddUserToCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddUserToCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**map[string]interface{}**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollection

> Collection CreateCollection(ctx, createCollectionRequest, optional)

Create a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createCollectionRequest** | [**CreateCollectionRequest**](CreateCollectionRequest.md)| Collection information | 
 **optional** | ***CreateCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**Collection**](Collection.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> map[string]interface{} DeleteCollection(ctx, key, optional)

Delete a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Collection to delete | 
 **optional** | ***DeleteCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**map[string]interface{}**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCollection

> []Collection FindCollection(ctx, findCollectionsRequest, optional)

Find a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**findCollectionsRequest** | [**FindCollectionsRequest**](FindCollectionsRequest.md)| Collection query | 
 **optional** | ***FindCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCollections

> []Collection GetAllCollections(ctx, optional)

Get all Collections

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllCollectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollection

> Collection GetCollection(ctx, key, optional)

Get a Collection's Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of collection to get info | 
 **optional** | ***GetCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**Collection**](Collection.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionChildren

> []Collection GetCollectionChildren(ctx, key, optional)

Get a Collection's child Collections

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of collection | 
 **optional** | ***GetCollectionChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCollectionChildrenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionFiles

> []File GetCollectionFiles(ctx, key, optional)

Get a Collection's Files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of collection | 
 **optional** | ***GetCollectionFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCollectionFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]File**](File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItems

> CollectionItems GetCollectionItems(ctx, key, optional)

Get a Collection's child items, including Collections and Files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of collection | 
 **optional** | ***GetCollectionItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCollectionItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**CollectionItems**](CollectionItems.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFileFromCollection

> map[string]interface{} RemoveFileFromCollection(ctx, key, fileKey, optional)

Remove a File from a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of Collection | 
**fileKey** | **string**| Key of File to remove from Collection | 
 **optional** | ***RemoveFileFromCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveFileFromCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

**map[string]interface{}**

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> Collection UpdateCollection(ctx, key, updateCollectionRequest, optional)

Update a Collection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of collection to get info | 
**updateCollectionRequest** | [**UpdateCollectionRequest**](UpdateCollectionRequest.md)| Collection information | 
 **optional** | ***UpdateCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**Collection**](Collection.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

