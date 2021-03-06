/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// AdditionalPropertiesArray struct for AdditionalPropertiesArray
type AdditionalPropertiesArray struct {
	Name *string `json:"name,omitempty"`
}

// NewAdditionalPropertiesArray instantiates a new AdditionalPropertiesArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalPropertiesArray() *AdditionalPropertiesArray {
    this := AdditionalPropertiesArray{}
    return &this
}

// NewAdditionalPropertiesArrayWithDefaults instantiates a new AdditionalPropertiesArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalPropertiesArrayWithDefaults() *AdditionalPropertiesArray {
    this := AdditionalPropertiesArray{}
    return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdditionalPropertiesArray) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalPropertiesArray) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdditionalPropertiesArray) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdditionalPropertiesArray) SetName(v string) {
	o.Name = &v
}

type NullableAdditionalPropertiesArray struct {
	Value AdditionalPropertiesArray
	ExplicitNull bool
}

func (v NullableAdditionalPropertiesArray) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAdditionalPropertiesArray) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
