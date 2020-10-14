# \ResourceAliasesApi

All URIs are relative to *https://resource-controller.cloud.ibm.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceAlias**](ResourceAliasesApi.md#CreateResourceAlias) | **Post** /resource_aliases | Create a new resource alias
[**DeleteResourceAlias**](ResourceAliasesApi.md#DeleteResourceAlias) | **Delete** /resource_aliases/{id} | Delete a resource alias
[**GetResourceAlias**](ResourceAliasesApi.md#GetResourceAlias) | **Get** /resource_aliases/{id} | Get a resource alias
[**ListResourceAliases**](ResourceAliasesApi.md#ListResourceAliases) | **Get** /resource_aliases | Get a list of all resource aliases
[**UpdateResourceAlias**](ResourceAliasesApi.md#UpdateResourceAlias) | **Patch** /resource_aliases/{id} | Update a resource alias



## CreateResourceAlias

> ResourceAlias CreateResourceAlias(ctx, authorization, resourceAliasPost)

Create a new resource alias

Alias a resource instance into a targeted environment's (name)space.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**resourceAliasPost** | [**ResourceAliasPost**](ResourceAliasPost.md)|  | 

### Return type

[**ResourceAlias**](ResourceAlias.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceAlias

> DeleteResourceAlias(ctx, authorization, id)

Delete a resource alias

Delete a resource alias by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the alias. | 

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


## GetResourceAlias

> ResourceAlias GetResourceAlias(ctx, authorization, id)

Get a resource alias

Retrieve a resource alias by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the alias. | 

### Return type

[**ResourceAlias**](ResourceAlias.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceAliases

> ResourceAliasesList ListResourceAliases(ctx, authorization, optional)

Get a list of all resource aliases

Get a list of all resource aliases.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
 **optional** | ***ListResourceAliasesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListResourceAliasesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guid** | **optional.String**| Short ID of the alias. | 
 **name** | **optional.String**| The human-readable name of the alias. | 
 **resourceInstanceId** | **optional.String**| Resource instance short ID. | 
 **regionInstanceId** | **optional.String**| Short ID of the instance in a specific targeted environment. For example, &#x60;service_instance_id&#x60; in a given IBM Cloud environment. | 
 **resourceId** | **optional.String**| The unique ID of the offering (service name). This value is provided by and stored in the global catalog. | 
 **resourceGroupId** | **optional.String**| Short ID of Resource group. | 
 **limit** | **optional.String**| Limit on how many items should be returned. | 

### Return type

[**ResourceAliasesList**](ResourceAliasesList.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceAlias

> ResourceAlias UpdateResourceAlias(ctx, authorization, id, resourceAliasPatch)

Update a resource alias

Update a resource alias by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the alias. | 
**resourceAliasPatch** | [**ResourceAliasPatch**](ResourceAliasPatch.md)|  | 

### Return type

[**ResourceAlias**](ResourceAlias.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

