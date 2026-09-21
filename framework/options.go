package framework

import (
	"time"

	"github.com/redhat-appstudio/helmet/api"
	"github.com/redhat-appstudio/helmet/internal/config"
	"github.com/redhat-appstudio/helmet/internal/mcptools"
)

// Option represents a functional option for the App runtime.
// These options configure runtime dependencies and behavior.
// For application metadata (name, version, etc.), use ContextOption with NewAppContext.
type Option func(*App)

// WithIntegrations sets the supported integrations for the application.
func WithIntegrations(modules ...api.IntegrationModule) Option {
	return func(a *App) {
		a.integrations = append(a.integrations, modules...)
	}
}

// WithLoadCreateConfig replaces the default "config --create" loader.
func WithLoadCreateConfig(fn config.CreateConfigLoader) Option {
	return func(a *App) {
		a.loadCreateConfig = fn
	}
}

// WithDistributedInstallerMergeLayout configures "config --create" without a file
// argument to merge installer/config/settings.yaml, installer/helmet.yaml,
// and charts/<chart>/config.yaml fragments (see config.MergeDistributedInstallerYAML).
// Overrides the default single-file loader.
func WithDistributedInstallerMergeLayout() Option {
	return func(a *App) {
		a.mergedInstallerConfig = true
		a.loadCreateConfig = distributedInstallerMergeLoader
	}
}

// WithImage sets the container image for the installer.
func WithImage(image string) Option {
	return func(a *App) {
		a.image = image
	}
}

// Deprecated: use WithImage instead.
func WithMCPImage(image string) Option {
	return WithImage(image)
}

// WithMCPToolsBuilder sets the MCP tools builder for the application.
func WithMCPToolsBuilder(builder mcptools.MCPToolsBuilder) Option {
	return func(a *App) {
		a.mcpToolsBuilder = builder
	}
}

// WithInstallerTarball sets the embedded installer tarball for the application.
func WithInstallerTarball(tarball []byte) Option {
	return func(a *App) {
		a.installerTarball = tarball
	}
}

// WithVerifyRetries sets how many times helm test runs after each chart deploy.
// Values less than 1 are treated as 1. Default is 3 when this option is omitted.
func WithVerifyRetries(retries int) Option {
	return func(a *App) {
		if retries < 1 {
			retries = 1
		}
		a.flags.VerifyRetries = retries
	}
}

// WithVerifyRetryDelay sets the pause between failed helm test attempts.
// Default is 1 minute when this option is omitted. Negative values become 0.
func WithVerifyRetryDelay(delay time.Duration) Option {
	return func(a *App) {
		if delay < 0 {
			delay = 0
		}
		a.flags.VerifyRetryDelay = delay
	}
}
