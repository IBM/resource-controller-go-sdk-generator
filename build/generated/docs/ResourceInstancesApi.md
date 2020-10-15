# \ResourceInstancesApi

All URIs are relative to *https://resource-controller.cloud.ibm.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceInstance**](ResourceInstancesApi.md#CreateResourceInstance) | **Post** /resource_instances | Create (provision) a new resource instance
[**DeleteResourceInstance**](ResourceInstancesApi.md#DeleteResourceInstance) | **Delete** /resource_instances/{id} | Delete a resource instance
[**GetResourceInstance**](ResourceInstancesApi.md#GetResourceInstance) | **Get** /resource_instances/{id} | Get a resource instance
[**ListResourceInstances**](ResourceInstancesApi.md#ListResourceInstances) | **Get** /resource_instances | Get a list of all resource instances
[**UpdateResourceInstance**](ResourceInstancesApi.md#UpdateResourceInstance) | **Patch** /resource_instances/{id} | Update a resource instance



## CreateResourceInstance

> ResourceInstance CreateResourceInstance(ctx, authorization, resourceInstancePost)

Create (provision) a new resource instance

Provision a new resource in the specified location for the selected plan.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**resourceInstancePost** | [**ResourceInstancePost**](ResourceInstancePost.md)|  | 

### Return type

[**ResourceInstance**](ResourceInstance.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceInstance

> ResourceInstance DeleteResourceInstance(ctx, authorization, id)

Delete a resource instance

Delete a resource instance by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the instance. | 

### Return type

[**ResourceInstance**](ResourceInstance.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceInstance

> ResourceInstance GetResourceInstance(ctx, authorization, id)

Get a resource instance

Retrieve a resource instance by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the instance. | 

### Return type

[**ResourceInstance**](ResourceInstance.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceInstances

> ResourceInstancesList ListResourceInstances(ctx, authorization, optional)

Get a list of all resource instances

Get a list of all resource instances.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
 **optional** | ***ListResourceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListResourceInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guid** | **optional.String**| When you provision a new resource in the specified location for the selected plan, a GUID (globally unique identifier) is created. This is a unique internal GUID managed by Resource controller that corresponds to the instance. | 
 **name** | **optional.String**| The human-readable name of the instance. | 
 **resourceGroupId** | **optional.String**| Short ID of a resource group. | 
 **resourceId** | **optional.String**| The unique ID of the offering. This value is provided by and stored in the global catalog. | 
 **resourcePlanId** | **optional.String**| The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog. | 
 **type_** | **optional.String**| The type of the instance. For example, &#x60;service_instance&#x60;. | 
 **subType** | **optional.String**| The sub-type of instance, e.g. &#x60;cfaas&#x60;. | 
 **limit** | **optional.String**| Limit on how many items should be returned. | 

### Return type

[**ResourceInstancesList**](ResourceInstancesList.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceInstance

> ResourceInstance UpdateResourceInstance(ctx, authorization, id, resourceInstancePatch)

Update a resource instance

Update a resource instance by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Authorization value should use the format Bearer [token] - where [token] is an opaque access token obtained by authenticating with IAM authentication services. | 
**id** | **string**| The short or long ID of the instance. | 
**resourceInstancePatch** | [**ResourceInstancePatch**](ResourceInstancePatch.md)|  | 

### Return type

[**ResourceInstance**](ResourceInstance.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

