/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gohiarc
import (
	"os"
)
// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	Request string `json:"request,omitempty"`
	File *os.File `json:"file,omitempty"`
}
