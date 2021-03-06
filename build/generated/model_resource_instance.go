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
// ResourceInstance A resource instance.
type ResourceInstance struct {
	// The ID associated with the instance.
	Id string `json:"id,omitempty"`
	// When you create a new resource, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the instance.
	Guid string `json:"guid,omitempty"`
	// The full Cloud Resource Name (CRN) associated with the instance. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).
	Crn string `json:"crn,omitempty"`
	// When you provision a new resource, a relative URL path is created identifying the location of the instance.
	Url string `json:"url,omitempty"`
	// The human-readable name of the instance.
	Name string `json:"name,omitempty"`
	// An alpha-numeric value identifying the account ID.
	AccountId string `json:"account_id,omitempty"`
	// The short ID of the resource group.
	ResourceGroupId string `json:"resource_group_id,omitempty"`
	// The long ID (full CRN) of the resource group.
	ResourceGroupCrn string `json:"resource_group_crn,omitempty"`
	// The unique ID of the offering. This value is provided by and stored in the global catalog.
	ResourceId string `json:"resource_id,omitempty"`
	// The unique ID of the plan associated with the offering. This value is provided by and stored in the global catalog.
	ResourcePlanId string `json:"resource_plan_id,omitempty"`
	// The full deployment CRN as defined in the global catalog. The Cloud Resource Name (CRN) of the deployment location where the instance is provisioned.
	TargetCrn string `json:"target_crn,omitempty"`
	// The current state of the instance. For example, if the instance is deleted, it will return removed.
	State string `json:"state,omitempty"`
	// The type of the instance, e.g. `service_instance`.
	Type string `json:"type,omitempty"`
	// The sub-type of instance, e.g. `cfaas`.
	SubType string `json:"sub_type,omitempty"`
	// The status of the last operation requested on the instance.
	LastOperation map[string]interface{} `json:"last_operation,omitempty"`
	// The resource-broker-provided URL to access administrative features of the instance.
	DashboardUrl string `json:"dashboard_url,omitempty"`
	// The plan history of the instance.
	PlanHistory []PlanHistoryItem `json:"plan_history,omitempty"`
	// The relative path to the resource aliases for the instance.
	ResourceAliasesUrl string `json:"resource_aliases_url,omitempty"`
	// The relative path to the resource bindings for the instance.
	ResourceBindingsUrl string `json:"resource_bindings_url,omitempty"`
	// The relative path to the resource keys for the instance.
	ResourceKeysUrl string `json:"resource_keys_url,omitempty"`
	// The date when the instance was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The date when the instance was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The date when the instance was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
