package types

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v56/github"
)

// GithubClientOptions sets up how we communicate with github.
type GithubClientOptions struct {
	HTTPClient          *http.Client
	UserAgent           string
	CustomConfigurators []func(*github.Client)
	Authentication      GithubAuthenticationProvider
	AuthTimeout         time.Duration
}

// NewOptions creates a new github client for a given repo.
func NewOptions(appName string, appVersion string) *GithubClientOptions {
	opt := &GithubClientOptions{AuthTimeout: time.Second * 2}
	return opt.WithApp(appName, appVersion)
}

// WithHTTPClient sets an override http client for GH API interactions.
func (o *GithubClientOptions) WithHTTPClient(cl *http.Client) *GithubClientOptions {
	o.HTTPClient = cl
	return o
}

// WithCustom allows for custom overrides of the GH client library.
func (o *GithubClientOptions) WithCustom(cfg func(client *github.Client)) *GithubClientOptions {
	o.CustomConfigurators = append(o.CustomConfigurators, cfg)
	return o
}

// WithApp sets the user agent of the GH client.
func (o *GithubClientOptions) WithApp(appName string, appVersion string) *GithubClientOptions {
	userAgent := fmt.Sprintf("%s/%s", appName, appVersion)
	o.UserAgent = userAgent
	return o
}

// WithAuthenticationProvider sets the authentication provider for this client.
func (o *GithubClientOptions) WithAuthenticationProvider(provider GithubAuthenticationProvider) *GithubClientOptions {
	o.Authentication = provider
	return o
}

// GithubAuthenticationProvider defines how the gh client gets authentication material (if relevant).
type GithubAuthenticationProvider interface {
	GetToken(ctx context.Context) (string, bool, error)
}
