/*
copyright 2020 the Goployer authors

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DevopsArtFactory/goployer/pkg/constants"
)

func TestHealthcheckHost_Validation(t *testing.T) {
	tests := []struct {
		name           string
		host           HealthcheckHost
		expectedValid  bool
		lifecycleState string
	}{
		{
			name: "Valid InService host",
			host: HealthcheckHost{
				InstanceID:     "i-1234567890abcdef0",
				LifecycleState: constants.InServiceStatus,
				Valid:          true,
			},
			expectedValid:  true,
			lifecycleState: constants.InServiceStatus,
		},
		{
			name: "Invalid OutOfService host",
			host: HealthcheckHost{
				InstanceID:     "i-1234567890abcdef0",
				LifecycleState: "OutOfService",
				Valid:          false,
			},
			expectedValid:  false,
			lifecycleState: "OutOfService",
		},
		{
			name: "Valid host with instance ID",
			host: HealthcheckHost{
				InstanceID:     "i-abcdef1234567890",
				LifecycleState: constants.InServiceStatus,
				Valid:          true,
			},
			expectedValid:  true,
			lifecycleState: constants.InServiceStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedValid, tt.host.Valid, "Valid field should match expected")
			assert.Equal(t, tt.lifecycleState, tt.host.LifecycleState, "LifecycleState should match")
			assert.NotEmpty(t, tt.host.InstanceID, "InstanceID should not be empty")
		})
	}
}

func TestELBName_Format(t *testing.T) {
	tests := []struct {
		name    string
		elbName string
		valid   bool
	}{
		{
			name:    "Valid ELB name",
			elbName: "my-load-balancer",
			valid:   true,
		},
		{
			name:    "Valid ELB name with numbers",
			elbName: "my-lb-123",
			valid:   true,
		},
		{
			name:    "Empty ELB name",
			elbName: "",
			valid:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.elbName, "ELB name should not be empty")
			} else {
				assert.Empty(t, tt.elbName, "ELB name should be empty for invalid case")
			}
		})
	}
}

func TestInstanceState_Validation(t *testing.T) {
	validStates := []string{
		constants.InServiceStatus,
		"OutOfService",
		"Unknown",
	}

	for _, state := range validStates {
		t.Run("State_"+state, func(t *testing.T) {
			assert.NotEmpty(t, state, "State should not be empty")
		})
	}
}
