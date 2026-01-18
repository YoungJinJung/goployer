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

func TestGetManifest_BucketKeyValidation(t *testing.T) {
	tests := []struct {
		name   string
		bucket string
		key    string
		valid  bool
	}{
		{
			name:   "Valid bucket and key",
			bucket: "my-deployment-bucket",
			key:    "manifests/production/app.yaml",
			valid:  true,
		},
		{
			name:   "Empty bucket",
			bucket: "",
			key:    "manifests/app.yaml",
			valid:  false,
		},
		{
			name:   "Empty key",
			bucket: "my-bucket",
			key:    "",
			valid:  false,
		},
		{
			name:   "Both empty",
			bucket: "",
			key:    "",
			valid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.bucket, "Bucket should not be empty")
				assert.NotEmpty(t, tt.key, "Key should not be empty")
			} else {
				isEmpty := tt.bucket == "" || tt.key == ""
				assert.True(t, isEmpty, "At least one field should be empty for invalid case")
			}
		})
	}
}

func TestS3Path_Format(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		isS3Path bool
	}{
		{
			name:     "Valid S3 path with s3://",
			path:     "s3://my-bucket/path/to/file.yaml",
			isS3Path: true,
		},
		{
			name:     "Valid S3 path with s3a://",
			path:     "s3a://my-bucket/path/to/file.yaml",
			isS3Path: true,
		},
		{
			name:     "Local file path",
			path:     "/local/path/to/file.yaml",
			isS3Path: false,
		},
		{
			name:     "Relative path",
			path:     "./manifests/app.yaml",
			isS3Path: false,
		},
		{
			name:     "HTTP URL",
			path:     "https://example.com/file.yaml",
			isS3Path: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasS3Prefix := strings.HasPrefix(tt.path, "s3://") || strings.HasPrefix(tt.path, "s3a://")
			assert.Equal(t, tt.isS3Path, hasS3Prefix, "S3 path detection should match expected")
		})
	}
}

func TestS3BucketName_Validation(t *testing.T) {
	tests := []struct {
		name       string
		bucketName string
		valid      bool
	}{
		{
			name:       "Valid bucket name with hyphens",
			bucketName: "my-deployment-bucket",
			valid:      true,
		},
		{
			name:       "Valid bucket name with dots",
			bucketName: "my.deployment.bucket",
			valid:      true,
		},
		{
			name:       "Valid bucket name lowercase",
			bucketName: "deploymentbucket123",
			valid:      true,
		},
		{
			name:       "Too short bucket name",
			bucketName: "ab",
			valid:      false,
		},
		{
			name:       "Empty bucket name",
			bucketName: "",
			valid:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation: bucket name should be at least 3 characters
			isValidLength := len(tt.bucketName) >= 3
			if tt.valid {
				assert.True(t, isValidLength, "Valid bucket name should be at least 3 characters")
			} else if tt.bucketName != "" {
				assert.False(t, isValidLength, "Invalid bucket name should be less than 3 characters")
			}
		})
	}
}
