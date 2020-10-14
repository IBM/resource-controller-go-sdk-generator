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
// ResourceAliasPatch Updated property values for a resource alias.
type ResourceAliasPatch struct {
	// The new name of the alias. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`.
	Name string `json:"name"`
}
