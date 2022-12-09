/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ModelUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUser{}

// ModelUser struct for ModelUser
type ModelUser struct {
	Company *string `json:"company,omitempty"`
	CompanyId *int32 `json:"company_id,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	Groups map[string]string `json:"groups,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	PasswordInvalidated *bool `json:"password_invalidated,omitempty"`
	Role *string `json:"role,omitempty"`
	RoleId *int32 `json:"role_id,omitempty"`
}

// NewModelUser instantiates a new ModelUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUser() *ModelUser {
	this := ModelUser{}
	return &this
}

// NewModelUserWithDefaults instantiates a new ModelUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUserWithDefaults() *ModelUser {
	this := ModelUser{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ModelUser) GetCompany() string {
	if o == nil || isNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetCompanyOk() (*string, bool) {
	if o == nil || isNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ModelUser) HasCompany() bool {
	if o != nil && !isNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ModelUser) SetCompany(v string) {
	o.Company = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ModelUser) GetCompanyId() int32 {
	if o == nil || isNil(o.CompanyId) {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetCompanyIdOk() (*int32, bool) {
	if o == nil || isNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ModelUser) HasCompanyId() bool {
	if o != nil && !isNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *ModelUser) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ModelUser) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ModelUser) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ModelUser) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ModelUser) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ModelUser) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ModelUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUser) GetGroups() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUser) GetGroupsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Groups) {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ModelUser) HasGroups() bool {
	if o != nil && isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]string and assigns it to the Groups field.
func (o *ModelUser) SetGroups(v map[string]string) {
	o.Groups = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelUser) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelUser) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelUser) SetId(v int32) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ModelUser) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ModelUser) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ModelUser) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ModelUser) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ModelUser) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ModelUser) SetLastName(v string) {
	o.LastName = &v
}

// GetPasswordInvalidated returns the PasswordInvalidated field value if set, zero value otherwise.
func (o *ModelUser) GetPasswordInvalidated() bool {
	if o == nil || isNil(o.PasswordInvalidated) {
		var ret bool
		return ret
	}
	return *o.PasswordInvalidated
}

// GetPasswordInvalidatedOk returns a tuple with the PasswordInvalidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetPasswordInvalidatedOk() (*bool, bool) {
	if o == nil || isNil(o.PasswordInvalidated) {
		return nil, false
	}
	return o.PasswordInvalidated, true
}

// HasPasswordInvalidated returns a boolean if a field has been set.
func (o *ModelUser) HasPasswordInvalidated() bool {
	if o != nil && !isNil(o.PasswordInvalidated) {
		return true
	}

	return false
}

// SetPasswordInvalidated gets a reference to the given bool and assigns it to the PasswordInvalidated field.
func (o *ModelUser) SetPasswordInvalidated(v bool) {
	o.PasswordInvalidated = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ModelUser) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ModelUser) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ModelUser) SetRole(v string) {
	o.Role = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ModelUser) GetRoleId() int32 {
	if o == nil || isNil(o.RoleId) {
		var ret int32
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetRoleIdOk() (*int32, bool) {
	if o == nil || isNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ModelUser) HasRoleId() bool {
	if o != nil && !isNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given int32 and assigns it to the RoleId field.
func (o *ModelUser) SetRoleId(v int32) {
	o.RoleId = &v
}

func (o ModelUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !isNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !isNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !isNil(o.PasswordInvalidated) {
		toSerialize["password_invalidated"] = o.PasswordInvalidated
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.RoleId) {
		toSerialize["role_id"] = o.RoleId
	}
	return toSerialize, nil
}

type NullableModelUser struct {
	value *ModelUser
	isSet bool
}

func (v NullableModelUser) Get() *ModelUser {
	return v.value
}

func (v *NullableModelUser) Set(val *ModelUser) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUser) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUser(val *ModelUser) *NullableModelUser {
	return &NullableModelUser{value: val, isSet: true}
}

func (v NullableModelUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

