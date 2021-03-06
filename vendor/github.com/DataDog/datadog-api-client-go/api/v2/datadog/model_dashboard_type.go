/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// DashboardType The type of the dashboard.
type DashboardType string

// List of DashboardType
const (
	DASHBOARDTYPE_CUSTOM_TIMEBOARD        DashboardType = "custom_timeboard"
	DASHBOARDTYPE_CUSTOM_SCREENBOARD      DashboardType = "custom_screenboard"
	DASHBOARDTYPE_INTEGRATION_SCREENBOARD DashboardType = "integration_screenboard"
	DASHBOARDTYPE_INTEGRATION_TIMEBOARD   DashboardType = "integration_timeboard"
	DASHBOARDTYPE_HOST_TIMEBOARD          DashboardType = "host_timeboard"
)

// Ptr returns reference to DashboardType value
func (v DashboardType) Ptr() *DashboardType {
	return &v
}

type NullableDashboardType struct {
	value *DashboardType
	isSet bool
}

func (v NullableDashboardType) Get() *DashboardType {
	return v.value
}

func (v *NullableDashboardType) Set(val *DashboardType) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardType) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardType(val *DashboardType) *NullableDashboardType {
	return &NullableDashboardType{value: val, isSet: true}
}

func (v NullableDashboardType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
