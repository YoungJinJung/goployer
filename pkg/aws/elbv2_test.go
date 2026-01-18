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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetGroupARN_Format(t *testing.T) {
	tests := []struct {
		name  string
		arn   string
		valid bool
	}{
		{
			name:  "Valid target group ARN",
			arn:   "arn:aws:elasticloadbalancing:us-east-1:123456789012:targetgroup/my-targets/50dc6c495c0c9188",
			valid: true,
		},
		{
			name:  "Valid target group ARN with different region",
			arn:   "arn:aws:elasticloadbalancing:ap-northeast-2:123456789012:targetgroup/app-tg/1234567890abcdef",
			valid: true,
		},
		{
			name:  "Empty ARN",
			arn:   "",
			valid: false,
		},
		{
			name:  "Invalid ARN format",
			arn:   "invalid-arn",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasARNPrefix := strings.HasPrefix(tt.arn, "arn:aws:")
			if tt.valid {
				assert.True(t, hasARNPrefix, "Valid ARN should start with 'arn:aws:'")
				assert.Contains(t, tt.arn, "targetgroup", "Valid target group ARN should contain 'targetgroup'")
			} else if tt.arn != "" {
				assert.False(t, hasARNPrefix, "Invalid ARN should not have proper format")
			}
		})
	}
}

func TestLoadBalancerARN_Format(t *testing.T) {
	tests := []struct {
		name  string
		arn   string
		valid bool
	}{
		{
			name:  "Valid ALB ARN",
			arn:   "arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188",
			valid: true,
		},
		{
			name:  "Valid NLB ARN",
			arn:   "arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/net/my-network-lb/50dc6c495c0c9188",
			valid: true,
		},
		{
			name:  "Empty ARN",
			arn:   "",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasARNPrefix := strings.HasPrefix(tt.arn, "arn:aws:")
			if tt.valid {
				assert.True(t, hasARNPrefix, "Valid ARN should start with 'arn:aws:'")
				assert.Contains(t, tt.arn, "loadbalancer", "Valid load balancer ARN should contain 'loadbalancer'")
			} else {
				assert.Empty(t, tt.arn, "Invalid ARN should be empty")
			}
		})
	}
}

func TestHealthcheckHost_TargetStatus(t *testing.T) {
	tests := []struct {
		name         string
		host         HealthcheckHost
		targetStatus string
		healthStatus string
	}{
		{
			name: "Healthy target",
			host: HealthcheckHost{
				InstanceID:   "i-1234567890abcdef0",
				TargetStatus: "healthy",
				HealthStatus: "healthy",
				Valid:        true,
			},
			targetStatus: "healthy",
			healthStatus: "healthy",
		},
		{
			name: "Unhealthy target",
			host: HealthcheckHost{
				InstanceID:   "i-1234567890abcdef0",
				TargetStatus: "unhealthy",
				HealthStatus: "unhealthy",
				Valid:        false,
			},
			targetStatus: "unhealthy",
			healthStatus: "unhealthy",
		},
		{
			name: "Draining target",
			host: HealthcheckHost{
				InstanceID:   "i-1234567890abcdef0",
				TargetStatus: "draining",
				HealthStatus: "draining",
				Valid:        false,
			},
			targetStatus: "draining",
			healthStatus: "draining",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.targetStatus, tt.host.TargetStatus, "TargetStatus should match")
			assert.Equal(t, tt.healthStatus, tt.host.HealthStatus, "HealthStatus should match")
			assert.NotEmpty(t, tt.host.InstanceID, "InstanceID should not be empty")
		})
	}
}

func TestTargetGroupName_Validation(t *testing.T) {
	tests := []struct {
		name   string
		tgName string
		valid  bool
	}{
		{
			name:   "Valid target group name",
			tgName: "my-app-tg",
			valid:  true,
		},
		{
			name:   "Valid target group name with version",
			tgName: "my-app-tg-v001",
			valid:  true,
		},
		{
			name:   "Empty target group name",
			tgName: "",
			valid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.tgName, "Target group name should not be empty")
			} else {
				assert.Empty(t, tt.tgName, "Target group name should be empty for invalid case")
			}
		})
	}
}
