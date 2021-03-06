/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gohiarc

// CreateFileRequest struct for CreateFileRequest
type CreateFileRequest struct {
	Key            string                 `json:"key,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	StorageService string                 `json:"storageService,omitempty"`
}
