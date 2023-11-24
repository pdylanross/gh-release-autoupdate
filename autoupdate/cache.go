package autoupdate

import "time"

type CacheOptions struct {
	// PackageName overrides the PackageName for use with caching.
	// The cache will store metadata in os.UserCacheDir()/PackageName
	PackageName *string

	// CheckInterval specifies the maximum frequency at which update checks will be performed
	CheckInterval time.Duration
}

func DefaultCacheOptions() *CacheOptions {
	return &CacheOptions{
		PackageName:   nil,
		CheckInterval: time.Hour * 24,
	}
}
