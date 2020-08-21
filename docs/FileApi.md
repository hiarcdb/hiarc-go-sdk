# \FileApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFile**](FileApi.md#CreateFile) | **Post** /files | Create a File
[**DownloadFile**](FileApi.md#DownloadFile) | **Get** /files/{key}/download | Download a File



## CreateFile

> File CreateFile(ctx).XHiarcUserKey(xHiarcUserKey).Request(request).File(file).Execute()

Create a File

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xHiarcUserKey := "xHiarcUserKey_example" // string |  (optional)
    request := openapiclient.EntityBase{Key: "Key_example", Name: "Name_example", Description: "Description_example", Metadata: "TODO"} // EntityBase |  (optional)
    file := 987 // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileApi.CreateFile(context.Background(), ).XHiarcUserKey(xHiarcUserKey).Request(request).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.CreateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFile`: File
    fmt.Fprintf(os.Stdout, "Response from `FileApi.CreateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHiarcUserKey** | **string** |  | 
 **request** | [**EntityBase**](EntityBase.md) |  | 
 **file** | ***os.File** |  | 

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


## DownloadFile

> *os.File DownloadFile(ctx, key).XHiarcUserKey(xHiarcUserKey).Execute()

Download a File

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | Key of file to download
    xHiarcUserKey := "xHiarcUserKey_example" // string | Optional key of user to impersonate (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FileApi.DownloadFile(context.Background(), key).XHiarcUserKey(xHiarcUserKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileApi.DownloadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FileApi.DownloadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key of file to download | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHiarcUserKey** | **string** | Optional key of user to impersonate | 

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

