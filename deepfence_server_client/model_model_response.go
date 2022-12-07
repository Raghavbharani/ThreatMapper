/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ModelResponse struct for ModelResponse
type ModelResponse struct {
	Data *ModelResponseAccessToken `json:"data,omitempty"`
	ErrorFields map[string]string `json:"error_fields,omitempty"`
	Message *string `json:"message,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewModelResponse instantiates a new ModelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelResponse() *ModelResponse {
	this := ModelResponse{}
	return &this
}

// NewModelResponseWithDefaults instantiates a new ModelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelResponseWithDefaults() *ModelResponse {
	this := ModelResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModelResponse) GetData() ModelResponseAccessToken {
	if o == nil || isNil(o.Data) {
		var ret ModelResponseAccessToken
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponse) GetDataOk() (*ModelResponseAccessToken, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModelResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ModelResponseAccessToken and assigns it to the Data field.
func (o *ModelResponse) SetData(v ModelResponseAccessToken) {
	o.Data = &v
}

// GetErrorFields returns the ErrorFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelResponse) GetErrorFields() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ErrorFields
}

// GetErrorFieldsOk returns a tuple with the ErrorFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelResponse) GetErrorFieldsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ErrorFields) {
    return nil, false
	}
	return &o.ErrorFields, true
}

// HasErrorFields returns a boolean if a field has been set.
func (o *ModelResponse) HasErrorFields() bool {
	if o != nil && isNil(o.ErrorFields) {
		return true
	}

	return false
}

// SetErrorFields gets a reference to the given map[string]string and assigns it to the ErrorFields field.
func (o *ModelResponse) SetErrorFields(v map[string]string) {
	o.ErrorFields = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ModelResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ModelResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ModelResponse) SetMessage(v string) {
	o.Message = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ModelResponse) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ModelResponse) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ModelResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o ModelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.ErrorFields != nil {
		toSerialize["error_fields"] = o.ErrorFields
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableModelResponse struct {
	value *ModelResponse
	isSet bool
}

func (v NullableModelResponse) Get() *ModelResponse {
	return v.value
}

func (v *NullableModelResponse) Set(val *ModelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelResponse(val *ModelResponse) *NullableModelResponse {
	return &NullableModelResponse{value: val, isSet: true}
}

func (v NullableModelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


