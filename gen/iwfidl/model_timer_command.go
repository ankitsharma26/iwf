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

// TimerCommand struct for TimerCommand
type TimerCommand struct {
	CommandId string `json:"commandId"`
	FiringUnixTimestampSeconds int64 `json:"firingUnixTimestampSeconds"`
}

// NewTimerCommand instantiates a new TimerCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerCommand(commandId string, firingUnixTimestampSeconds int64) *TimerCommand {
	this := TimerCommand{}
	this.CommandId = commandId
	this.FiringUnixTimestampSeconds = firingUnixTimestampSeconds
	return &this
}

// NewTimerCommandWithDefaults instantiates a new TimerCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerCommandWithDefaults() *TimerCommand {
	this := TimerCommand{}
	return &this
}

// GetCommandId returns the CommandId field value
func (o *TimerCommand) GetCommandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value
// and a boolean to check if the value has been set.
func (o *TimerCommand) GetCommandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandId, true
}

// SetCommandId sets field value
func (o *TimerCommand) SetCommandId(v string) {
	o.CommandId = v
}

// GetFiringUnixTimestampSeconds returns the FiringUnixTimestampSeconds field value
func (o *TimerCommand) GetFiringUnixTimestampSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FiringUnixTimestampSeconds
}

// GetFiringUnixTimestampSecondsOk returns a tuple with the FiringUnixTimestampSeconds field value
// and a boolean to check if the value has been set.
func (o *TimerCommand) GetFiringUnixTimestampSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiringUnixTimestampSeconds, true
}

// SetFiringUnixTimestampSeconds sets field value
func (o *TimerCommand) SetFiringUnixTimestampSeconds(v int64) {
	o.FiringUnixTimestampSeconds = v
}

func (o TimerCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commandId"] = o.CommandId
	}
	if true {
		toSerialize["firingUnixTimestampSeconds"] = o.FiringUnixTimestampSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableTimerCommand struct {
	value *TimerCommand
	isSet bool
}

func (v NullableTimerCommand) Get() *TimerCommand {
	return v.value
}

func (v *NullableTimerCommand) Set(val *TimerCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerCommand(val *TimerCommand) *NullableTimerCommand {
	return &NullableTimerCommand{value: val, isSet: true}
}

func (v NullableTimerCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

