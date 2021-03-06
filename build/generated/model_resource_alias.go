/*
 * (C) Copyright IBM Corp. 2020
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * IBM Cloud Resource Controller API
 *
 * Manage lifecycle of your Cloud resources using Resource Controller APIs. Resources are provisioned globally in an account scope. Supports asynchronous provisioning of resources. Enables consumption of a global resource through a Cloud Foundry space in any region.
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package generated
import (
	"time"
)
// ResourceAlias A resource alias.
type ResourceAlias struct {
	// The ID associated with the alias.
	Id string `json:"id,omitempty"`
	// When you create a new alias, a globally unique identifier (GUID) is assigned. This GUID is a unique internal indentifier managed by the resource controller that corresponds to the alias.
	Guid string `json:"guid,omitempty"`
	// The full Cloud Resource Name (CRN) associated with the alias. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).
	Crn string `json:"crn,omitempty"`
	// When you created a new alias, a relative URL path is created identifying the location of the alias.
	Url string `json:"url,omitempty"`
	// The human-readable name of the alias.
	Name string `json:"name,omitempty"`
	// An alpha-numeric value identifying the account ID.
	AccountId string `json:"account_id,omitempty"`
	// The short ID of the resource group.
	ResourceGroupId string `json:"resource_group_id,omitempty"`
	// The long ID (full CRN) of the resource group.
	ResourceGroupCrn string `json:"resource_group_crn,omitempty"`
	// The CRN of the target namespace in the specific environment.
	TargetCrn string `json:"target_crn,omitempty"`
	// The state of the alias.
	State string `json:"state,omitempty"`
	// The short ID of the resource instance that is being aliased.
	ResourceInstanceId string `json:"resource_instance_id,omitempty"`
	// The short ID of the instance in the specific target environment, e.g. `service_instance_id` in a given IBM Cloud environment.
	RegionInstanceId string `json:"region_instance_id,omitempty"`
	// The relative path to the instance.
	ResourceInstanceUrl string `json:"resource_instance_url,omitempty"`
	// The relative path to the resource bindings for the alias.
	ResourceBindingsUrl string `json:"resource_bindings_url,omitempty"`
	// The relative path to the resource keys for the alias.
	ResourceKeysUrl string `json:"resource_keys_url,omitempty"`
	// The date when the alias was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The date when the alias was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The date when the alias was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
