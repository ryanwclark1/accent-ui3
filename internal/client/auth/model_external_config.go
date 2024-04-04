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

// checks if the ExternalConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalConfig{}

// ExternalConfig struct for ExternalConfig
type ExternalConfig struct {
	// Client ID for the given authentication type. Required only for `google` and `microsoft` authentication types.
	ClientId *string `json:"client_id,omitempty"`
	// Client secret for the given authentication type. Required only for `google` and `microsoft` authentication types.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The API key to use for Firebase Cloud Messaging
	FcmApiKey *string `json:"fcm_api_key,omitempty"`
	// Public certificate to use for Apple Push Notification Service
	IosApnCertificate *string `json:"ios_apn_certificate,omitempty"`
	// Private key to use for Apple Push Notification Service
	IosApnPrivate *bool `json:"ios_apn_private,omitempty"`
	// Whether to use sandbox for Apple Push Notification Service
	UseSandbox *bool `json:"use_sandbox,omitempty"`
}

// NewExternalConfig instantiates a new ExternalConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalConfig() *ExternalConfig {
	this := ExternalConfig{}
	return &this
}

// NewExternalConfigWithDefaults instantiates a new ExternalConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalConfigWithDefaults() *ExternalConfig {
	this := ExternalConfig{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ExternalConfig) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConfig) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ExternalConfig) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ExternalConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ExternalConfig) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ExternalConfig) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ExternalConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetFcmApiKey returns the FcmApiKey field value if set, zero value otherwise.
func (o *ExternalConfig) GetFcmApiKey() string {
	if o == nil || IsNil(o.FcmApiKey) {
		var ret string
		return ret
	}
	return *o.FcmApiKey
}

// GetFcmApiKeyOk returns a tuple with the FcmApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConfig) GetFcmApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FcmApiKey) {
		return nil, false
	}
	return o.FcmApiKey, true
}

// HasFcmApiKey returns a boolean if a field has been set.
func (o *ExternalConfig) HasFcmApiKey() bool {
	if o != nil && !IsNil(o.FcmApiKey) {
		return true
	}

	return false
}

// SetFcmApiKey gets a reference to the given string and assigns it to the FcmApiKey field.
func (o *ExternalConfig) SetFcmApiKey(v string) {
	o.FcmApiKey = &v
}

// GetIosApnCertificate returns the IosApnCertificate field value if set, zero value otherwise.
func (o *ExternalConfig) GetIosApnCertificate() string {
	if o == nil || IsNil(o.IosApnCertificate) {
		var ret string
		return ret
	}
	return *o.IosApnCertificate
}

// GetIosApnCertificateOk returns a tuple with the IosApnCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConfig) GetIosApnCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.IosApnCertificate) {
		return nil, false
	}
	return o.IosApnCertificate, true
}

// HasIosApnCertificate returns a boolean if a field has been set.
func (o *ExternalConfig) HasIosApnCertificate() bool {
	if o != nil && !IsNil(o.IosApnCertificate) {
		return true
	}

	return false
}

// SetIosApnCertificate gets a reference to the given string and assigns it to the IosApnCertificate field.
func (o *ExternalConfig) SetIosApnCertificate(v string) {
	o.IosApnCertificate = &v
}

// GetIosApnPrivate returns the IosApnPrivate field value if set, zero value otherwise.
func (o *ExternalConfig) GetIosApnPrivate() bool {
	if o == nil || IsNil(o.IosApnPrivate) {
		var ret bool
		return ret
	}
	return *o.IosApnPrivate
}

// GetIosApnPrivateOk returns a tuple with the IosApnPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConfig) GetIosApnPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IosApnPrivate) {
		return nil, false
	}
	return o.IosApnPrivate, true
}

// HasIosApnPrivate returns a boolean if a field has been set.
func (o *ExternalConfig) HasIosApnPrivate() bool {
	if o != nil && !IsNil(o.IosApnPrivate) {
		return true
	}

	return false
}

// SetIosApnPrivate gets a reference to the given bool and assigns it to the IosApnPrivate field.
func (o *ExternalConfig) SetIosApnPrivate(v bool) {
	o.IosApnPrivate = &v
}

// GetUseSandbox returns the UseSandbox field value if set, zero value otherwise.
func (o *ExternalConfig) GetUseSandbox() bool {
	if o == nil || IsNil(o.UseSandbox) {
		var ret bool
		return ret
	}
	return *o.UseSandbox
}

// GetUseSandboxOk returns a tuple with the UseSandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConfig) GetUseSandboxOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSandbox) {
		return nil, false
	}
	return o.UseSandbox, true
}

// HasUseSandbox returns a boolean if a field has been set.
func (o *ExternalConfig) HasUseSandbox() bool {
	if o != nil && !IsNil(o.UseSandbox) {
		return true
	}

	return false
}

// SetUseSandbox gets a reference to the given bool and assigns it to the UseSandbox field.
func (o *ExternalConfig) SetUseSandbox(v bool) {
	o.UseSandbox = &v
}

func (o ExternalConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if !IsNil(o.FcmApiKey) {
		toSerialize["fcm_api_key"] = o.FcmApiKey
	}
	if !IsNil(o.IosApnCertificate) {
		toSerialize["ios_apn_certificate"] = o.IosApnCertificate
	}
	if !IsNil(o.IosApnPrivate) {
		toSerialize["ios_apn_private"] = o.IosApnPrivate
	}
	if !IsNil(o.UseSandbox) {
		toSerialize["use_sandbox"] = o.UseSandbox
	}
	return toSerialize, nil
}

type NullableExternalConfig struct {
	value *ExternalConfig
	isSet bool
}

func (v NullableExternalConfig) Get() *ExternalConfig {
	return v.value
}

func (v *NullableExternalConfig) Set(val *ExternalConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalConfig(val *ExternalConfig) *NullableExternalConfig {
	return &NullableExternalConfig{value: val, isSet: true}
}

func (v NullableExternalConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
