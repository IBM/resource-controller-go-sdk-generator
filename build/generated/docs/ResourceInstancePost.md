# ResourceInstancePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the instance. Must be 180 characters or less and cannot include any special characters other than &#x60;(space) - . _ :&#x60;. | 
**Target** | **string** | The deployment location where the instance should be hosted. | 
**ResourceGroup** | **string** | Short or long ID of resource group | 
**ResourcePlanId** | **string** | The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog. | 
**Tags** | **[]string** | Tags that are attached to the instance after provisioning. These tags can be searched and managed through the Tagging API in IBM Cloud. | [optional] 
**Parameters** | **map[string]map[string]interface{}** | Configuration options represented as key-value pairs that are passed through to the target resource brokers. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


