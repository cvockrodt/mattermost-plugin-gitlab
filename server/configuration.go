package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-plugin-gitlab/server/gitlab"
)

// configuration captures the plugin's external configuration as exposed in the Mattermost server
// configuration, as well as values computed from the configuration. Any public fields will be
// deserialized from the Mattermost server configuration in OnConfigurationChange.
//
// As plugins are inherently concurrent (hooks being called asynchronously), and the plugin
// configuration can change at any time, access to the configuration must be synchronized. The
// strategy used in this plugin is to guard a pointer to the configuration, and clone the entire
// struct whenever it changes. You may replace this with whatever strategy you choose.
//
// If you add non-reference types to your configuration struct, be sure to rewrite Clone as a deep
// copy appropriate for your types.
type configuration struct {
	GitlabURL                   string
	GitlabOAuthClientID         string
	GitlabOAuthClientSecret     string
	WebhookSecret               string
	EncryptionKey               string
	GitlabGroup                 string
	EnablePrivateRepo           bool
	PluginsDirectory            string
	UsePreregisteredApplication bool
}

// Clone shallow copies the configuration. Your implementation may require a deep copy if
// your configuration has reference types.
func (c *configuration) Clone() *configuration {
	var clone = *c
	return &clone
}

// IsValid checks if all needed fields are set.
func (c *configuration) IsValid() error {
	if _, err := url.ParseRequestURI(c.GitlabURL); err != nil {
		return errors.New("must have a valid GitLab URL")
	}

	if !c.UsePreregisteredApplication {
		if c.GitlabOAuthClientID == "" {
			return fmt.Errorf("must have a GitLab oauth client id")
		}
		if c.GitlabOAuthClientSecret == "" {
			return fmt.Errorf("must have a GitLab oauth client secret")
		}
	}

	gitLabURL := strings.TrimSuffix(c.GitlabURL, "/")
	if c.UsePreregisteredApplication && gitLabURL != "https://gitlab.com" {
		return errors.New("pre-registered application can only be used with official public GitLab")
	}

	if c.EncryptionKey == "" {
		return fmt.Errorf("must have an encryption key")
	}

	return nil
}

// getConfiguration retrieves the active configuration under lock, making it safe to use
// concurrently. The active configuration may change underneath the client of this method, but
// the struct returned by this API call is considered immutable.
func (p *Plugin) getConfiguration() *configuration {
	p.configurationLock.RLock()
	defer p.configurationLock.RUnlock()

	if p.configuration == nil {
		return &configuration{}
	}

	return p.configuration
}

// setConfiguration replaces the active configuration under lock.
//
// Do not call setConfiguration while holding the configurationLock, as sync.Mutex is not
// reentrant. In particular, avoid using the plugin API entirely, as this may in turn trigger a
// hook back into the plugin. If that hook attempts to acquire this lock, a deadlock may occur.
//
// This method panics if setConfiguration is called with the existing configuration. This almost
// certainly means that the configuration was modified without being cloned and may result in
// an unsafe access.
func (p *Plugin) setConfiguration(configuration *configuration, serverConfiguration *model.Config) {
	p.configurationLock.Lock()
	defer p.configurationLock.Unlock()

	if configuration != nil && p.configuration == configuration {
		// Ignore assignment if the configuration struct is empty. Go will optimize the
		// allocation for same to point at the same memory address, breaking the check
		// above.
		if reflect.ValueOf(*configuration).NumField() == 0 {
			return
		}

		panic("setConfiguration called with the existing configuration")
	}

	// PluginDirectory should be set based on server configuration and not the plugin configuration
	if serverConfiguration.PluginSettings.Directory != nil {
		configuration.PluginsDirectory = *serverConfiguration.PluginSettings.Directory
	}

	p.configuration = configuration
}

// OnConfigurationChange is invoked when configuration changes may have been made.
func (p *Plugin) OnConfigurationChange() error {
	var configuration = new(configuration)

	// Load the public configuration fields from the Mattermost server configuration.
	if err := p.API.LoadPluginConfiguration(configuration); err != nil {
		return errors.Wrap(err, "failed to load plugin configuration")
	}

	serverConfiguration := p.API.GetConfig()

	p.setConfiguration(configuration, serverConfiguration)

	if err := configuration.IsValid(); err != nil {
		return err
	}

	p.GitlabClient = gitlab.New(configuration.GitlabURL, configuration.GitlabGroup, p.isNamespaceAllowed)

	return nil
}
