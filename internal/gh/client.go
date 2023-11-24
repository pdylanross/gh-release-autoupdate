package gh

import (
	"context"
	"errors"
	"net/http"

	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"

	"github.com/google/go-github/v56/github"
)

func NewGithubClient(opts *types.GithubClientOptions) (*github.Client, error) {
	if opts != nil {
		return configureClient(opts)
	}

	return nil, errors.New("nil github client options")
}

func configureClient(opts *types.GithubClientOptions) (*github.Client, error) {
	if opts.UserAgent == "" {
		return nil, errors.New("GithubClientOptions.UserAgent is required")
	}

	if opts.HTTPClient == nil {
		opts.HTTPClient = &http.Client{}
	}

	cl := github.NewClient(opts.HTTPClient)
	cl.UserAgent = opts.UserAgent

	if opts.Authentication != nil {
		ctx, cancel := context.WithTimeout(context.Background(), opts.AuthTimeout)
		defer cancel()
		token, found, err := opts.Authentication.GetToken(ctx)
		if err != nil {
			return nil, err
		}

		if found {
			cl = cl.WithAuthToken(token)
		}
	}

	return cl, nil
}
