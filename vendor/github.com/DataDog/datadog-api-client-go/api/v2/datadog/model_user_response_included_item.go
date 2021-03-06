/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// UserResponseIncludedItem An object related to a user.
type UserResponseIncludedItem struct {
	UserResponseIncludedItemInterface interface{ GetType() string }
}

func (s UserResponseIncludedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.UserResponseIncludedItemInterface)
}

func (s *UserResponseIncludedItem) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaled map[string]interface{}
	err = json.Unmarshal(src, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["type"]; ok {
		switch v {
		case "Organization":
			var result *Organization = &Organization{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.UserResponseIncludedItemInterface = result
			return nil
		case "Permission":
			var result *Permission = &Permission{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.UserResponseIncludedItemInterface = result
			return nil
		case "Role":
			var result *Role = &Role{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.UserResponseIncludedItemInterface = result
			return nil
		case "orgs":
			var result *Organization = &Organization{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.UserResponseIncludedItemInterface = result
			return nil
		case "permissions":
			var result *Permission = &Permission{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.UserResponseIncludedItemInterface = result
			return nil
		case "roles":
			var result *Role = &Role{}
			err = json.Unmarshal(src, result)
			if err != nil {
				return err
			}
			s.UserResponseIncludedItemInterface = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'type' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'type' not found in unmarshaled payload: %+v", unmarshaled)
	}
}

type NullableUserResponseIncludedItem struct {
	value *UserResponseIncludedItem
	isSet bool
}

func (v NullableUserResponseIncludedItem) Get() *UserResponseIncludedItem {
	return v.value
}

func (v *NullableUserResponseIncludedItem) Set(val *UserResponseIncludedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseIncludedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseIncludedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseIncludedItem(val *UserResponseIncludedItem) *NullableUserResponseIncludedItem {
	return &NullableUserResponseIncludedItem{value: val, isSet: true}
}

func (v NullableUserResponseIncludedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseIncludedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
