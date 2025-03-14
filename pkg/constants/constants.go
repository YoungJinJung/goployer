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

package constants

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Current Base Year
	YearNow = 2020

	// DefaultLogLevel is the default global verbosity
	DefaultLogLevel = logrus.WarnLevel

	// DefaultRegionVariable is the default region id
	DefaultRegionVariable = "AWS_DEFAULT_REGION"

	// EmptyString is the empty string
	EmptyString = ""

	// ZeroFloat64 means 0 in float64 format
	ZeroFloat64 = float64(0)

	// DefaultProfile is the default aws profile
	DefaultProfile = "default"

	// NoManifestFileExists is the error message when there is no manifest file
	NoManifestFileExists = "manifest file does not exist"

	// InServiceStatus means status in service
	InServiceStatus = "InService"

	// PendingStatus means pending status
	PendingStatus = "Pending"

	// TestString is "test"
	TestString = "test"

	// StringText is "string"
	StringText = "string"

	// DefaultSpotAllocationStrategy means default spot strategy for spot allocation
	DefaultSpotAllocationStrategy = "lowest-price"

	// DefaultDeploymentTimeout is default deployment timeout
	DefaultDeploymentTimeout = 60 * time.Minute

	// DefaultPollingInterval is default polling interval
	DefaultPollingInterval = 60 * time.Second

	// MinPollingInterval is minimum polling interval
	MinPollingInterval = 5 * time.Second

	// MetricYamlPath is the path of metric manifest file
	MetricYamlPath = "metrics.yaml"

	// TestMetricYamlPath is the relative path for test metrics file
	TestMetricYamlPath = "../../test/metrics_test.yaml"

	// DefaultMetricStorageType is the default storage type for metrics
	DefaultMetricStorageType = "dynamodb"

	// DefaultRegion is a default region
	DefaultRegion = "us-east-1"

	// DefaultHealthcheckType is the default healthcheck type
	DefaultHealthcheckType = "EC2"

	// DefaultHealthcheckGracePeriod is the default healthcheck grace period
	DefaultHealthcheckGracePeriod = 300

	// DefaultInstanceWarmup is the default duration for instance warmup
	DefaultInstanceWarmup = 300

	// DefaultMinHealthyPercentage is the default value of minimum healthy instance percentage for refresh
	DefaultMinHealthyPercentage = 90

	// S3Prefix is prefix of s3 URL
	S3Prefix = "s3://"

	// HashKey is the default value of hash key for metric table
	HashKey = "identifier"

	// DefaultReadThroughput is the default value of RCU
	DefaultReadThroughput = int64(5)

	// DefaultWriteThroughput is the default value of RCU
	DefaultWriteThroughput = int64(5)

	// InitialStatus is the initial status of instances in classic LB
	InitialStatus = "Not Found"

	// MonthToSec changes a month to seconds
	MonthToSec = float64(2592000)

	// DayToSec changes a day to seconds
	DayToSec = int64(86400)

	// HourToSec changes an hour to seconds
	HourToSec = int64(3600)

	// StepCheckPrevious = CheckPrevious
	StepCheckPrevious = int64(1)

	// StepDeploy = Deploy
	StepDeploy = int64(2)

	// StepAdditionalWork = FinishAdditionalWork
	StepAdditionalWork = int64(3)

	// StepTriggerLifecycleCallback = TriggerLifecycleCallbacks
	StepTriggerLifecycleCallback = int64(4)

	// StepCleanPreviousVersion = CleanPreviousVersion
	StepCleanPreviousVersion = int64(5)

	// StepCleanChecking = CleanChecking
	StepCleanChecking = int64(6)

	// StepGatherMetrics = GatherMetrics
	StepGatherMetrics = int64(7)

	// StepRunAPI = RunAPI
	StepRunAPI = int64(8)

	// DefaultEnableStats is whether or not to enable gathering stats
	DefaultEnableStats = true

	// ALL means ALL as string
	ALL = "ALL"

	// SlackToken is environment key for slack token
	SlackToken = "SLACK_TOKEN"

	// SlackChannel is environment key for slack channel
	SlackChannel = "SLACK_CHANNEL"

	// SlackWebHookURL is environment key for slack webhook url
	SlackWebHookURL = "SLACK_WEBHOOK_URL"

	// MinAPITestDuration is minimum duration of API test
	MinAPITestDuration = 1 * time.Second

	// DeploymentTagKey is a tag key for indicating deployment
	DeploymentTagKey = "goployer-deployment"

	// CanaryMark is a mark indicating that resources are related to Canary deployment
	CanaryMark = "canary"

	// Deployment Methods
	BlueGreenDeployment     = "bluegreen"
	CanaryDeployment        = "canary"
	RollingUpdateDeployment = "rollingupdate"
	DeployOnly              = "deployonly"

	DelimiterRegex = "[,/|!@$%^&*_=`~]+"
)

var (
	// LogLevelMapper is the default global verbosity
	LogLevelMapper = map[string]logrus.Level{
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"warn":  logrus.WarnLevel,
		"trace": logrus.TraceLevel,
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
	}

	// AWSCredentialsPath is the file path of aws credentials
	AWSCredentialsPath = HomeDir() + "/.aws/credentials"

	// AWSConfigPath is the file path of aws config
	AWSConfigPath = HomeDir() + "/.aws/config"

	// AvailableBlockTypes is a list of available ebs block types
	AvailableBlockTypes = []string{"io1", "io2", "gp2", "gp3", "st1", "sc1"}

	// IopsRequiredBlockType is a list of ebs type which requires iops
	IopsRequiredBlockType = []string{"io1", "io2"}

	// AllowedRequestMethod is a list of request method
	AllowedRequestMethod = []string{"GET", "POST", "PUT"}

	// TimeFields is a list of time.Time field
	TimeFields = []string{"timeout", "polling-interval"}

	// ProhibitedTags is a list of prohibited tags which are going to be attached by goployer
	ProhibitedTags = []string{"Name", "stack"}

	// StatusTimeStampKey is a map of timestamp keys with deployment status
	StatusTimeStampKey = map[string]string{
		"deployed":   "deployed_date",
		"terminated": "terminated_date",
	}

	// AllowedAnswerYes is a list of allowed answers with yes
	AllowedAnswerYes = []string{"y", "yes"}

	// DaysOfWeek is a list of possible string value for cron expression
	DaysOfWeek = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN", "0", "1", "2", "3", "4", "5", "6", "7"}

	// MinTimestamp means minimum timestamp YEAR/01/01 00:00:00 UTC
	MinTimestamp = time.Date(YearNow, time.January, 1, 0, 0, 0, 0, time.UTC)
)

// Get Home Directory
func HomeDir() string {
	if h := os.Getenv("HOME"); h != EmptyString {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
