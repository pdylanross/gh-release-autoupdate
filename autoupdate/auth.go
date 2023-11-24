package autoupdate

import (
	"context"
	"os"

	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/pdylanross/gh-release-autoupdate/autoupdate/types"
)

func DefaultAuthProvider() types.GithubAuthenticationProvider {
	return &defaultAuthProvider{}
}

type defaultAuthProvider struct{}

func (d *defaultAuthProvider) GetToken(ctx context.Context) (string, bool, error) {
	if err := ctx.Err(); err != nil {
		return "", false, err
	}

	if val, ok := os.LookupEnv("GH_TOKEN"); ok {
		return val, true, nil
	}

	token, _ := auth.TokenForHost("github.com")
	if token != "" {
		return token, true, nil
	}

	return "", false, nil
}
