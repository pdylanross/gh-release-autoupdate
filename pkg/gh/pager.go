package gh

import "github.com/google/go-github/v56/github"

// Pager is a utility to page through github API responses.
type Pager[T any] struct {
	gh          *github.Client
	getter      func(gh *github.Client, page *github.ListOptions) ([]T, *github.Response, error)
	currentPage int
}

// NewPager creates a new pager instance.
func NewPager[T any](gh *github.Client, getter func(gh *github.Client, page *github.ListOptions) ([]T, *github.Response, error)) (*Pager[T], error) {
	return &Pager[T]{gh: gh, getter: getter, currentPage: 1}, nil
}

// NextPage fetches the next page from gh.
func (p *Pager[T]) NextPage() ([]T, bool, error) {
	opts := &github.ListOptions{Page: p.currentPage}
	items, response, err := p.getter(p.gh, opts)
	if err != nil {
		return nil, false, err
	}

	p.currentPage++
	finished := false
	if p.currentPage > response.LastPage {
		finished = true
	}

	return items, finished, nil
}
