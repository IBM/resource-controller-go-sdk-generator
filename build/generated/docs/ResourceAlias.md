# ResourceAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID associated with the alias. | [optional] 
**Guid** | **string** | When you create a new alias, a globally unique identifier (GUID) is assigned. This GUID is a unique internal indentifier managed by the resource controller that corresponds to the alias. | [optional] 
**Crn** | **string** | The full Cloud Resource Name (CRN) associated with the alias. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic&#x3D;overview-crn). | [optional] 
**Url** | **string** | When you created a new alias, a relative URL path is created identifying the location of the alias. | [optional] 
**Name** | **string** | The human-readable name of the alias. | [optional] 
**AccountId** | **string** | An alpha-numeric value identifying the account ID. | [optional] 
**ResourceGroupId** | **string** | The short ID of the resource group. | [optional] 
**ResourceGroupCrn** | **string** | The long ID (full CRN) of the resource group. | [optional] 
**TargetCrn** | **string** | The CRN of the target namespace in the specific environment. | [optional] 
**State** | **string** | The state of the alias. | [optional] 
**ResourceInstanceId** | **string** | The short ID of the resource instance that is being aliased. | [optional] 
**RegionInstanceId** | **string** | The short ID of the instance in the specific target environment, e.g. &#x60;service_instance_id&#x60; in a given IBM Cloud environment. | [optional] 
**ResourceInstanceUrl** | **string** | The relative path to the instance. | [optional] 
**ResourceBindingsUrl** | **string** | The relative path to the resource bindings for the alias. | [optional] 
**ResourceKeysUrl** | **string** | The relative path to the resource keys for the alias. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date when the alias was created. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date when the alias was last updated. | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) | The date when the alias was deleted. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


