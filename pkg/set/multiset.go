package set

import "github.com/safeguardproperties/gomodels/cache"

//MultiSetter has method to set values for given primaryContext, secondaryContext, cacheItem
type MultiSetter interface {
	Set(primaryContext string, secondaryContext string, items []cache.Item) map[string]cache.Result
}
