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

// checks if the TenantCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantCreate{}

// TenantCreate struct for TenantCreate
type TenantCreate struct {
	Address *TenantAddress `json:"address,omitempty"`
	// The contact user's UUID
	Contact *string `json:"contact,omitempty"`
	// A list containing human readeable unique domain names, associated with a specific tenant
	DomainNames []string `json:"domain_names,omitempty"`
	// The tenant's name
	Name *string `json:"name,omitempty"`
	// The tenant's contact phone number
	Phone *string `json:"phone,omitempty"`
	// A unique, human readeable identifier for this tenant. This field cannot be modified and will be auto-generated if missing.
	Slug *string `json:"slug,omitempty"`
	// The tenant's UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewTenantCreate instantiates a new TenantCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantCreate() *TenantCreate {
	this := TenantCreate{}
	return &this
}

// NewTenantCreateWithDefaults instantiates a new TenantCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantCreateWithDefaults() *TenantCreate {
	this := TenantCreate{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TenantCreate) GetAddress() TenantAddress {
	if o == nil || IsNil(o.Address) {
		var ret TenantAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreate) GetAddressOk() (*TenantAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TenantCreate) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given TenantAddress and assigns it to the Address field.
func (o *TenantCreate) SetAddress(v TenantAddress) {
	o.Address = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *TenantCreate) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreate) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *TenantCreate) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *TenantCreate) SetContact(v string) {
	o.Contact = &v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *TenantCreate) GetDomainNames() []string {
	if o == nil || IsNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreate) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *TenantCreate) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *TenantCreate) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TenantCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TenantCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TenantCreate) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *TenantCreate) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreate) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *TenantCreate) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *TenantCreate) SetPhone(v string) {
	o.Phone = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *TenantCreate) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreate) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *TenantCreate) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *TenantCreate) SetSlug(v string) {
	o.Slug = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *TenantCreate) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreate) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *TenantCreate) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *TenantCreate) SetUuid(v string) {
	o.Uuid = &v
}

func (o TenantCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.DomainNames) {
		toSerialize["domain_names"] = o.DomainNames
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableTenantCreate struct {
	value *TenantCreate
	isSet bool
}

func (v NullableTenantCreate) Get() *TenantCreate {
	return v.value
}

func (v *NullableTenantCreate) Set(val *TenantCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantCreate(val *TenantCreate) *NullableTenantCreate {
	return &NullableTenantCreate{value: val, isSet: true}
}

func (v NullableTenantCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}