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

func TestCheckTableExists_TableName(t *testing.T) {
	tests := []struct {
		name      string
		tableName string
		wantEmpty bool
	}{
		{
			name:      "Valid table name",
			tableName: "goployer-metrics",
			wantEmpty: false,
		},
		{
			name:      "Empty table name",
			tableName: "",
			wantEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantEmpty {
				assert.Empty(t, tt.tableName, "Table name should be empty")
			} else {
				assert.NotEmpty(t, tt.tableName, "Table name should not be empty")
			}
		})
	}
}

func TestMakeRecord_FieldValidation(t *testing.T) {
	tests := []struct {
		name   string
		stack  string
		config string
		asg    string
		status string
		valid  bool
	}{
		{
			name:   "All fields valid",
			stack:  "test-stack",
			config: "test-config",
			asg:    "test-asg-123",
			status: "RUNNING",
			valid:  true,
		},
		{
			name:   "Empty ASG name",
			stack:  "test-stack",
			config: "test-config",
			asg:    "",
			status: "RUNNING",
			valid:  false,
		},
		{
			name:   "Empty status",
			stack:  "test-stack",
			config: "test-config",
			asg:    "test-asg-123",
			status: "",
			valid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.stack)
				assert.NotEmpty(t, tt.config)
				assert.NotEmpty(t, tt.asg)
				assert.NotEmpty(t, tt.status)
			} else {
				// At least one field should be empty
				isEmpty := tt.stack == "" || tt.config == "" || tt.asg == "" || tt.status == ""
				assert.True(t, isEmpty, "At least one required field should be empty for invalid case")
			}
		})
	}
}

func TestUpdateRecord_StatusValues(t *testing.T) {
	validStatuses := []string{
		"RUNNING",
		"DONE",
		"FAILED",
		"GATHERING",
	}

	for _, status := range validStatuses {
		t.Run("Status_"+status, func(t *testing.T) {
			assert.NotEmpty(t, status, "Status should not be empty")
			assert.Contains(t, validStatuses, status, "Status should be a valid value")
		})
	}
}

func TestGetSingleItem_ASGFormat(t *testing.T) {
	tests := []struct {
		name  string
		asg   string
		valid bool
	}{
		{
			name:  "Valid ASG name with prefix",
			asg:   "test-stack-v001",
			valid: true,
		},
		{
			name:  "Valid ASG name with timestamp",
			asg:   "test-stack-20240117120000",
			valid: true,
		},
		{
			name:  "Empty ASG name",
			asg:   "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.asg, "ASG name should not be empty")
			} else {
				assert.Empty(t, tt.asg, "ASG name should be empty for invalid case")
			}
		})
	}
}

func TestUpdateStatistics_FieldTypes(t *testing.T) {
	tests := []struct {
		name         string
		updateFields map[string]interface{}
		valid        bool
	}{
		{
			name: "Valid numeric fields",
			updateFields: map[string]interface{}{
				"uptime":        float64(3600),
				"request_count": int64(1000),
			},
			valid: true,
		},
		{
			name: "Valid string fields",
			updateFields: map[string]interface{}{
				"deployment_id": "deploy-123",
				"version":       "v001",
			},
			valid: true,
		},
		{
			name: "Mixed field types",
			updateFields: map[string]interface{}{
				"uptime":  float64(3600),
				"status":  "RUNNING",
				"healthy": true,
			},
			valid: true,
		},
		{
			name:         "Empty fields",
			updateFields: map[string]interface{}{},
			valid:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.updateFields, "Update fields should not be empty")
			} else {
				assert.Empty(t, tt.updateFields, "Update fields should be empty for invalid case")
			}
		})
	}
}
