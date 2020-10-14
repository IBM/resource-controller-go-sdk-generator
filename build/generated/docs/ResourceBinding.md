# ResourceBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID associated with the binding. | [optional] 
**Guid** | **string** | When you create a new binding, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the binding. | [optional] 
**Crn** | **string** | The full Cloud Resource Name (CRN) associated with the binding. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic&#x3D;overview-crn). | [optional] 
**Url** | **string** | When you provision a new binding, a relative URL path is created identifying the location of the binding. | [optional] 
**Name** | **string** | The human-readable name of the binding. | [optional] 
**AccountId** | **string** | An alpha-numeric value identifying the account ID. | [optional] 
**ResourceGroupId** | **string** | The short ID of the resource group. | [optional] 
**SourceCrn** | **string** | The CRN of resource alias associated to the binding. | [optional] 
**TargetCrn** | **string** | The CRN of target resource, e.g. application, in a specific environment. | [optional] 
**RegionBindingId** | **string** | The short ID of the binding in specific targeted environment, e.g. &#x60;service_binding_id&#x60; in a given IBM Cloud environment. | [optional] 
**State** | **string** | The state of the binding. | [optional] 
**Credentials** | [**Credentials**](Credentials.md) | The credentials for the binding. Additional key-value pairs are passed through from the resource brokers.  For additional details, see the service’s documentation. | [optional] 
**IamCompatible** | **bool** | Specifies whether the binding’s credentials support IAM. | [optional] 
**ResourceAliasUrl** | **string** | The relative path to the resource alias that this binding is associated with. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The date when the binding was created. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The date when the binding was last updated. | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) | The date when the binding was deleted. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


