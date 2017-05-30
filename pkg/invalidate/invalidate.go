package invalidate

import "github.com/safeguardproperties/gomodels/cache"

//Invalidator has method to invalidate cache for given primaryContext, secondaryContext, key
type Invalidator interface {
	Invalidate(primaryContext string, secondaryContext string, key string) *cache.Result
}

//Client
type Client struct {
}

//New
func New() *Client {
	return &Client{}
}

//InvalidatePriSecKey erases cache value for provided primaryContext, secondaryContext, key
func (i *Impl) InvalidatePriSecKey(primaryContext string, secondaryContext string, key string) *cache.Result {
	return &cache.Result{}
}
