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
// ResourceAliasesList A list of resource aliases.
type ResourceAliasesList struct {
	// The URL for requesting the next page of results.
	NextUrl string `json:"next_url"`
	// A list of resource aliases.
	Resources []ResourceAlias `json:"resources"`
	// The number of resource aliases in `resources`.
	RowsCount float32 `json:"rows_count"`
}