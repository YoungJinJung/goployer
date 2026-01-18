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

package deployer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DevopsArtFactory/goployer/pkg/schemas"
)

func TestRetrieveNextCapacity(t *testing.T) {
	tests := []struct {
		name                  string
		currentCapacity       schemas.Capacity
		targetCapacity        schemas.Capacity
		increaseInstanceCount int64
		expectedMin           int64
		expectedMax           int64
		expectedDesired       int64
		expectError           bool
	}{
		{
			name:                  "Increase capacity by 1",
			currentCapacity:       schemas.Capacity{Min: 1, Max: 1, Desired: 1},
			targetCapacity:        schemas.Capacity{Min: 3, Max: 3, Desired: 3},
			increaseInstanceCount: 1,
			expectedMin:           2,
			expectedMax:           2,
			expectedDesired:       2,
			expectError:           false,
		},
		{
			name:                  "Increase capacity by 2",
			currentCapacity:       schemas.Capacity{Min: 2, Max: 2, Desired: 2},
			targetCapacity:        schemas.Capacity{Min: 6, Max: 6, Desired: 6},
			increaseInstanceCount: 2,
			expectedMin:           4,
			expectedMax:           4,
			expectedDesired:       4,
			expectError:           false,
		},
		{
			name:                  "Already at target capacity",
			currentCapacity:       schemas.Capacity{Min: 3, Max: 3, Desired: 3},
			targetCapacity:        schemas.Capacity{Min: 3, Max: 3, Desired: 3},
			increaseInstanceCount: 1,
			expectedMin:           3,
			expectedMax:           3,
			expectedDesired:       3,
			expectError:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capacity := tt.currentCapacity
			err := RetrieveNextCapacity(&capacity, tt.targetCapacity, tt.increaseInstanceCount)

			if tt.expectError {
				assert.Error(t, err, "Should return error")
			} else {
				assert.NoError(t, err, "Should not return error")
				assert.Equal(t, tt.expectedMin, capacity.Min, "Min should match")
				assert.Equal(t, tt.expectedMax, capacity.Max, "Max should match")
				assert.Equal(t, tt.expectedDesired, capacity.Desired, "Desired should match")
			}
		})
	}
}

func TestIsFinishedRollingUpdate(t *testing.T) {
	tests := []struct {
		name            string
		currentCapacity schemas.Capacity
		targetCapacity  schemas.Capacity
		expected        bool
	}{
		{
			name:            "Rolling update finished",
			currentCapacity: schemas.Capacity{Min: 3, Max: 3, Desired: 3},
			targetCapacity:  schemas.Capacity{Min: 3, Max: 3, Desired: 3},
			expected:        true,
		},
		{
			name:            "Rolling update not finished - desired mismatch",
			currentCapacity: schemas.Capacity{Min: 2, Max: 3, Desired: 2},
			targetCapacity:  schemas.Capacity{Min: 3, Max: 3, Desired: 3},
			expected:        false,
		},
		{
			name:            "Rolling update not finished - min mismatch",
			currentCapacity: schemas.Capacity{Min: 2, Max: 3, Desired: 3},
			targetCapacity:  schemas.Capacity{Min: 3, Max: 3, Desired: 3},
			expected:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsFinishedRollingUpdate(tt.currentCapacity, tt.targetCapacity)
			assert.Equal(t, tt.expected, result, "IsFinishedRollingUpdate result should match expected")
		})
	}
}
