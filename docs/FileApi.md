# \FileApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClassificationToFile**](FileApi.md#AddClassificationToFile) | **Put** /files/{key}/classifications | Add a Classification to a File
[**AddGroupToFile**](FileApi.md#AddGroupToFile) | **Put** /files/{key}/groups | Give a group access to a File
[**AddRetentionPolicyToFile**](FileApi.md#AddRetentionPolicyToFile) | **Put** /files/{key}/retentionpolicies | Add a Retention Policy to a File
[**AddUserToFile**](FileApi.md#AddUserToFile) | **Put** /files/{key}/users | Give a user access to a File
[**AddVersion**](FileApi.md#AddVersion) | **Put** /files/{key}/versions | Add a new File Version
[**AttachToExisitingFile**](FileApi.md#AttachToExisitingFile) | **Put** /files/{key}/attach | Attach to an existing File
[**CopyFile**](FileApi.md#CopyFile) | **Put** /files/{key}/copy | Copy a File
[**CreateDirectUploadUrl**](FileApi.md#CreateDirectUploadUrl) | **Post** /files/directuploadurl | Create a direct upload url to a storage service
[**CreateFile**](FileApi.md#CreateFile) | **Post** /files | Create a File
[**DeleteFile**](FileApi.md#DeleteFile) | **Delete** /files/{key} | Delete a File
[**DownloadFile**](FileApi.md#DownloadFile) | **Get** /files/{key}/download | Download a File
[**GetCollectionsForFile**](FileApi.md#GetCollectionsForFile) | **Get** /files/{key}/collections | Get a list of Collections for a File
[**GetDirectDownloadUrl**](FileApi.md#GetDirectDownloadUrl) | **Get** /files/{key}/directdownloadurl | Get a direct download URL to a File
[**GetFile**](FileApi.md#GetFile) | **Get** /files/{key} | Get a File&#39;s Info
[**GetRetentionPolicies**](FileApi.md#GetRetentionPolicies) | **Get** /files/{key}/retentionpolicies | Get a list of Retention Policies on a File
[**GetVersions**](FileApi.md#GetVersions) | **Get** /files/{key}/versions | Get a list of File Versions
[**UpdateFile**](FileApi.md#UpdateFile) | **Put** /files/{key} | Update a File



## AddClassificationToFile

> map[string]interface{} AddClassificationToFile(ctx, key, addClassificationToFileRequest, optional)

Add a Classification to a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file | 
**addClassificationToFileRequest** | [**AddClassificationToFileRequest**](AddClassificationToFileRequest.md)| Classification information | 
 **optional** | ***AddClassificationToFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddClassificationToFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**|  | 

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


## AddGroupToFile

> map[string]interface{} AddGroupToFile(ctx, key, addGroupToFileRequest, optional)

Give a group access to a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file | 
**addGroupToFileRequest** | [**AddGroupToFileRequest**](AddGroupToFileRequest.md)| Group information | 
 **optional** | ***AddGroupToFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddGroupToFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**|  | 

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


## AddRetentionPolicyToFile

> map[string]interface{} AddRetentionPolicyToFile(ctx, key, addRetentionPolicyToFileRequest, optional)

Add a Retention Policy to a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file | 
**addRetentionPolicyToFileRequest** | [**AddRetentionPolicyToFileRequest**](AddRetentionPolicyToFileRequest.md)| Retention Policy information | 
 **optional** | ***AddRetentionPolicyToFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddRetentionPolicyToFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**|  | 

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


## AddUserToFile

> map[string]interface{} AddUserToFile(ctx, key, addUserToFileRequest, optional)

Give a user access to a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file | 
**addUserToFileRequest** | [**AddUserToFileRequest**](AddUserToFileRequest.md)| User information | 
 **optional** | ***AddUserToFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddUserToFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**|  | 

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


## AddVersion

> File AddVersion(ctx, key, optional)

Add a new File Version

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to add a version | 
 **optional** | ***AddVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**|  | 
 **request** | **optional.String**|  | 
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**File**](File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachToExisitingFile

> File AttachToExisitingFile(ctx, key, attachToExistingFileRequest, optional)

Attach to an existing File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to attach to | 
**attachToExistingFileRequest** | [**AttachToExistingFileRequest**](AttachToExistingFileRequest.md)| File information | 
 **optional** | ***AttachToExisitingFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AttachToExisitingFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**|  | 

### Return type

[**File**](File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyFile

> File CopyFile(ctx, key, copyFileRequest, optional)

Copy a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to attach to | 
**copyFileRequest** | [**CopyFileRequest**](CopyFileRequest.md)| File information | 
 **optional** | ***CopyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CopyFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHiarcUserKey** | **optional.String**|  | 

### Return type

[**File**](File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDirectUploadUrl

> FileDirectUpload CreateDirectUploadUrl(ctx, createDirectUploadUrlRequest, optional)

Create a direct upload url to a storage service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createDirectUploadUrlRequest** | [**CreateDirectUploadUrlRequest**](CreateDirectUploadUrlRequest.md)| Storage service information | 
 **optional** | ***CreateDirectUploadUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDirectUploadUrlOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 
 **expiresInSeconds** | **optional.Int32**| When access to the url should expire | 

### Return type

[**FileDirectUpload**](FileDirectUpload.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFile

> File CreateFile(ctx, optional)

Create a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHiarcUserKey** | **optional.String**|  | 
 **request** | **optional.String**|  | 
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**File**](File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFile

> map[string]interface{} DeleteFile(ctx, key, optional)

Delete a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to delete | 
 **optional** | ***DeleteFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteFileOpts struct


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


## DownloadFile

> *os.File DownloadFile(ctx, key, optional)

Download a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to download | 
 **optional** | ***DownloadFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionsForFile

> []Collection GetCollectionsForFile(ctx, key, optional)

Get a list of Collections for a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to get all collections | 
 **optional** | ***GetCollectionsForFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCollectionsForFileOpts struct


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


## GetDirectDownloadUrl

> FileDirectDownload GetDirectDownloadUrl(ctx, key, optional)

Get a direct download URL to a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to get download URL | 
 **optional** | ***GetDirectDownloadUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDirectDownloadUrlOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 
 **expiresInSeconds** | **optional.Int32**| When access to the url should expire | 

### Return type

[**FileDirectDownload**](FileDirectDownload.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> File GetFile(ctx, key, optional)

Get a File's Info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to get info | 
 **optional** | ***GetFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**File**](File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRetentionPolicies

> []RetentionPolicyApplication GetRetentionPolicies(ctx, key, optional)

Get a list of Retention Policies on a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to get all retention policies | 
 **optional** | ***GetRetentionPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRetentionPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]RetentionPolicyApplication**](RetentionPolicyApplication.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersions

> []FileVersion GetVersions(ctx, key, optional)

Get a list of File Versions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to get all versions | 
 **optional** | ***GetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **optional.String**| Optional key of user to impersonate | 

### Return type

[**[]FileVersion**](FileVersion.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth), [JWTBearerAuth](../README.md#JWTBearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFile

> File UpdateFile(ctx, key, updateFileRequest)

Update a File

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Key of file to update | 
**updateFileRequest** | [**UpdateFileRequest**](UpdateFileRequest.md)| File information | 

### Return type

[**File**](File.md)

### Authorization

[AdminApiKeyAuth](../README.md#AdminApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

