# ResourceKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID associated with the key. | [optional] 
**Guid** | **string** | When you create a new key, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the key. | [optional] 
**Crn** | **string** | The full Cloud Resource Name (CRN) associated with the key. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic&#x3D;overview-crn). | [optional] 
**Url** | **string** | When you created a new key, a relative URL path is created identifying the location of the key. | [optional] 
**Name** | **string** | The human-readable name of the key. | [optional] 
**AccountId** | **string** | An alpha-numeric value identifying the account ID. | [optional] 
**ResourceGroupId** | **string** | The short ID of the resource group. | [optional] 
**SourceCrn** | **string** | The CRN of resource instance or alias associated to the key. | [optional] 
**State** | **string** | The state of the key. | [optional] 
**Credentials** | [**Credentials**](Credentials.md) | The credentials for the key. Additional key-value pairs are passed through from the resource brokers.  Refer to service’s documentation for additional details. | [optional] 
**IamCompatible** | **bool** | Specifies whether the key’s credentials support IAM. | [optional] 
**ResourceInstanceUrl** | **string** | The relative path to the resource. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date when the key was created. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date when the key was last updated. | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) | The date when the key was deleted. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


