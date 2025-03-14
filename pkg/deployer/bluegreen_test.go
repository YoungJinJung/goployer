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

	"github.com/DevopsArtFactory/goployer/pkg/schemas"
)

func TestGetStackName(t *testing.T) {
	b := BlueGreen{&Deployer{Stack: schemas.Stack{Stack: "Test"}}}

	input := b.GetStackName()
	expected := "Test"

	if input != expected {
		t.Error(input)
	}
}

func TestCheckRegionExist(t *testing.T) {
	target := "ap-northeast-2"
	regionList := []schemas.RegionConfig{
		{
			Region: "us-east-1",
		},
	}
	input := CheckRegionExist(target, regionList)
	expected := false

	if input != expected {
		t.Error(regionList, target)
	}

	regionList = append(regionList, schemas.RegionConfig{
		Region: target,
	})

	input = CheckRegionExist(target, regionList)
	expected = true

	if input != expected {
		t.Error(regionList, target)
	}
}
