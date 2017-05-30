package get

import "github.com/safeguardproperties/gomodels/cache"

//MultiGetter has method to get values for given primaryContext, secondaryContext, array of keys
type MultiGetter interface {
	Get(primaryContext string, secondaryContext string, keys []string) map[string]cache.Result
}
