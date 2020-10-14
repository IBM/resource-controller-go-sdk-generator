# \ResourceKeysApi

All URIs are relative to *https://resource-controller.cloud.ibm.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceKey**](ResourceKeysApi.md#CreateResourceKey) | **Post** /resource_keys | Create a new resource key
[**DeleteResourceKey**](ResourceKeysApi.md#DeleteResourceKey) | **Delete** /resource_keys/{id} | Delete a resource key by ID
[**GetResourceKey**](ResourceKeysApi.md#GetResourceKey) | **Get** /resource_keys/{id} | Get resource key by ID
[**ListResourceKeys**](ResourceKeysApi.md#ListResourceKeys) | **Get** /resource_keys | Get a list of resource keys
[**UpdateResourceKey**](ResourceKeysApi.md#UpdateResourceKey) | **Patch** /resource_keys/{id} | Update a resource key



## CreateResourceKey

> ResourceKey CreateResourceKey(ctx, authorization, resourceKeyPost)

Create a new resource key

Create a new resource key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**resourceKeyPost** | [**ResourceKeyPost**](ResourceKeyPost.md)|  | 

### Return type

[**ResourceKey**](ResourceKey.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceKey

> DeleteResourceKey(ctx, authorization, id)

Delete a resource key by ID

Delete a resource key by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the key. | 

### Return type

 (empty response body)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceKey

> ResourceKey GetResourceKey(ctx, authorization, id)

Get resource key by ID

Get resource key by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the key. | 

### Return type

[**ResourceKey**](ResourceKey.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceKeys

> ResourceKeysList ListResourceKeys(ctx, authorization, optional)

Get a list of resource keys

List all resource keys

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
 **optional** | ***ListResourceKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListResourceKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guid** | **optional.String**| When you create a new key, a GUID (globally unique identifier) is assigned. This is a unique internal GUID managed by Resource controller that corresponds to the key. | 
 **name** | **optional.String**| The human-readable name of the key. | 
 **resourceGroupId** | **optional.String**| The short ID of the resource group. | 
 **resourceId** | **optional.String**| The unique ID of the offering. This value is provided by and stored in the global catalog. | 
 **limit** | **optional.String**| Limit on how many items should be returned. | 

### Return type

[**ResourceKeysList**](ResourceKeysList.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceKey

> ResourceKey UpdateResourceKey(ctx, authorization, id, resourceKeyPatch)

Update a resource key

Update a resource key by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the key. | 
**resourceKeyPatch** | [**ResourceKeyPatch**](ResourceKeyPatch.md)|  | 

### Return type

[**ResourceKey**](ResourceKey.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

