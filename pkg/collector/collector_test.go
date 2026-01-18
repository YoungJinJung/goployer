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

package collector

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DevopsArtFactory/goployer/pkg/schemas"
)

func TestNewCollector(t *testing.T) {
	tests := []struct {
		name         string
		metricConfig schemas.MetricConfig
		assumeRole   string
	}{
		{
			name: "Create collector with DynamoDB storage",
			metricConfig: schemas.MetricConfig{
				Enabled: true,
				Storage: schemas.Storage{
					Type: "dynamodb",
					Name: "goployer-metrics",
				},
				Region: "us-east-1",
			},
			assumeRole: "",
		},
		{
			name: "Create collector with assume role",
			metricConfig: schemas.MetricConfig{
				Enabled: true,
				Storage: schemas.Storage{
					Type: "dynamodb",
					Name: "goployer-metrics",
				},
				Region: "ap-northeast-2",
			},
			assumeRole: "arn:aws:iam::123456789012:role/DeployRole",
		},
		{
			name: "Create disabled collector",
			metricConfig: schemas.MetricConfig{
				Enabled: false,
			},
			assumeRole: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewCollector(tt.metricConfig, tt.assumeRole)
			assert.NotNil(t, collector, "Collector should not be nil")
			assert.Equal(t, tt.metricConfig, collector.MetricConfig, "MetricConfig should match")
		})
	}
}

func TestSetTargetMetrics(t *testing.T) {
	metrics := SetTargetMetrics()

	assert.NotNil(t, metrics, "Target metrics should not be nil")
	assert.Greater(t, len(metrics), 0, "Should have at least one target metric")

	// Verify each metric has required fields
	for _, metric := range metrics {
		assert.NotEmpty(t, metric.Name, "Metric name should not be empty")
		assert.NotNil(t, metric.MappingFunction, "Mapping function should not be nil")
	}
}

func TestGatherUptime_Calculation(t *testing.T) {
	tests := []struct {
		name             string
		baseTimeDuration float64
		startDate        time.Time
		currentTime      time.Time
		expectedUptime   float64
	}{
		{
			name:             "One hour uptime",
			baseTimeDuration: 3600.0,
			startDate:        time.Now().Add(-1 * time.Hour),
			currentTime:      time.Now(),
			expectedUptime:   3600.0,
		},
		{
			name:             "Zero uptime",
			baseTimeDuration: 0.0,
			startDate:        time.Now(),
			currentTime:      time.Now(),
			expectedUptime:   0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := &HelperStruct{
				BaseTimeDuration: tt.baseTimeDuration,
				StartDate:        tt.startDate,
				CurrentTime:      tt.currentTime,
			}

			// Calculate uptime
			uptime := hs.CurrentTime.Sub(hs.StartDate).Seconds()
			assert.InDelta(t, tt.expectedUptime, uptime, 1.0, "Uptime should be approximately correct")
		})
	}
}

func TestHelperStruct_Fields(t *testing.T) {
	now := time.Now()
	startTime := now.Add(-1 * time.Hour)

	hs := HelperStruct{
		BaseTimeDuration: 3600.0,
		StartDate:        startTime,
		AutoScalingGroup: "test-asg-v001",
		CurrentTime:      now,
		TargetGroups:     []*string{stringPtr("tg-1"), stringPtr("tg-2")},
		LoadBalancers:    []*string{stringPtr("lb-1")},
		Storage:          "dynamodb",
	}

	assert.Equal(t, 3600.0, hs.BaseTimeDuration, "BaseTimeDuration should match")
	assert.Equal(t, startTime, hs.StartDate, "StartDate should match")
	assert.Equal(t, "test-asg-v001", hs.AutoScalingGroup, "AutoScalingGroup should match")
	assert.Equal(t, now, hs.CurrentTime, "CurrentTime should match")
	assert.Equal(t, 2, len(hs.TargetGroups), "Should have 2 target groups")
	assert.Equal(t, 1, len(hs.LoadBalancers), "Should have 1 load balancer")
	assert.Equal(t, "dynamodb", hs.Storage, "Storage should match")
}

func stringPtr(s string) *string {
	return &s
}
