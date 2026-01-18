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

package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DevopsArtFactory/goployer/pkg/constants"
)

func TestInitStartStatus(t *testing.T) {
	// Call the function
	status := InitStartStatus()

	// Verify the map is not nil
	assert.NotNil(t, status, "InitStartStatus should return a non-nil map")

	// Verify the map has exactly 8 entries
	expectedCount := 8
	assert.Equal(t, expectedCount, len(status), "InitStartStatus should return a map with %d entries", expectedCount)

	// Verify all expected keys exist and are set to false
	expectedKeys := []int64{
		constants.StepCheckPrevious,
		constants.StepDeploy,
		constants.StepAdditionalWork,
		constants.StepTriggerLifecycleCallback,
		constants.StepCleanPreviousVersion,
		constants.StepCleanChecking,
		constants.StepGatherMetrics,
		constants.StepRunAPI,
	}

	for _, key := range expectedKeys {
		value, exists := status[key]
		assert.True(t, exists, "Key %d should exist in the status map", key)
		assert.False(t, value, "Key %d should be initialized to false", key)
	}
}

func TestInitStartStatus_AllKeysFalse(t *testing.T) {
	status := InitStartStatus()

	// Verify all values are false
	for key, value := range status {
		assert.False(t, value, "All status values should be false, but key %d is true", key)
	}
}

func TestInitStartStatus_NoExtraKeys(t *testing.T) {
	status := InitStartStatus()

	// Define all valid keys
	validKeys := map[int64]bool{
		constants.StepCheckPrevious:            true,
		constants.StepDeploy:                   true,
		constants.StepAdditionalWork:           true,
		constants.StepTriggerLifecycleCallback: true,
		constants.StepCleanPreviousVersion:     true,
		constants.StepCleanChecking:            true,
		constants.StepGatherMetrics:            true,
		constants.StepRunAPI:                   true,
	}

	// Verify no extra keys exist
	for key := range status {
		assert.True(t, validKeys[key], "Unexpected key %d found in status map", key)
	}
}
