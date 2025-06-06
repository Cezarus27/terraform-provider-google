// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/container/state_util.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package container

// A StateType represents the specific type of resting state that a state value
// is.
type StateType int

const (
	UndefinedState StateType = iota
	// A special resting state, that generally requires special consideration
	// Interactive states like PENDING_PARTNER in interconnects are an example
	RestingState
	// An error state is a state that indicates that a resource is not working
	// correctly. If this is Create, it should be tainted by returning an error
	ErrorState
	// A ready resource is fully provisioned, and ready to accept traffic/work
	ReadyState
)

type RestingStates map[string]StateType
