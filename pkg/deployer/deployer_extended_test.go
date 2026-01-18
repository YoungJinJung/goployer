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

	"github.com/DevopsArtFactory/goployer/pkg/aws"
	"github.com/DevopsArtFactory/goployer/pkg/constants"
	"github.com/DevopsArtFactory/goployer/pkg/schemas"
)

func TestGetValidHostCount(t *testing.T) {
	tests := []struct {
		name          string
		hosts         []aws.HealthcheckHost
		expectedCount int64
	}{
		{
			name: "All hosts valid",
			hosts: []aws.HealthcheckHost{
				{InstanceID: "i-001", Valid: true},
				{InstanceID: "i-002", Valid: true},
				{InstanceID: "i-003", Valid: true},
			},
			expectedCount: 3,
		},
		{
			name: "Some hosts invalid",
			hosts: []aws.HealthcheckHost{
				{InstanceID: "i-001", Valid: true},
				{InstanceID: "i-002", Valid: false},
				{InstanceID: "i-003", Valid: true},
			},
			expectedCount: 2,
		},
		{
			name: "No valid hosts",
			hosts: []aws.HealthcheckHost{
				{InstanceID: "i-001", Valid: false},
				{InstanceID: "i-002", Valid: false},
			},
			expectedCount: 0,
		},
		{
			name:          "Empty host list",
			hosts:         []aws.HealthcheckHost{},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deployer := &Deployer{}
			count := deployer.GetValidHostCount(tt.hosts)
			assert.Equal(t, tt.expectedCount, count, "Valid host count should match expected")
		})
	}
}

func TestCompareWithCurrentCapacity(t *testing.T) {
	tests := []struct {
		name                  string
		forceManifestCapacity bool
		manifestCapacity      schemas.Capacity
		prevCapacity          schemas.Capacity
		expectedCapacity      schemas.Capacity
		region                string
	}{
		{
			name:                  "Force manifest capacity",
			forceManifestCapacity: true,
			manifestCapacity:      schemas.Capacity{Min: 2, Max: 4, Desired: 3},
			prevCapacity:          schemas.Capacity{Min: 1, Max: 2, Desired: 1},
			expectedCapacity:      schemas.Capacity{Min: 2, Max: 4, Desired: 3},
			region:                "us-east-1",
		},
		{
			name:                  "Use previous capacity when not forced",
			forceManifestCapacity: false,
			manifestCapacity:      schemas.Capacity{Min: 1, Max: 2, Desired: 1},
			prevCapacity:          schemas.Capacity{Min: 3, Max: 6, Desired: 4},
			expectedCapacity:      schemas.Capacity{Min: 3, Max: 6, Desired: 4},
			region:                "us-east-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deployer := &Deployer{
				Stack: schemas.Stack{
					Capacity: tt.manifestCapacity,
				},
				PrevInstanceCount: map[string]schemas.Capacity{
					tt.region: tt.prevCapacity,
				},
			}

			result := deployer.CompareWithCurrentCapacity(tt.forceManifestCapacity, tt.region)
			assert.Equal(t, tt.expectedCapacity.Min, result.Min, "Min capacity should match")
			assert.Equal(t, tt.expectedCapacity.Max, result.Max, "Max capacity should match")
			assert.Equal(t, tt.expectedCapacity.Desired, result.Desired, "Desired capacity should match")
		})
	}
}

func TestSelectClientFromList(t *testing.T) {
	tests := []struct {
		name        string
		clients     []aws.Client
		region      string
		expectError bool
	}{
		{
			name: "Find matching client",
			clients: []aws.Client{
				{Region: "us-east-1"},
				{Region: "us-west-2"},
			},
			region:      "us-west-2",
			expectError: false,
		},
		{
			name: "Client not found",
			clients: []aws.Client{
				{Region: "us-east-1"},
			},
			region:      "ap-northeast-2",
			expectError: true,
		},
		{
			name:        "Empty client list",
			clients:     []aws.Client{},
			region:      "us-east-1",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := selectClientFromList(tt.clients, tt.region)
			if tt.expectError {
				assert.Error(t, err, "Should return error when client not found")
			} else {
				assert.NoError(t, err, "Should not return error when client found")
				assert.Equal(t, tt.region, client.Region, "Client region should match requested region")
			}
		})
	}
}

func TestCheckSpotInstanceOption_Extended(t *testing.T) {
	tests := []struct {
		name             string
		overrideTypes    string
		instanceTypeList []string
		instanceType     string
		expectError      bool
	}{
		{
			name:             "Valid spot instance types",
			overrideTypes:    "t2.small|t3.small",
			instanceTypeList: []string{},
			instanceType:     "t2.micro",
			expectError:      false,
		},
		{
			name:             "Empty override types",
			overrideTypes:    "",
			instanceTypeList: []string{},
			instanceType:     "t2.micro",
			expectError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkSpotInstanceOption(tt.overrideTypes, tt.instanceTypeList, tt.instanceType)
			if tt.expectError {
				assert.Error(t, err, "Should return error for invalid spot instance configuration")
			} else {
				assert.NoError(t, err, "Should not return error for valid spot instance configuration")
			}
		})
	}
}

func TestDeployer_GetStackName(t *testing.T) {
	tests := []struct {
		name         string
		stackName    string
		expectedName string
	}{
		{
			name:         "Simple stack name",
			stackName:    "my-app",
			expectedName: "my-app",
		},
		{
			name:         "Stack name with environment",
			stackName:    "my-app-production",
			expectedName: "my-app-production",
		},
		{
			name:         "Empty stack name",
			stackName:    "",
			expectedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deployer := &Deployer{
				Stack: schemas.Stack{
					Stack: tt.stackName,
				},
			}
			result := deployer.GetStackName()
			assert.Equal(t, tt.expectedName, result, "Stack name should match")
		})
	}
}

func TestDeployer_SkipDeployStep(t *testing.T) {
	deployer := &Deployer{
		StepStatus: map[int64]bool{
			constants.StepDeploy: false,
		},
	}

	// Initially deploy step should be false
	assert.False(t, deployer.StepStatus[constants.StepDeploy], "Deploy step should be false initially")

	// Skip the deploy step
	deployer.SkipDeployStep()

	// After skipping, deploy step should be true
	assert.True(t, deployer.StepStatus[constants.StepDeploy], "Deploy step should be true after skipping")
}
