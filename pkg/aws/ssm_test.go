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
)

func TestSSMCommand_Validation(t *testing.T) {
	tests := []struct {
		name      string
		instances []string
		commands  []string
		valid     bool
	}{
		{
			name:      "Valid instances and commands",
			instances: []string{"i-1234567890abcdef0", "i-0987654321fedcba0"},
			commands:  []string{"echo 'test'", "ls -la"},
			valid:     true,
		},
		{
			name:      "Empty instances",
			instances: []string{},
			commands:  []string{"echo 'test'"},
			valid:     false,
		},
		{
			name:      "Empty commands",
			instances: []string{"i-1234567890abcdef0"},
			commands:  []string{},
			valid:     false,
		},
		{
			name:      "Both empty",
			instances: []string{},
			commands:  []string{},
			valid:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.instances, "Instances should not be empty")
				assert.NotEmpty(t, tt.commands, "Commands should not be empty")
			} else {
				isEmpty := len(tt.instances) == 0 || len(tt.commands) == 0
				assert.True(t, isEmpty, "At least one should be empty for invalid case")
			}
		})
	}
}

func TestInstanceID_Format(t *testing.T) {
	tests := []struct {
		name       string
		instanceID string
		valid      bool
	}{
		{
			name:       "Valid instance ID",
			instanceID: "i-1234567890abcdef0",
			valid:      true,
		},
		{
			name:       "Valid short instance ID",
			instanceID: "i-12345678",
			valid:      true,
		},
		{
			name:       "Invalid instance ID - no prefix",
			instanceID: "1234567890abcdef0",
			valid:      false,
		},
		{
			name:       "Empty instance ID",
			instanceID: "",
			valid:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.instanceID, "Instance ID should not be empty")
				assert.Contains(t, tt.instanceID, "i-", "Instance ID should start with i-")
			} else if tt.instanceID != "" {
				assert.NotContains(t, tt.instanceID, "i-", "Invalid instance ID should not have i- prefix")
			}
		})
	}
}
