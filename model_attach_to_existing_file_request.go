/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gohiarc
// AttachToExistingFileRequest struct for AttachToExistingFileRequest
type AttachToExistingFileRequest struct {
	Name string `json:"name,omitempty"`
	StorageService string `json:"storageService,omitempty"`
	StorageId string `json:"storageId,omitempty"`
}
