# \ResourceBindingsApi

All URIs are relative to *https://resource-controller.cloud.ibm.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceBinding**](ResourceBindingsApi.md#CreateResourceBinding) | **Post** /resource_bindings | Create a new resource binding
[**DeleteResourceBinding**](ResourceBindingsApi.md#DeleteResourceBinding) | **Delete** /resource_bindings/{id} | Delete a resource binding
[**GetResourceBinding**](ResourceBindingsApi.md#GetResourceBinding) | **Get** /resource_bindings/{id} | Get a resource binding
[**ListResourceBindings**](ResourceBindingsApi.md#ListResourceBindings) | **Get** /resource_bindings | Get a list of all resource bindings
[**UpdateResourceBinding**](ResourceBindingsApi.md#UpdateResourceBinding) | **Patch** /resource_bindings/{id} | Update a resource binding



## CreateResourceBinding

> ResourceBinding CreateResourceBinding(ctx, authorization, resourceBindingPost)

Create a new resource binding

Create a new resource binding.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**resourceBindingPost** | [**ResourceBindingPost**](ResourceBindingPost.md)|  | 

### Return type

[**ResourceBinding**](ResourceBinding.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceBinding

> DeleteResourceBinding(ctx, authorization, id)

Delete a resource binding

Delete a resource binding by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the binding. | 

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


## GetResourceBinding

> ResourceBinding GetResourceBinding(ctx, authorization, id)

Get a resource binding

Retrieve a resource binding by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the binding. | 

### Return type

[**ResourceBinding**](ResourceBinding.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceBindings

> ResourceBindingsList ListResourceBindings(ctx, authorization, optional)

Get a list of all resource bindings

Get a list of all resource bindings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
 **optional** | ***ListResourceBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListResourceBindingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guid** | **optional.String**| The short ID of the binding. | 
 **name** | **optional.String**| The human-readable name of the binding. | 
 **resourceGroupId** | **optional.String**| Short ID of the resource group. | 
 **resourceId** | **optional.String**| The unique ID of the offering (service name). This value is provided by and stored in the global catalog. | 
 **regionBindingId** | **optional.String**| Short ID of the binding in the specific targeted environment, e.g. service_binding_id in a given IBM Cloud environment. | 
 **limit** | **optional.String**| Limit on how many items should be returned | 

### Return type

[**ResourceBindingsList**](ResourceBindingsList.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceBinding

> ResourceBinding UpdateResourceBinding(ctx, authorization, id, resourceBindingPatch)

Update a resource binding

Update a resource binding by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the binding. | 
**resourceBindingPatch** | [**ResourceBindingPatch**](ResourceBindingPatch.md)|  | 

### Return type

[**ResourceBinding**](ResourceBinding.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

