# ResourceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID associated with the instance. | [optional] 
**Guid** | **string** | When you create a new resource, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the instance. | [optional] 
**Crn** | **string** | The full Cloud Resource Name (CRN) associated with the instance. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic&#x3D;overview-crn). | [optional] 
**Url** | **string** | When you provision a new resource, a relative URL path is created identifying the location of the instance. | [optional] 
**Name** | **string** | The human-readable name of the instance. | [optional] 
**AccountId** | **string** | An alpha-numeric value identifying the account ID. | [optional] 
**ResourceGroupId** | **string** | The short ID of the resource group. | [optional] 
**ResourceGroupCrn** | **string** | The long ID (full CRN) of the resource group. | [optional] 
**ResourceId** | **string** | The unique ID of the offering. This value is provided by and stored in the global catalog. | [optional] 
**ResourcePlanId** | **string** | The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog. | [optional] 
**TargetCrn** | **string** | The full deployment CRN as defined in the global catalog. The Cloud Resource Name (CRN) of the deployment location where the instance is provisioned. | [optional] 
**State** | **string** | The current state of the instance. For example, if the instance is deleted, it will return removed. | [optional] 
**Type** | **string** | The type of the instance, e.g. &#x60;service_instance&#x60;. | [optional] 
**SubType** | **string** | The sub-type of instance, e.g. &#x60;cfaas&#x60;. | [optional] 
**LastOperation** | **map[string]interface{}** | The status of the last operation requested on the instance. | [optional] 
**DashboardUrl** | **string** | The resource-broker-provided URL to access administrative features of the instance. | [optional] 
**PlanHistory** | [**[]PlanHistoryItem**](PlanHistoryItem.md) | The plan history of the instance. | [optional] 
**ResourceAliasesUrl** | **string** | The relative path to the resource aliases for the instance. | [optional] 
**ResourceBindingsUrl** | **string** | The relative path to the resource bindings for the instance. | [optional] 
**ResourceKeysUrl** | **string** | The relative path to the resource keys for the instance. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date when the instance was created. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date when the instance was last updated. | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) | The date when the instance was deleted. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


