# ResourceBindingPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the binding. Must be 180 characters or less and cannot include any special characters other than &#x60;(space) - . _ :&#x60;. | [optional] 
**Source** | **string** | The short or long ID of resource alias. | 
**Target** | **string** | The CRN of application to bind to in a specific environment, e.g. Dallas YP, CFEE instance | 
**Parameters** | **map[string]map[string]interface{}** | Configuration options represented as key-value pairs that are passed through to the target resource brokers. | [optional] 
**Role** | **string** | The role name or it&#39;s CRN. | [optional] [default to Writer]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


