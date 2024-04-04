/*
accent-auth

Accent's authentication service

API version: 0.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"encoding/json"
)

// checks if the UserEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEmail{}

// UserEmail struct for UserEmail
type UserEmail struct {
	Address   *string `json:"address,omitempty"`
	Confirmed *bool   `json:"confirmed,omitempty"`
	Main      *bool   `json:"main,omitempty"`
}

// NewUserEmail instantiates a new UserEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEmail() *UserEmail {
	this := UserEmail{}
	return &this
}

// NewUserEmailWithDefaults instantiates a new UserEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEmailWithDefaults() *UserEmail {
	this := UserEmail{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UserEmail) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEmail) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UserEmail) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UserEmail) SetAddress(v string) {
	o.Address = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *UserEmail) GetConfirmed() bool {
	if o == nil || IsNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEmail) GetConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.Confirmed) {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *UserEmail) HasConfirmed() bool {
	if o != nil && !IsNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *UserEmail) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *UserEmail) GetMain() bool {
	if o == nil || IsNil(o.Main) {
		var ret bool
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEmail) GetMainOk() (*bool, bool) {
	if o == nil || IsNil(o.Main) {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *UserEmail) HasMain() bool {
	if o != nil && !IsNil(o.Main) {
		return true
	}

	return false
}

// SetMain gets a reference to the given bool and assigns it to the Main field.
func (o *UserEmail) SetMain(v bool) {
	o.Main = &v
}

func (o UserEmail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !IsNil(o.Main) {
		toSerialize["main"] = o.Main
	}
	return toSerialize, nil
}

type NullableUserEmail struct {
	value *UserEmail
	isSet bool
}

func (v NullableUserEmail) Get() *UserEmail {
	return v.value
}

func (v *NullableUserEmail) Set(val *UserEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEmail(val *UserEmail) *NullableUserEmail {
	return &NullableUserEmail{value: val, isSet: true}
}

func (v NullableUserEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}