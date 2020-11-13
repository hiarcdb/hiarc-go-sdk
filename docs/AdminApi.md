# \AdminApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InitDB**](AdminApi.md#InitDB) | **Post** /admin/database/init | Init the Hiarc database
[**ResetDB**](AdminApi.md#ResetDB) | **Put** /admin/database/reset | Reset the Hiarc database



## InitDB

> map[string]interface{} InitDB(ctx, )

Init the Hiarc database

### Required Parameters

This endpoint does not need any parameter.

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


## ResetDB

> map[string]interface{} ResetDB(ctx, )

Reset the Hiarc database

### Required Parameters

This endpoint does not need any parameter.

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

