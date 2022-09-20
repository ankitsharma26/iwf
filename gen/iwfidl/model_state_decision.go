/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// StateDecision struct for StateDecision
type StateDecision struct {
	WaitForMoreCommandResults *bool `json:"waitForMoreCommandResults,omitempty"`
	NextStates []StateMovement `json:"nextStates,omitempty"`
	UpsertSearchAttributes []SearchAttribute `json:"upsertSearchAttributes,omitempty"`
	UpsertQueryAttributes []KeyValue `json:"upsertQueryAttributes,omitempty"`
}

// NewStateDecision instantiates a new StateDecision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateDecision() *StateDecision {
	this := StateDecision{}
	return &this
}

// NewStateDecisionWithDefaults instantiates a new StateDecision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateDecisionWithDefaults() *StateDecision {
	this := StateDecision{}
	return &this
}

// GetWaitForMoreCommandResults returns the WaitForMoreCommandResults field value if set, zero value otherwise.
func (o *StateDecision) GetWaitForMoreCommandResults() bool {
	if o == nil || o.WaitForMoreCommandResults == nil {
		var ret bool
		return ret
	}
	return *o.WaitForMoreCommandResults
}

// GetWaitForMoreCommandResultsOk returns a tuple with the WaitForMoreCommandResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateDecision) GetWaitForMoreCommandResultsOk() (*bool, bool) {
	if o == nil || o.WaitForMoreCommandResults == nil {
		return nil, false
	}
	return o.WaitForMoreCommandResults, true
}

// HasWaitForMoreCommandResults returns a boolean if a field has been set.
func (o *StateDecision) HasWaitForMoreCommandResults() bool {
	if o != nil && o.WaitForMoreCommandResults != nil {
		return true
	}

	return false
}

// SetWaitForMoreCommandResults gets a reference to the given bool and assigns it to the WaitForMoreCommandResults field.
func (o *StateDecision) SetWaitForMoreCommandResults(v bool) {
	o.WaitForMoreCommandResults = &v
}

// GetNextStates returns the NextStates field value if set, zero value otherwise.
func (o *StateDecision) GetNextStates() []StateMovement {
	if o == nil || o.NextStates == nil {
		var ret []StateMovement
		return ret
	}
	return o.NextStates
}

// GetNextStatesOk returns a tuple with the NextStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateDecision) GetNextStatesOk() ([]StateMovement, bool) {
	if o == nil || o.NextStates == nil {
		return nil, false
	}
	return o.NextStates, true
}

// HasNextStates returns a boolean if a field has been set.
func (o *StateDecision) HasNextStates() bool {
	if o != nil && o.NextStates != nil {
		return true
	}

	return false
}

// SetNextStates gets a reference to the given []StateMovement and assigns it to the NextStates field.
func (o *StateDecision) SetNextStates(v []StateMovement) {
	o.NextStates = v
}

// GetUpsertSearchAttributes returns the UpsertSearchAttributes field value if set, zero value otherwise.
func (o *StateDecision) GetUpsertSearchAttributes() []SearchAttribute {
	if o == nil || o.UpsertSearchAttributes == nil {
		var ret []SearchAttribute
		return ret
	}
	return o.UpsertSearchAttributes
}

// GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateDecision) GetUpsertSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || o.UpsertSearchAttributes == nil {
		return nil, false
	}
	return o.UpsertSearchAttributes, true
}

// HasUpsertSearchAttributes returns a boolean if a field has been set.
func (o *StateDecision) HasUpsertSearchAttributes() bool {
	if o != nil && o.UpsertSearchAttributes != nil {
		return true
	}

	return false
}

// SetUpsertSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the UpsertSearchAttributes field.
func (o *StateDecision) SetUpsertSearchAttributes(v []SearchAttribute) {
	o.UpsertSearchAttributes = v
}

// GetUpsertQueryAttributes returns the UpsertQueryAttributes field value if set, zero value otherwise.
func (o *StateDecision) GetUpsertQueryAttributes() []KeyValue {
	if o == nil || o.UpsertQueryAttributes == nil {
		var ret []KeyValue
		return ret
	}
	return o.UpsertQueryAttributes
}

// GetUpsertQueryAttributesOk returns a tuple with the UpsertQueryAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateDecision) GetUpsertQueryAttributesOk() ([]KeyValue, bool) {
	if o == nil || o.UpsertQueryAttributes == nil {
		return nil, false
	}
	return o.UpsertQueryAttributes, true
}

// HasUpsertQueryAttributes returns a boolean if a field has been set.
func (o *StateDecision) HasUpsertQueryAttributes() bool {
	if o != nil && o.UpsertQueryAttributes != nil {
		return true
	}

	return false
}

// SetUpsertQueryAttributes gets a reference to the given []KeyValue and assigns it to the UpsertQueryAttributes field.
func (o *StateDecision) SetUpsertQueryAttributes(v []KeyValue) {
	o.UpsertQueryAttributes = v
}

func (o StateDecision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WaitForMoreCommandResults != nil {
		toSerialize["waitForMoreCommandResults"] = o.WaitForMoreCommandResults
	}
	if o.NextStates != nil {
		toSerialize["nextStates"] = o.NextStates
	}
	if o.UpsertSearchAttributes != nil {
		toSerialize["upsertSearchAttributes"] = o.UpsertSearchAttributes
	}
	if o.UpsertQueryAttributes != nil {
		toSerialize["upsertQueryAttributes"] = o.UpsertQueryAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableStateDecision struct {
	value *StateDecision
	isSet bool
}

func (v NullableStateDecision) Get() *StateDecision {
	return v.value
}

func (v *NullableStateDecision) Set(val *StateDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableStateDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableStateDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateDecision(val *StateDecision) *NullableStateDecision {
	return &NullableStateDecision{value: val, isSet: true}
}

func (v NullableStateDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


