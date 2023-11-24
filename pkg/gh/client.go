package gh

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v56/github"
)

type GithubClientOptions struct {
	HTTPClient          *http.Client
	UserAgent           string
	CustomConfigurators []func(*github.Client)
}

func NewOptions(appName string, appVersion string) *GithubClientOptions {
	opt := &GithubClientOptions{}
	return opt.WithApp(appName, appVersion)
}

func (o *GithubClientOptions) WithHTTPClient(cl *http.Client) *GithubClientOptions {
	o.HTTPClient = cl
	return o
}

func (o *GithubClientOptions) WithCustom(cfg func(client *github.Client)) *GithubClientOptions {
	o.CustomConfigurators = append(o.CustomConfigurators, cfg)
	return o
}

func (o *GithubClientOptions) WithApp(appName string, appVersion string) *GithubClientOptions {
	userAgent := fmt.Sprintf("%s/%s", appName, appVersion)
	o.UserAgent = userAgent
	return o
}

func NewGithubClient(opts *GithubClientOptions) (*github.Client, error) {
	if opts != nil {
		return configureClient(opts)
	}

	return nil, errors.New("nil github client options")
}

func configureClient(opts *GithubClientOptions) (*github.Client, error) {
	if opts.UserAgent == "" {
		return nil, errors.New("GithubClientOptions.UserAgent is required")
	}

	if opts.HTTPClient == nil {
		opts.HTTPClient = &http.Client{}
	}

	cl := github.NewClient(opts.HTTPClient)
	cl.UserAgent = opts.UserAgent

	return cl, nil
}
