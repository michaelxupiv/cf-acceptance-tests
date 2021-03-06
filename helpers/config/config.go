package config

import (
	"time"
)

type CatsConfig interface {
	GetIncludeApps() bool
	GetIncludeBackendCompatiblity() bool
	GetIncludeContainerNetworking() bool
	GetIncludeDetect() bool
	GetIncludeDocker() bool
	GetIncludeInternetDependent() bool
	GetIncludePrivilegedContainerSupport() bool
	GetIncludeRouteServices() bool
	GetIncludeRouting() bool
	GetIncludeZipkin() bool
	GetIncludeSSO() bool
	GetIncludeSecurityGroups() bool
	GetIncludeServices() bool
	GetIncludeSsh() bool
	GetIncludeTasks() bool
	GetIncludeV3() bool
	GetShouldKeepUser() bool
	GetSkipSSLValidation() bool
	GetUseExistingUser() bool

	GetAdminPassword() string
	GetAdminUser() string
	GetApiEndpoint() string
	GetAppsDomain() string
	GetArtifactsDirectory() string
	GetBackend() string
	GetBinaryBuildpackName() string
	GetConfigurableTestPassword() string
	GetExistingUser() string
	GetExistingUserPassword() string
	GetGoBuildpackName() string
	GetJavaBuildpackName() string
	GetNamePrefix() string
	GetNodejsBuildpackName() string
	GetPersistentAppHost() string
	GetPersistentAppOrg() string
	GetPersistentAppQuotaName() string
	GetPersistentAppSpace() string
	GetRubyBuildpackName() string
	Protocol() string

	AsyncServiceOperationTimeoutDuration() time.Duration
	BrokerStartTimeoutDuration() time.Duration
	CfPushTimeoutDuration() time.Duration
	DefaultTimeoutDuration() time.Duration
	DetectTimeoutDuration() time.Duration
	GetScaledTimeout(time.Duration) time.Duration
	LongCurlTimeoutDuration() time.Duration
	LongTimeoutDuration() time.Duration
	SleepTimeoutDuration() time.Duration
}

func NewCatsConfig(path string) (CatsConfig, error) {
	return NewConfig(path)
}
